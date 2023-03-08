// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hello

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type HelloParser struct {
	*antlr.BaseParser
}

var helloParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func helloParserInit() {
  staticData := &helloParserStaticData
  staticData.symbolicNames = []string{
    "", "ID", "WS", "LINE_COMMENT", "ERRCHAR",
  }
  staticData.ruleNames = []string{
    "hello", "spec",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 4, 14, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 5, 0, 6, 8, 0, 10, 0, 12, 0, 
	9, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0, 12, 0, 7, 1, 0, 0, 
	0, 2, 10, 1, 0, 0, 0, 4, 6, 3, 2, 1, 0, 5, 4, 1, 0, 0, 0, 6, 9, 1, 0, 0, 
	0, 7, 5, 1, 0, 0, 0, 7, 8, 1, 0, 0, 0, 8, 1, 1, 0, 0, 0, 9, 7, 1, 0, 0, 
	0, 10, 11, 5, 1, 0, 0, 11, 12, 5, 1, 0, 0, 12, 3, 1, 0, 0, 0, 1, 7,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// HelloParserInit initializes any static state used to implement HelloParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHelloParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HelloParserInit() {
  staticData := &helloParserStaticData
  staticData.once.Do(helloParserInit)
}

// NewHelloParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHelloParser(input antlr.TokenStream) *HelloParser {
	HelloParserInit()
	this := new(HelloParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &helloParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hello.g4"

	return this
}


// HelloParser tokens.
const (
	HelloParserEOF = antlr.TokenEOF
	HelloParserID = 1
	HelloParserWS = 2
	HelloParserLINE_COMMENT = 3
	HelloParserERRCHAR = 4
)

// HelloParser rules.
const (
	HelloParserRULE_hello = 0
	HelloParserRULE_spec = 1
)

// IHelloContext is an interface to support dynamic dispatch.
type IHelloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHelloContext differentiates from other interfaces.
	IsHelloContext()
}

type HelloContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelloContext() *HelloContext {
	var p = new(HelloContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HelloParserRULE_hello
	return p
}

func (*HelloContext) IsHelloContext() {}

func NewHelloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelloContext {
	var p = new(HelloContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HelloParserRULE_hello

	return p
}

func (s *HelloContext) GetParser() antlr.Parser { return s.parser }

func (s *HelloContext) AllSpec() []ISpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecContext); ok {
			len++
		}
	}

	tst := make([]ISpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecContext); ok {
			tst[i] = t.(ISpecContext)
			i++
		}
	}

	return tst
}

func (s *HelloContext) Spec(i int) ISpecContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *HelloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *HelloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.EnterHello(s)
	}
}

func (s *HelloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.ExitHello(s)
	}
}

func (s *HelloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HelloVisitor:
		return t.VisitHello(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *HelloParser) Hello() (localctx IHelloContext) {
	this := p
	_ = this

	localctx = NewHelloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HelloParserRULE_hello)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == HelloParserID {
		{
			p.SetState(4)
			p.Spec()
		}


		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HelloParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HelloParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(HelloParserID)
}

func (s *SpecContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(HelloParserID, i)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HelloListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HelloVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *HelloParser) Spec() (localctx ISpecContext) {
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HelloParserRULE_spec)

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
		p.Match(HelloParserID)
	}
	{
		p.SetState(11)
		p.Match(HelloParserID)
	}



	return localctx
}


