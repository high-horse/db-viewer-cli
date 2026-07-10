package mysql

import (
	"context"
	"database/sql"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/transports"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


type Connection struct{
	config entities.ConnectionConfig

	transport transports.Transport

	db *sql.DB
	connected bool
}

func New(config entities.ConnectionConfig, transport transports.Transport) *Connection {
	return &Connection{
		config: config,
		transport: transport,
	}
}

func(c *Connection) ID() string {
	return c.config.ID
}

func(c *Connection) Name() string {
	return c.config.Name
}

func(c *Connection) Type() string {
	return "mysql"
}

func(c *Connection) DB() *sql.DB {
	return  c.db
}

func (c *Connection) dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		c.config.User,
		c.config.Password,
		c.transport.Address(),
		c.config.Database,
	)
}

func(c *Connection) Connect(ctx context.Context) error {
	fmt.Println("transporting ...")
	if err := c.transport.Connect(ctx); err != nil {
		return err
	}

	fmt.Println("transporting success\nOpening database")
	fmt.Println("dsn", c.dsn())
	db, err := sql.Open(
		"mysql",
		c.dsn(),
	)

	if err != nil {
		fmt.Println("Failed to open with dsn", c.dsn())
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return err
	}

	c.db = db
	c.connected = true

	return nil
}

func(c *Connection) Disconnect() error {
	if c.db == nil {
		return nil
	}

	err := c.db.Close()
	c.connected = false
	return err
}

func (c *Connection) Ping(ctx context.Context) error {

	if c.db == nil {
		return fmt.Errorf("mysql connection not initialized")
	}

	return c.db.PingContext(ctx)
}

func (c *Connection) IsConnected() bool {
	return c.connected
}