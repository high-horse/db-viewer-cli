package metadata

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
)


type Inspector interface {
	ListDatabases(ctx context.Context, conn manager.Connection) ([]entities.DatabaseInfo, error)

	ListTables(ctx context.Context, conn manager.Connection)([]entities.InspectTableInfo, error)

	ListColumns (ctx context.Context, conn manager.Connection, table string) ([]entities.InspectColumnInfo, error)

	// TODO: for later
	// ListIndexes()
	// ListForeignKeys()
	// ListViews()	
	// ListProcedures()
}