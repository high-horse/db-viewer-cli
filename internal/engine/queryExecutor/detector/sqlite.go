package detector

import "strings"

type SQLite struct{}

func (d SQLite) Detect(statement string) StatementKind {
	upper := strings.ToUpper(strings.TrimSpace(statement))

	switch {
	case strings.HasPrefix(upper, "INSERT"),
		strings.HasPrefix(upper, "UPDATE"),
		strings.HasPrefix(upper, "DELETE"):
		return KindExec
	default:
		// PRAGMA is ambiguous (some pragmas return rows, some don't) — stay safe.
		return KindQuery
	}
}