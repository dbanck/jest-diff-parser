package transformer

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dbanck/jest-diff-parser/internal/parser"
)

type transformListener struct {
	*parser.BaseJestDiffListener
	result strings.Builder
}

func (l *transformListener) EnterObj(ctx *parser.ObjContext) {
	l.result.WriteString("{")
}

func (l *transformListener) ExitObj(ctx *parser.ObjContext) {
	// TODO: check if last child and omit comma
	l.result.WriteString("},")
}

func (l *transformListener) EnterArr(ctx *parser.ArrContext) {
	l.result.WriteString("[")
}

func (l *transformListener) ExitArr(ctx *parser.ArrContext) {
	// TODO: check if last child and omit comma
	l.result.WriteString("],")
}

func (l *transformListener) EnterPair(ctx *parser.PairContext) {
	l.result.WriteString(ctx.STRING().GetText() + ":")
}

func (l *transformListener) ExitPair(ctx *parser.PairContext) {
	// TODO: check if last child and omit comma
	l.result.WriteString(",")
}

func (l *transformListener) EnterValue(ctx *parser.ValueContext) {
	if ctx.Arr() != nil || ctx.Obj() != nil {
		return
	}

	l.result.WriteString(ctx.GetText())
}

func Transform(source string) string {
	is := antlr.NewInputStream(source)

	// Create the Lexer
	lexer := parser.NewJestDiffLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJestDiffParser(stream)
	p.BuildParseTrees = true

	tree := p.JestDiff()
	result := transformListener{}
	antlr.ParseTreeWalkerDefault.Walk(&result, tree)

	return result.result.String()
}
