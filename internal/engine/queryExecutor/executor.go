package queryexecutor

import (
	"context"
	manager "db-viewer/internal/engine/connectionManager"
	"time"
)

type ColumnInfo struct {
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
}

type QueryResult struct {
	Columns      []ColumnInfo    `json:"columns"`
	Rows         [][]interface{} `json:"rows"`
	RowsAffected int64           `json:"rowsAffected"`
	LastInsertId int64           `json:"lastInsertId"`
	Duration     time.Duration   `json:"duration"`
	IsQuery      bool            `json:"isQuery"`
}

type Executor interface {
	Execute(ctx context.Context, conn manager.Connection, query string, args ...any) (*QueryResult, error)
}

