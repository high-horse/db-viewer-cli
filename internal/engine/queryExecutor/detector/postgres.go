package detector

import (
	"regexp"
	"strings"
)

var pgReturning = regexp.MustCompile(`(?i)\bRETURNING\b`)

type Postgres struct{}

func (d Postgres) Detect(statement string) StatementKind {
	upper := strings.ToUpper(strings.TrimSpace(statement))

	if pgReturning.MatchString(statement) {
		return KindQuery // INSERT/UPDATE/DELETE ... RETURNING produces rows
	}

	switch {
	case strings.HasPrefix(upper, "INSERT"),
		strings.HasPrefix(upper, "UPDATE"),
		strings.HasPrefix(upper, "DELETE"):
		return KindExec
	default:
		return KindQuery
	}
}