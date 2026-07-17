package detector

var registry = map[string]Detector {
	"mysql": MySQL{},
	"mariadb": MySQL{},
	"postgres": Postgres{},
	"sqlite": SQLite{},
}

func For(driverName string) Detector {
	if d, ok := registry[driverName]; ok {
		return d
	}
	return fallback{}
}

type fallback struct {}

func(fallback) Detect(string) StatementKind {return KindQuery}