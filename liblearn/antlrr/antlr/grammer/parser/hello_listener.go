// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HelloListener is a complete listener for a parse tree produced by HelloParser.
type HelloListener interface {
	antlr.ParseTreeListener

	// EnterHello is called when entering the hello production.
	EnterHello(c *HelloContext)

	// EnterSpec is called when entering the spec production.
	EnterSpec(c *SpecContext)

	// ExitHello is called when exiting the hello production.
	ExitHello(c *HelloContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)
}
