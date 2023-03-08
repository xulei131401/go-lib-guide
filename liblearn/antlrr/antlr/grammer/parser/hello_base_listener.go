// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHelloListener is a complete listener for a parse tree produced by HelloParser.
type BaseHelloListener struct{}

var _ HelloListener = &BaseHelloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHelloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHelloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHelloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHelloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHello is called when production hello is entered.
func (s *BaseHelloListener) EnterHello(ctx *HelloContext) {}

// ExitHello is called when production hello is exited.
func (s *BaseHelloListener) ExitHello(ctx *HelloContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseHelloListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseHelloListener) ExitSpec(ctx *SpecContext) {}
