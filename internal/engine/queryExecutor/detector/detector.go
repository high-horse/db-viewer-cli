package detector


type StatementKind int

const (
	KindUnknown StatementKind = iota
	KindQuery  // expect rows back
	KindExec // expect RowsAffected/LastInsertId	
)

type Detector interface {
	Detect(statement string) StatementKind
}



