package mysql

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/metadata"
	"db-viewer/internal/engine/metadata/mySQLInspector"
	
	queryexecutor "db-viewer/internal/engine/queryExecutor"
	"db-viewer/internal/engine/queryExecutor/sqlExecutor"
	"db-viewer/internal/engine/transports"
)

type Driver struct{}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Name() string {
	return "mysql"
}

func (d *Driver) Create(ctx context.Context, config entities.ConnectionConfig, transport transports.Transport) (manager.Connection, error) {

	conn := New(config, transport)
	return conn, nil
}

func (d *Driver) Executor() queryexecutor.Executor {
	return sqlExecutor.New()
}

func (d *Driver) Inspector() metadata.Inspector {
	return mySQLInspector.NewInspector()
}

