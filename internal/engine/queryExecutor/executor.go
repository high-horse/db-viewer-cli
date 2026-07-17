package queryexecutor

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
)



type Executor interface {
	Execute(ctx context.Context, conn manager.Connection, query string, args ...any) (*entities.QueryResult, error)
}

