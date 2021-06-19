package main

import (
	// "fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dbanck/jest-diff-parser/internal/parser"
)

func main() {
	is, _ := antlr.NewFileStream("example.diff")

	// Create the Lexer
	lexer := parser.NewJestDiffLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJestDiffParser(stream)
	p.Obj()

}
