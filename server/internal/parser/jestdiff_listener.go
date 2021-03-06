// Code generated from JestDiff.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JestDiff

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JestDiffListener is a complete listener for a parse tree produced by JestDiffParser.
type JestDiffListener interface {
	antlr.ParseTreeListener

	// EnterJestDiff is called when entering the jestDiff production.
	EnterJestDiff(c *JestDiffContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitJestDiff is called when exiting the jestDiff production.
	ExitJestDiff(c *JestDiffContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
