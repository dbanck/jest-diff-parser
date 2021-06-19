// Code generated from JestDiff.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JestDiff

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJestDiffListener is a complete listener for a parse tree produced by JestDiffParser.
type BaseJestDiffListener struct{}

var _ JestDiffListener = &BaseJestDiffListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJestDiffListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJestDiffListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJestDiffListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJestDiffListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseJestDiffListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseJestDiffListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseJestDiffListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseJestDiffListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJestDiffListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJestDiffListener) ExitPair(ctx *PairContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseJestDiffListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseJestDiffListener) ExitArr(ctx *ArrContext) {}

// EnterValue is called when production value is entered.
func (s *BaseJestDiffListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseJestDiffListener) ExitValue(ctx *ValueContext) {}
