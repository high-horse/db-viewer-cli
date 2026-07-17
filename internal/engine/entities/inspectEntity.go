package entities

import "time"

type InspectTableInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // TABLE, VIEW
	Database    string `json:"database"`
	Schema      string `json:"schema,omitempty"`

	Rows        int64  `json:"rows,omitempty"`
	Engine      string `json:"engine,omitempty"`
	Comment     string `json:"comment,omitempty"`

	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type InspectColumnInfo struct {
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
	Nullable     bool   `json:"nullable"`
	DefaultValue any    `json:"defaultValue"`
	PrimaryKey   bool   `json:"primaryKey"`
	AutoIncrement bool  `json:"autoIncrement"`
	Length         int64  `json:"length,omitempty"`
}