package parser

type AbapFile struct {
	Statements []Stmt
}

type Node interface {
	Node()
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type IntroductoryStmt interface {
	Stmt
	introductoryStmtNode()
}

type Program struct {
	IntroductoryStmt IntroductoryStmt
}

type ReportStmt struct {
	Name                 string
	MsgID                *string
	ReducedFunctionality *bool
	ListOptions          *ListOptions
}

// Node implements IntroductoryStmt.
func (r *ReportStmt) Node() {

}

// introductoryStmtNode implements IntroductoryStmt.
func (r *ReportStmt) introductoryStmtNode() {
}

// stmtNode implements Stmt.
func (r *ReportStmt) stmtNode() {

}

var _ IntroductoryStmt = &ReportStmt{}

type pageLines struct {
	PageLines   *int
	FooterLines *int
}

type ListOptions struct {
	NoStandardPageHeading *bool
	LineSize              *int
	PageLines             *int
	FooterLines           *int
}

type reportOpts struct {
	MsgID                *string
	ReducedFunctionality *bool
	ListOptions          *ListOptions
}
