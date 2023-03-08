// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by HelloParser.
type HelloVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HelloParser#hello.
	VisitHello(ctx *HelloContext) interface{}

	// Visit a parse tree produced by HelloParser#spec.
	VisitSpec(ctx *SpecContext) interface{}

}