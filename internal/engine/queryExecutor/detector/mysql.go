package detector


import "strings"

type MySQL struct{}

func(d MySQL) Detect(statement string) StatementKind {
	upper := strings.ToUpper(strings.TrimSpace(statement))

	switch {
	case strings.HasPrefix(upper, "INSERT"),
		strings.HasPrefix(upper, "UPDATE"),
		strings.HasPrefix(upper, "DELETE"),
		strings.HasPrefix(upper, "REPLACE"):
		return KindExec
	default:
		// SELECT, SHOW, EXPLAIN, WITH, CALL, DDL, anything unsure — go through QueryContext.
		return KindQuery
	}
}