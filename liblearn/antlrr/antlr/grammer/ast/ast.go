package ast

import (
	"fmt"
	"holy-cmd/antlr/grammer/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Expr struct {
	Line   uint
	Column uint
	Text   string
}

type HelloSpec struct {
	Statements []HelloStatement
}

type HelloStatement struct {
	HelloExpr Expr
	MsgExpr   Expr
}

type HelloVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (h *HelloVisitor) VisitHello(ctx *parser.HelloContext) interface{} {
	var list HelloSpec
	for _, e := range ctx.AllSpec() {
		spec := e.Accept(h)
		if statement, ok := spec.(HelloStatement); ok {
			list.Statements = append(list.Statements, statement)
		}
	}
	return list
}

func (h *HelloVisitor) VisitSpec(ctx *parser.SpecContext) interface{} {
	if len(ctx.AllID()) != 2 {
		panic("invalid statement")
	}
	var statement HelloStatement
	hello := ctx.ID(0)
	msg := ctx.ID(1)

	statement.HelloExpr = h.getExpr(hello)
	statement.MsgExpr = h.getExpr(msg)
	if statement.HelloExpr.Text != "Hello" {
		panic(
			fmt.Sprintf(
				"Expected Hello, actual is %q", statement.HelloExpr.Text,
			),
		)
	}
	return statement
}

func (h *HelloVisitor) getExpr(node antlr.TerminalNode) Expr {
	return Expr{
		Line:   uint(node.GetSymbol().GetLine()),
		Column: uint(node.GetSymbol().GetColumn()),
		Text:   node.GetText(),
	}
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(content string) HelloSpec {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewHelloLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser1 := parser.NewHelloParser(tokens)
	v := parser1.Hello().Accept(&HelloVisitor{})
	if ret, ok := v.(HelloSpec); ok {
		return ret
	}
	return HelloSpec{}
}
