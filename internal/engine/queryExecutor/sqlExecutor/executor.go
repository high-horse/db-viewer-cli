package sqlExecutor

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"fmt"

	manager "db-viewer/internal/engine/connectionManager"
	executor "db-viewer/internal/engine/queryExecutor"
	"db-viewer/internal/engine/queryExecutor/detector"
)

type Executor struct{}

func New() *Executor {
	return &Executor{}
}

func (e *Executor) ExecuteWithDetector(
	ctx context.Context,
	conn manager.Connection,
	query string,
	args ...any,
) (*executor.QueryResult, error) {
	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection %q is not a SQL connection", conn.ID())
	}

	db := sqlConn.DB()
	if db == nil {
		return nil, fmt.Errorf("connection %q is not connected", conn.ID())
	}

	kind := detector.For(conn.Type()).Detect(query)
	start := time.Now()

	if kind == detector.KindExec {
		res, err := db.ExecContext(ctx, query, args...)
		if err != nil {
			rowsAffected, _ := res.RowsAffected()
			lastInsertId, _ := res.LastInsertId()
			return &executor.QueryResult{
				RowsAffected: rowsAffected,
				LastInsertId: lastInsertId,
				Duration:     time.Since(start),
				IsQuery:      false,
			}, nil
		}
		// Detector guessed wrong (statement actually returns rows) — fall through to QueryContext.
	}

	return e.runQuery(ctx, db, query, start, args...)

}

func (e *Executor) Execute(
	ctx context.Context,
	conn manager.Connection,
	query string,
	args ...any,
) (*executor.QueryResult, error) {
	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection %q is not a SQL connection", conn.ID())
	}

	db := sqlConn.DB()
	if db == nil {
		return nil, fmt.Errorf("connection %q is not connected", conn.ID())
	}

	start := time.Now()
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	// No columns => it wasn't a row-returning statement (INSERT/UPDATE/DDL/etc).
	// It already ran, exactly once, via QueryContext above.
	if len(colTypes) == 0 {
		return &executor.QueryResult{
			Duration: time.Since(start),
			IsQuery:  false,
			// RowsAffected intentionally omitted — not obtainable from sql.Rows.
			// See note below if you need it.
		}, nil
	}

	columns := make([]executor.ColumnInfo, len(colTypes))
	for i, ct := range colTypes {
		columns[i] = executor.ColumnInfo{Name: ct.Name(), DatabaseType: ct.DatabaseTypeName()}
	}

	var result [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(colTypes))
		ptrs := make([]interface{}, len(colTypes))
		for i := range values {
			ptrs[i] = &values[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		result = append(result, normalizeRow(values))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &executor.QueryResult{
		Columns:  columns,
		Rows:     result,
		Duration: time.Since(start),
		IsQuery:  true,
	}, nil
}


func (e *Executor) ExecuteOld(
	ctx context.Context,
	conn manager.Connection,
	query string,
	args ...any,
) (*executor.QueryResult, error) {

	sqlConn, ok := conn.(manager.SQLConnection)
	if !ok {
		return nil, fmt.Errorf("connection %q is not a SQL connection", conn.ID())
	}

	db := sqlConn.DB()
	if db == nil {
		return nil, fmt.Errorf("connection %q is not connected", conn.ID())
	}

	start := time.Now()

	if isQueryStatement(query) {
		return e.runQuery(ctx, db, query, start, args...)
	}
	return e.runExec(ctx, db, query, start, args...)
}

func (e *Executor) runQuery(ctx context.Context, db *sql.DB, query string, start time.Time, args ...any) (*executor.QueryResult, error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	columns := make([]executor.ColumnInfo, len(colTypes))
	for i, ct := range colTypes {
		columns[i] = executor.ColumnInfo{
			Name:         ct.Name(),
			DatabaseType: ct.DatabaseTypeName(),
		}
	}

	var results [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(colTypes))
		ptrs := make([]interface{}, len(colTypes))

		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		results = append(results, normalizeRow(values))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &executor.QueryResult{
		Columns:  columns,
		Rows:     results,
		Duration: time.Since(start),
		IsQuery:  true,
	}, nil
}

func (e *Executor) runExec(ctx context.Context, db *sql.DB, query string, start time.Time, args ...any) (*executor.QueryResult, error) {
	res, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()
	lastInsertId, _ := res.LastInsertId()

	return &executor.QueryResult{
		RowsAffected: rowsAffected,
		LastInsertId: lastInsertId,
		Duration:     time.Since(start),
		IsQuery:      false,
	}, nil
}

// normalizeRow converts []byte (how most drivers return TEXT/VARCHAR/DECIMAL)
// into string so the result is JSON/JS-friendly instead of base64.
func normalizeRow(values []interface{}) []interface{} {
	out := make([]interface{}, len(values))
	for i, v := range values {
		if b, ok := v.([]byte); ok {
			out[i] = string(b)
		} else {
			out[i] = v
		}
	}
	return out
}

func isQueryStatement(query string) bool {
	trimmed := strings.TrimSpace(query)
	if trimmed == "" {
		return false
	}

	for strings.HasPrefix(trimmed, "--") {
		lines := strings.SplitN(trimmed, "\n", 2)
		if len(lines) == 1 {
			return false
		}
		trimmed = strings.TrimSpace(lines[1])
	}

	firstWord := strings.ToUpper(trimmed)

	// Remove SQL comments
	for strings.HasPrefix(trimmed, "--") {
		lines := strings.SplitN(trimmed, "\n", 2)
		if len(lines) == 1 {
			return false
		}
		trimmed = strings.TrimSpace(lines[1])
	}

	for _, prefix := range []string{
		"SELECT",
		"SHOW",
		"EXPLAIN",
		"WITH",
		"DESCRIBE",
		"DESC",
		"PRAGMA", "SELECT",
		"SHOW",
		"DESCRIBE",
		"DESC",
		"EXPLAIN",
		"ANALYZE",
		"WITH",
		"VALUES",
		"TABLE",
		"CALL",
	} {
		if strings.HasPrefix(firstWord, prefix) {
			return true
		}
	}
	return false
}
