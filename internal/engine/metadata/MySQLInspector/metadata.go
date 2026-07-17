package MySQLInspector

import (
	"context"
	"database/sql"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/entities"
	"fmt"
)


type MySQLInspector struct {}

func NewInspector() *MySQLInspector {
	return &MySQLInspector{}
}

func (m *MySQLInspector) ListDatabases(ctx context.Context, conn manager.Connection) ([]entities.DatabaseInfo, error) {
	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection is not SQL")
	}
	rows, err := sqlConn.DB().QueryContext(ctx, "SHOW DATABASES;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var results []entities.DatabaseInfo

	for rows.Next(){
		var name string
		rows.Scan(&name)
		results = append(results, entities.DatabaseInfo{Name: name})
	}
	return results, nil
}

func (m *MySQLInspector) ListTables(ctx context.Context, conn manager.Connection)([]entities.InspectTableInfo, error) {
	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection is not SQL")
	}


	query := `
		SELECT
			TABLE_NAME,
			TABLE_TYPE,
			TABLE_SCHEMA,
			ENGINE,
			TABLE_ROWS,
			TABLE_COMMENT,
			CREATE_TIME,
			UPDATE_TIME
		FROM information_schema.TABLES
		WHERE TABLE_SCHEMA = ?
		ORDER BY TABLE_NAME
	`

	rows, err := sqlConn.DB().QueryContext(ctx, query, conn.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []entities.InspectTableInfo

	for rows.Next() {

		var table entities.InspectTableInfo 
		var (
			engine  sql.NullString
			comment sql.NullString
			rowsCnt sql.NullInt64

			created sql.NullTime
			updated sql.NullTime
		)

		err := rows.Scan(
			&table.Name,
			&table.Type,
			&table.Database,
			&engine,
			&rowsCnt,
			&comment,
			&created,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		table.Schema = table.Database

		if engine.Valid {
			table.Engine = engine.String
		}

		if comment.Valid {
			table.Comment = comment.String
		}

		if rowsCnt.Valid {
			table.Rows = rowsCnt.Int64
		}

		if created.Valid {
			table.CreatedAt = &created.Time
		}

		if updated.Valid {
			table.UpdatedAt = &updated.Time
		}


		tables = append(tables, table)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}


func (i *MySQLInspector) ListColumns(
	ctx context.Context,
	conn manager.Connection,
	database string,
	table string,
) ([]entities.InspectColumnInfo, error) {

	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection is not SQL")
	}

	query := `
		SELECT
			COLUMN_NAME,
			DATA_TYPE,
			IS_NULLABLE,
			COLUMN_DEFAULT,
			COLUMN_KEY,
			EXTRA
		FROM information_schema.COLUMNS
		WHERE TABLE_SCHEMA = ?
		AND TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION
	`

	rows, err := sqlConn.DB().QueryContext(
		ctx,
		query,
		database,
		table,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var columns []entities.InspectColumnInfo

	for rows.Next() {

		var (
			col entities.InspectColumnInfo

			isNullable string
			columnKey  string
			extra      string
		)

		err := rows.Scan(
			&col.Name,
			&col.DatabaseType,
			&isNullable,
			&col.DefaultValue,
			&columnKey,
			&extra,
		)

		if err != nil {
			return nil, err
		}


		col.Nullable = isNullable == "YES"

		col.PrimaryKey = columnKey == "PRI"

		col.AutoIncrement =
			extra == "auto_increment"

		columns = append(columns, col)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}