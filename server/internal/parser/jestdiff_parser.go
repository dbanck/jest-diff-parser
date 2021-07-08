// Code generated from JestDiff.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JestDiff

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 59, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 19, 10, 3, 12, 3, 14, 3, 22, 11, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 28, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 7, 5, 38, 10, 5, 12, 5, 14, 5, 41, 11, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 47, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 57, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 64, 2, 12, 3, 2,
	2, 2, 4, 27, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 56,
	3, 2, 2, 2, 12, 13, 5, 10, 6, 2, 13, 3, 3, 2, 2, 2, 14, 15, 7, 3, 2, 2,
	15, 20, 5, 6, 4, 2, 16, 17, 7, 4, 2, 2, 17, 19, 5, 6, 4, 2, 18, 16, 3,
	2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21,
	23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7, 5, 2, 2, 24, 28, 3, 2, 2,
	2, 25, 26, 7, 3, 2, 2, 26, 28, 7, 5, 2, 2, 27, 14, 3, 2, 2, 2, 27, 25,
	3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 13, 2, 2, 30, 31, 7, 6, 2, 2,
	31, 32, 5, 10, 6, 2, 32, 7, 3, 2, 2, 2, 33, 34, 7, 7, 2, 2, 34, 39, 5,
	10, 6, 2, 35, 36, 7, 4, 2, 2, 36, 38, 5, 10, 6, 2, 37, 35, 3, 2, 2, 2,
	38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3,
	2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 7, 8, 2, 2, 43, 47, 3, 2, 2, 2, 44,
	45, 7, 7, 2, 2, 45, 47, 7, 8, 2, 2, 46, 33, 3, 2, 2, 2, 46, 44, 3, 2, 2,
	2, 47, 9, 3, 2, 2, 2, 48, 57, 7, 13, 2, 2, 49, 57, 7, 14, 2, 2, 50, 57,
	5, 4, 3, 2, 51, 57, 5, 8, 5, 2, 52, 57, 7, 9, 2, 2, 53, 57, 7, 10, 2, 2,
	54, 57, 7, 11, 2, 2, 55, 57, 7, 12, 2, 2, 56, 48, 3, 2, 2, 2, 56, 49, 3,
	2, 2, 2, 56, 50, 3, 2, 2, 2, 56, 51, 3, 2, 2, 2, 56, 52, 3, 2, 2, 2, 56,
	53, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 11, 3, 2, 2,
	2, 7, 20, 27, 39, 46, 56,
}
var literalNames = []string{
	"", "'Object {'", "','", "'}'", "':'", "'Array ['", "']'", "'true'", "'false'",
	"'null'", "'undefined'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var ruleNames = []string{
	"jestDiff", "obj", "pair", "arr", "value",
}

type JestDiffParser struct {
	*antlr.BaseParser
}

// NewJestDiffParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *JestDiffParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJestDiffParser(input antlr.TokenStream) *JestDiffParser {
	this := new(JestDiffParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JestDiff.g4"

	return this
}

// JestDiffParser tokens.
const (
	JestDiffParserEOF    = antlr.TokenEOF
	JestDiffParserT__0   = 1
	JestDiffParserT__1   = 2
	JestDiffParserT__2   = 3
	JestDiffParserT__3   = 4
	JestDiffParserT__4   = 5
	JestDiffParserT__5   = 6
	JestDiffParserT__6   = 7
	JestDiffParserT__7   = 8
	JestDiffParserT__8   = 9
	JestDiffParserT__9   = 10
	JestDiffParserSTRING = 11
	JestDiffParserNUMBER = 12
	JestDiffParserWS     = 13
)

// JestDiffParser rules.
const (
	JestDiffParserRULE_jestDiff = 0
	JestDiffParserRULE_obj      = 1
	JestDiffParserRULE_pair     = 2
	JestDiffParserRULE_arr      = 3
	JestDiffParserRULE_value    = 4
)

// IJestDiffContext is an interface to support dynamic dispatch.
type IJestDiffContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJestDiffContext differentiates from other interfaces.
	IsJestDiffContext()
}

type JestDiffContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJestDiffContext() *JestDiffContext {
	var p = new(JestDiffContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JestDiffParserRULE_jestDiff
	return p
}

func (*JestDiffContext) IsJestDiffContext() {}

func NewJestDiffContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JestDiffContext {
	var p = new(JestDiffContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JestDiffParserRULE_jestDiff

	return p
}

func (s *JestDiffContext) GetParser() antlr.Parser { return s.parser }

func (s *JestDiffContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *JestDiffContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JestDiffContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JestDiffContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.EnterJestDiff(s)
	}
}

func (s *JestDiffContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.ExitJestDiff(s)
	}
}

func (p *JestDiffParser) JestDiff() (localctx IJestDiffContext) {
	localctx = NewJestDiffContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JestDiffParserRULE_jestDiff)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Value()
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JestDiffParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JestDiffParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *JestDiffParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JestDiffParserRULE_obj)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Match(JestDiffParserT__0)
		}
		{
			p.SetState(13)
			p.Pair()
		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JestDiffParserT__1 {
			{
				p.SetState(14)
				p.Match(JestDiffParserT__1)
			}
			{
				p.SetState(15)
				p.Pair()
			}

			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(21)
			p.Match(JestDiffParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(JestDiffParserT__0)
		}
		{
			p.SetState(24)
			p.Match(JestDiffParserT__2)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JestDiffParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JestDiffParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(JestDiffParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *JestDiffParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JestDiffParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(JestDiffParserSTRING)
	}
	{
		p.SetState(28)
		p.Match(JestDiffParserT__3)
	}
	{
		p.SetState(29)
		p.Value()
	}

	return localctx
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JestDiffParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JestDiffParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.ExitArr(s)
	}
}

func (p *JestDiffParser) Arr() (localctx IArrContext) {
	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JestDiffParserRULE_arr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(JestDiffParserT__4)
		}
		{
			p.SetState(32)
			p.Value()
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JestDiffParserT__1 {
			{
				p.SetState(33)
				p.Match(JestDiffParserT__1)
			}
			{
				p.SetState(34)
				p.Value()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(40)
			p.Match(JestDiffParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(JestDiffParserT__4)
		}
		{
			p.SetState(43)
			p.Match(JestDiffParserT__5)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JestDiffParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JestDiffParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JestDiffParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JestDiffParserNUMBER, 0)
}

func (s *ValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Arr() IArrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JestDiffListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *JestDiffParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JestDiffParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JestDiffParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(JestDiffParserSTRING)
		}

	case JestDiffParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(JestDiffParserNUMBER)
		}

	case JestDiffParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Obj()
		}

	case JestDiffParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Arr()
		}

	case JestDiffParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(JestDiffParserT__6)
		}

	case JestDiffParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.Match(JestDiffParserT__7)
		}

	case JestDiffParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(52)
			p.Match(JestDiffParserT__8)
		}

	case JestDiffParserT__9:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(53)
			p.Match(JestDiffParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
