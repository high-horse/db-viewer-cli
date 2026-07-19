package drivers

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/metadata"
	queryexecutor "db-viewer/internal/engine/queryExecutor"
	"db-viewer/internal/engine/transports"
)


type Driver interface {
	Name() string

	Create(
		ctx context.Context,
		config entities.ConnectionConfig,
		transport transports.Transport,
	) (manager.Connection, error)

	Executor() queryexecutor.Executor

	Inspector() metadata.Inspector
}