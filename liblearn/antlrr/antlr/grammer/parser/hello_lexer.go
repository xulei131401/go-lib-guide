// Code generated from /Users/xulei/jungle/golangworkspace/holy-cmd/antlr/grammer/Hello.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type HelloLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var hellolexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func hellolexerLexerInit() {
  staticData := &hellolexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.symbolicNames = []string{
    "", "ID", "WS", "LINE_COMMENT", "ERRCHAR",
  }
  staticData.ruleNames = []string{
    "ID", "WS", "LINE_COMMENT", "ERRCHAR", "Letter",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 4, 45, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 1, 4, 1, 19, 
	8, 1, 11, 1, 12, 1, 20, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 29, 8, 
	2, 10, 2, 12, 2, 32, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 
	4, 1, 4, 1, 4, 3, 4, 44, 8, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 0, 1, 
	0, 6, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 10, 10, 13, 13, 4, 0, 36, 36, 
	65, 90, 95, 95, 97, 122, 2, 0, 0, 127, 55296, 56319, 1, 0, 55296, 56319, 
	1, 0, 56320, 57343, 48, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 
	0, 0, 0, 7, 1, 0, 0, 0, 1, 14, 1, 0, 0, 0, 3, 18, 1, 0, 0, 0, 5, 24, 1, 
	0, 0, 0, 7, 35, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11, 13, 3, 9, 4, 0, 12, 
	11, 1, 0, 0, 0, 13, 16, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 
	0, 15, 2, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 17, 19, 7, 0, 0, 0, 18, 17, 1, 
	0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 
	22, 1, 0, 0, 0, 22, 23, 6, 1, 0, 0, 23, 4, 1, 0, 0, 0, 24, 25, 5, 47, 0, 
	0, 25, 26, 5, 47, 0, 0, 26, 30, 1, 0, 0, 0, 27, 29, 8, 1, 0, 0, 28, 27, 
	1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 
	31, 33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 6, 2, 0, 0, 34, 6, 1, 0, 
	0, 0, 35, 36, 9, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 6, 3, 0, 0, 38, 8, 
	1, 0, 0, 0, 39, 44, 7, 2, 0, 0, 40, 44, 8, 3, 0, 0, 41, 42, 7, 4, 0, 0, 
	42, 44, 7, 5, 0, 0, 43, 39, 1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 43, 41, 1, 
	0, 0, 0, 44, 10, 1, 0, 0, 0, 5, 0, 14, 20, 30, 43, 1, 0, 1, 0,
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

// HelloLexerInit initializes any static state used to implement HelloLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHelloLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HelloLexerInit() {
  staticData := &hellolexerLexerStaticData
  staticData.once.Do(hellolexerLexerInit)
}

// NewHelloLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHelloLexer(input antlr.CharStream) *HelloLexer {
  HelloLexerInit()
	l := new(HelloLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &hellolexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Hello.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HelloLexer tokens.
const (
	HelloLexerID = 1
	HelloLexerWS = 2
	HelloLexerLINE_COMMENT = 3
	HelloLexerERRCHAR = 4
)

