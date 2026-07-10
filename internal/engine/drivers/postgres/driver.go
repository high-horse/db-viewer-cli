
package postgres

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/transports"
)

type Driver struct{}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Name() string {
	return "pgx"
}

func (d *Driver) Create(ctx context.Context, config entities.ConnectionConfig, transport transports.Transport) (manager.Connection, error) {

	conn := New(config, transport)
	return conn, nil
}
