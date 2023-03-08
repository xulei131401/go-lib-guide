// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseHelloVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHelloVisitor) VisitHello(ctx *HelloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHelloVisitor) VisitSpec(ctx *SpecContext) interface{} {
	return v.VisitChildren(ctx)
}
