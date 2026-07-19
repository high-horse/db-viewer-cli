package postgres

import (
	"context"
	"database/sql"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/transports"
	"fmt"
	"log"

	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
)


type Connection struct {
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

func(c *Connection) DatabaseName() string {
	return c.config.Database
}

func(c *Connection) Type() string {
	return "postgres"
}

func(c *Connection) DB() *sql.DB {
	return  c.db
}

func (c *Connection) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		c.transport.Address(),
		c.config.User,
		c.config.Password,
		c.config.Database,
	)
}

func(c *Connection) Connect(ctx context.Context) error {
	log.Println("conneccting to postgres pgx")
	if err := c.transport.Connect(ctx); err != nil {
		return err
	}

	db, err := sql.Open("pgx", c.dsn())
	if err != nil {
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

func(c *Connection) Ping(ctx context.Context) error  {
	if c.db == nil {
		return fmt.Errorf("postgres connection not initialized")
	}

	return c.db.PingContext(ctx)
}

func(c *Connection) IsConnected() bool {
	return c.connected
}
