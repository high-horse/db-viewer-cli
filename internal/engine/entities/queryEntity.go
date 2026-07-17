package entities

import "time"

type ColumnInfo struct {
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
	Nullable     bool   `json:"nullable"`
}

type QueryResult struct {
	Columns      []ColumnInfo    `json:"columns"`
	Rows         [][]interface{} `json:"rows"`
	RowsAffected int64           `json:"rowsAffected"`
	LastInsertId int64           `json:"lastInsertId"`
	Duration     time.Duration   `json:"duration"`
	IsQuery      bool            `json:"isQuery"`
}


type TableInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // TABLE, VIEW
	Database string `json:"database"`
}

type Column struct {
	Name string
	Type string
}
