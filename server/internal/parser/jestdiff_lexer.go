// Code generated from JestDiff.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 146,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 7, 11, 84, 10, 11, 12, 11, 14, 11, 87, 11, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 5, 12, 94, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 5, 16, 107, 10, 16, 3, 16, 3,
	16, 3, 16, 6, 16, 112, 10, 16, 13, 16, 14, 16, 113, 5, 16, 116, 10, 16,
	3, 16, 5, 16, 119, 10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 124, 10, 17, 12,
	17, 14, 17, 127, 11, 17, 5, 17, 129, 10, 17, 3, 18, 3, 18, 5, 18, 133,
	10, 18, 3, 18, 3, 18, 3, 19, 5, 19, 138, 10, 19, 3, 19, 6, 19, 141, 10,
	19, 13, 19, 14, 19, 142, 3, 19, 3, 19, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 2, 25, 2, 27, 2, 29, 2,
	31, 13, 33, 2, 35, 2, 37, 14, 3, 2, 10, 10, 2, 36, 36, 49, 49, 94, 94,
	100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72,
	99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 3, 2, 50, 59, 3, 2, 51, 59, 4, 2,
	71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	151, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 48, 3, 2, 2, 2, 7, 50, 3, 2, 2, 2, 9,
	52, 3, 2, 2, 2, 11, 54, 3, 2, 2, 2, 13, 62, 3, 2, 2, 2, 15, 64, 3, 2, 2,
	2, 17, 69, 3, 2, 2, 2, 19, 75, 3, 2, 2, 2, 21, 80, 3, 2, 2, 2, 23, 90,
	3, 2, 2, 2, 25, 95, 3, 2, 2, 2, 27, 101, 3, 2, 2, 2, 29, 103, 3, 2, 2,
	2, 31, 106, 3, 2, 2, 2, 33, 128, 3, 2, 2, 2, 35, 130, 3, 2, 2, 2, 37, 137,
	3, 2, 2, 2, 39, 40, 7, 81, 2, 2, 40, 41, 7, 100, 2, 2, 41, 42, 7, 108,
	2, 2, 42, 43, 7, 103, 2, 2, 43, 44, 7, 101, 2, 2, 44, 45, 7, 118, 2, 2,
	45, 46, 7, 34, 2, 2, 46, 47, 7, 125, 2, 2, 47, 4, 3, 2, 2, 2, 48, 49, 7,
	46, 2, 2, 49, 6, 3, 2, 2, 2, 50, 51, 7, 127, 2, 2, 51, 8, 3, 2, 2, 2, 52,
	53, 7, 60, 2, 2, 53, 10, 3, 2, 2, 2, 54, 55, 7, 67, 2, 2, 55, 56, 7, 116,
	2, 2, 56, 57, 7, 116, 2, 2, 57, 58, 7, 99, 2, 2, 58, 59, 7, 123, 2, 2,
	59, 60, 7, 34, 2, 2, 60, 61, 7, 93, 2, 2, 61, 12, 3, 2, 2, 2, 62, 63, 7,
	95, 2, 2, 63, 14, 3, 2, 2, 2, 64, 65, 7, 118, 2, 2, 65, 66, 7, 116, 2,
	2, 66, 67, 7, 119, 2, 2, 67, 68, 7, 103, 2, 2, 68, 16, 3, 2, 2, 2, 69,
	70, 7, 104, 2, 2, 70, 71, 7, 99, 2, 2, 71, 72, 7, 110, 2, 2, 72, 73, 7,
	117, 2, 2, 73, 74, 7, 103, 2, 2, 74, 18, 3, 2, 2, 2, 75, 76, 7, 112, 2,
	2, 76, 77, 7, 119, 2, 2, 77, 78, 7, 110, 2, 2, 78, 79, 7, 110, 2, 2, 79,
	20, 3, 2, 2, 2, 80, 85, 7, 36, 2, 2, 81, 84, 5, 23, 12, 2, 82, 84, 5, 29,
	15, 2, 83, 81, 3, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85,
	83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2,
	2, 88, 89, 7, 36, 2, 2, 89, 22, 3, 2, 2, 2, 90, 93, 7, 94, 2, 2, 91, 94,
	9, 2, 2, 2, 92, 94, 5, 25, 13, 2, 93, 91, 3, 2, 2, 2, 93, 92, 3, 2, 2,
	2, 94, 24, 3, 2, 2, 2, 95, 96, 7, 119, 2, 2, 96, 97, 5, 27, 14, 2, 97,
	98, 5, 27, 14, 2, 98, 99, 5, 27, 14, 2, 99, 100, 5, 27, 14, 2, 100, 26,
	3, 2, 2, 2, 101, 102, 9, 3, 2, 2, 102, 28, 3, 2, 2, 2, 103, 104, 10, 4,
	2, 2, 104, 30, 3, 2, 2, 2, 105, 107, 7, 47, 2, 2, 106, 105, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 115, 5, 33, 17, 2, 109,
	111, 7, 48, 2, 2, 110, 112, 9, 5, 2, 2, 111, 110, 3, 2, 2, 2, 112, 113,
	3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2,
	2, 2, 115, 109, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2,
	117, 119, 5, 35, 18, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	32, 3, 2, 2, 2, 120, 129, 7, 50, 2, 2, 121, 125, 9, 6, 2, 2, 122, 124,
	9, 5, 2, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2,
	2, 2, 125, 126, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2,
	128, 120, 3, 2, 2, 2, 128, 121, 3, 2, 2, 2, 129, 34, 3, 2, 2, 2, 130, 132,
	9, 7, 2, 2, 131, 133, 9, 8, 2, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 5, 33, 17, 2, 135, 36, 3, 2, 2, 2,
	136, 138, 9, 8, 2, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	140, 3, 2, 2, 2, 139, 141, 9, 9, 2, 2, 140, 139, 3, 2, 2, 2, 141, 142,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2,
	2, 2, 144, 145, 8, 19, 2, 2, 145, 38, 3, 2, 2, 2, 15, 2, 83, 85, 93, 106,
	113, 115, 118, 125, 128, 132, 137, 142, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'Object {'", "','", "'}'", "':'", "'Array ['", "']'", "'true'", "'false'",
	"'null'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP",
	"WS",
}

type JestDiffLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewJestDiffLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *JestDiffLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJestDiffLexer(input antlr.CharStream) *JestDiffLexer {
	l := new(JestDiffLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "JestDiff.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JestDiffLexer tokens.
const (
	JestDiffLexerT__0   = 1
	JestDiffLexerT__1   = 2
	JestDiffLexerT__2   = 3
	JestDiffLexerT__3   = 4
	JestDiffLexerT__4   = 5
	JestDiffLexerT__5   = 6
	JestDiffLexerT__6   = 7
	JestDiffLexerT__7   = 8
	JestDiffLexerT__8   = 9
	JestDiffLexerSTRING = 10
	JestDiffLexerNUMBER = 11
	JestDiffLexerWS     = 12
)
