// Code generated from kochanowski.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kochanowski

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type kochanowskiParser struct {
	*antlr.BaseParser
}

var KochanowskiParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kochanowskiParserInit() {
	staticData := &KochanowskiParserStaticData
	staticData.LiteralNames = []string{
		"", "'Zdefiniuj'", "'zmienn\\u0105'", "'ca\\u0142kowit\\u0105'", "'o nazwie'",
		"'o warto\\u015Bci'", "'.'", "'Przypisz'", "'Wypisz'", "'Wczytaj'",
		"'z'", "'wi\\u0119ksze ni\\u017C'", "'wi\\u0119ksze lub r\\u00F3wne'",
		"'r\\u00F3wne'", "'mniejsze lub r\\u00F3wne'", "'mniejsze ni\\u017C'",
		"'r\\u00F3\\u017Cne od'", "'modulo'", "'and'", "'or'", "'xor'", "'plus'",
		"'minus'", "'razy'", "'podzieli\\u0107 przez'", "'do pot\\u0119gi'",
		"'nie'", "'wpierw'", "'policz'",
	}
	staticData.SymbolicNames = []string{
		"", "DEFINE", "VARIABLE", "INT32", "NAMED", "WITH_VALUE", "DOT", "ASSIGN",
		"PRINT_WORD", "READ_WORD", "FROM", "GREATER", "GREATEREQUAL", "EQUAL",
		"LESSEQUAL", "LESS", "NOTEQUAL", "MODULO", "AND", "OR", "XOR", "PLUS",
		"MINUS", "TIMES", "DIVIDE", "POWER", "NOT", "FIRST", "CALCULATE", "INTEGER",
		"DECIMAL", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"body", "statement", "var_create", "type", "var_assign", "read", "print",
		"expr", "expr_compare", "expr_mod", "expr_bit", "expr_add", "expr_mult",
		"expr_power", "expr_paren", "unary", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 175, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 45, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 106, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 113, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 128, 8, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 139, 8, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 150, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 157, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 164, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 171, 8, 15, 1, 16, 1, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 1, 1, 0, 29, 31, 180, 0,
		37, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0,
		8, 66, 1, 0, 0, 0, 10, 71, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14, 79, 1, 0,
		0, 0, 16, 105, 1, 0, 0, 0, 18, 112, 1, 0, 0, 0, 20, 127, 1, 0, 0, 0, 22,
		138, 1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 156, 1, 0, 0, 0, 28, 163, 1,
		0, 0, 0, 30, 170, 1, 0, 0, 0, 32, 172, 1, 0, 0, 0, 34, 36, 3, 2, 1, 0,
		35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1,
		0, 0, 0, 38, 1, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 45, 3, 4, 2, 0, 41,
		45, 3, 8, 4, 0, 42, 45, 3, 12, 6, 0, 43, 45, 3, 10, 5, 0, 44, 40, 1, 0,
		0, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 3,
		1, 0, 0, 0, 46, 47, 5, 1, 0, 0, 47, 48, 5, 2, 0, 0, 48, 49, 3, 6, 3, 0,
		49, 50, 5, 4, 0, 0, 50, 51, 5, 31, 0, 0, 51, 52, 5, 6, 0, 0, 52, 63, 1,
		0, 0, 0, 53, 54, 5, 1, 0, 0, 54, 55, 5, 2, 0, 0, 55, 56, 3, 6, 3, 0, 56,
		57, 5, 4, 0, 0, 57, 58, 5, 31, 0, 0, 58, 59, 5, 5, 0, 0, 59, 60, 3, 14,
		7, 0, 60, 61, 5, 6, 0, 0, 61, 63, 1, 0, 0, 0, 62, 46, 1, 0, 0, 0, 62, 53,
		1, 0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 65, 5, 3, 0, 0, 65, 7, 1, 0, 0, 0, 66,
		67, 5, 7, 0, 0, 67, 68, 5, 31, 0, 0, 68, 69, 3, 14, 7, 0, 69, 70, 5, 6,
		0, 0, 70, 9, 1, 0, 0, 0, 71, 72, 5, 9, 0, 0, 72, 73, 5, 31, 0, 0, 73, 74,
		5, 6, 0, 0, 74, 11, 1, 0, 0, 0, 75, 76, 5, 8, 0, 0, 76, 77, 3, 14, 7, 0,
		77, 78, 5, 6, 0, 0, 78, 13, 1, 0, 0, 0, 79, 80, 3, 16, 8, 0, 80, 15, 1,
		0, 0, 0, 81, 106, 3, 18, 9, 0, 82, 83, 3, 18, 9, 0, 83, 84, 5, 11, 0, 0,
		84, 85, 3, 16, 8, 0, 85, 106, 1, 0, 0, 0, 86, 87, 3, 18, 9, 0, 87, 88,
		5, 12, 0, 0, 88, 89, 3, 16, 8, 0, 89, 106, 1, 0, 0, 0, 90, 91, 3, 18, 9,
		0, 91, 92, 5, 13, 0, 0, 92, 93, 3, 16, 8, 0, 93, 106, 1, 0, 0, 0, 94, 95,
		3, 18, 9, 0, 95, 96, 5, 14, 0, 0, 96, 97, 3, 16, 8, 0, 97, 106, 1, 0, 0,
		0, 98, 99, 3, 18, 9, 0, 99, 100, 5, 15, 0, 0, 100, 101, 3, 16, 8, 0, 101,
		106, 1, 0, 0, 0, 102, 103, 3, 18, 9, 0, 103, 104, 5, 16, 0, 0, 104, 106,
		1, 0, 0, 0, 105, 81, 1, 0, 0, 0, 105, 82, 1, 0, 0, 0, 105, 86, 1, 0, 0,
		0, 105, 90, 1, 0, 0, 0, 105, 94, 1, 0, 0, 0, 105, 98, 1, 0, 0, 0, 105,
		102, 1, 0, 0, 0, 106, 17, 1, 0, 0, 0, 107, 113, 3, 20, 10, 0, 108, 109,
		3, 20, 10, 0, 109, 110, 5, 17, 0, 0, 110, 111, 3, 18, 9, 0, 111, 113, 1,
		0, 0, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 113, 19, 1, 0, 0,
		0, 114, 128, 3, 22, 11, 0, 115, 116, 3, 22, 11, 0, 116, 117, 5, 18, 0,
		0, 117, 118, 3, 20, 10, 0, 118, 128, 1, 0, 0, 0, 119, 120, 3, 22, 11, 0,
		120, 121, 5, 19, 0, 0, 121, 122, 3, 20, 10, 0, 122, 128, 1, 0, 0, 0, 123,
		124, 3, 22, 11, 0, 124, 125, 5, 20, 0, 0, 125, 126, 3, 20, 10, 0, 126,
		128, 1, 0, 0, 0, 127, 114, 1, 0, 0, 0, 127, 115, 1, 0, 0, 0, 127, 119,
		1, 0, 0, 0, 127, 123, 1, 0, 0, 0, 128, 21, 1, 0, 0, 0, 129, 139, 3, 24,
		12, 0, 130, 131, 3, 24, 12, 0, 131, 132, 5, 21, 0, 0, 132, 133, 3, 22,
		11, 0, 133, 139, 1, 0, 0, 0, 134, 135, 3, 24, 12, 0, 135, 136, 5, 22, 0,
		0, 136, 137, 3, 22, 11, 0, 137, 139, 1, 0, 0, 0, 138, 129, 1, 0, 0, 0,
		138, 130, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 139, 23, 1, 0, 0, 0, 140, 150,
		3, 26, 13, 0, 141, 142, 3, 26, 13, 0, 142, 143, 5, 23, 0, 0, 143, 144,
		3, 24, 12, 0, 144, 150, 1, 0, 0, 0, 145, 146, 3, 26, 13, 0, 146, 147, 5,
		24, 0, 0, 147, 148, 3, 24, 12, 0, 148, 150, 1, 0, 0, 0, 149, 140, 1, 0,
		0, 0, 149, 141, 1, 0, 0, 0, 149, 145, 1, 0, 0, 0, 150, 25, 1, 0, 0, 0,
		151, 157, 3, 28, 14, 0, 152, 153, 3, 28, 14, 0, 153, 154, 5, 25, 0, 0,
		154, 155, 3, 26, 13, 0, 155, 157, 1, 0, 0, 0, 156, 151, 1, 0, 0, 0, 156,
		152, 1, 0, 0, 0, 157, 27, 1, 0, 0, 0, 158, 164, 3, 30, 15, 0, 159, 160,
		5, 27, 0, 0, 160, 161, 3, 14, 7, 0, 161, 162, 5, 28, 0, 0, 162, 164, 1,
		0, 0, 0, 163, 158, 1, 0, 0, 0, 163, 159, 1, 0, 0, 0, 164, 29, 1, 0, 0,
		0, 165, 171, 3, 32, 16, 0, 166, 167, 5, 22, 0, 0, 167, 171, 3, 14, 7, 0,
		168, 169, 5, 26, 0, 0, 169, 171, 3, 14, 7, 0, 170, 165, 1, 0, 0, 0, 170,
		166, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 31, 1, 0, 0, 0, 172, 173, 7,
		0, 0, 0, 173, 33, 1, 0, 0, 0, 11, 37, 44, 62, 105, 112, 127, 138, 149,
		156, 163, 170,
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

// kochanowskiParserInit initializes any static state used to implement kochanowskiParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewkochanowskiParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KochanowskiParserInit() {
	staticData := &KochanowskiParserStaticData
	staticData.once.Do(kochanowskiParserInit)
}

// NewkochanowskiParser produces a new parser instance for the optional input antlr.TokenStream.
func NewkochanowskiParser(input antlr.TokenStream) *kochanowskiParser {
	KochanowskiParserInit()
	this := new(kochanowskiParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KochanowskiParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "kochanowski.g4"

	return this
}

// kochanowskiParser tokens.
const (
	kochanowskiParserEOF          = antlr.TokenEOF
	kochanowskiParserDEFINE       = 1
	kochanowskiParserVARIABLE     = 2
	kochanowskiParserINT32        = 3
	kochanowskiParserNAMED        = 4
	kochanowskiParserWITH_VALUE   = 5
	kochanowskiParserDOT          = 6
	kochanowskiParserASSIGN       = 7
	kochanowskiParserPRINT_WORD   = 8
	kochanowskiParserREAD_WORD    = 9
	kochanowskiParserFROM         = 10
	kochanowskiParserGREATER      = 11
	kochanowskiParserGREATEREQUAL = 12
	kochanowskiParserEQUAL        = 13
	kochanowskiParserLESSEQUAL    = 14
	kochanowskiParserLESS         = 15
	kochanowskiParserNOTEQUAL     = 16
	kochanowskiParserMODULO       = 17
	kochanowskiParserAND          = 18
	kochanowskiParserOR           = 19
	kochanowskiParserXOR          = 20
	kochanowskiParserPLUS         = 21
	kochanowskiParserMINUS        = 22
	kochanowskiParserTIMES        = 23
	kochanowskiParserDIVIDE       = 24
	kochanowskiParserPOWER        = 25
	kochanowskiParserNOT          = 26
	kochanowskiParserFIRST        = 27
	kochanowskiParserCALCULATE    = 28
	kochanowskiParserINTEGER      = 29
	kochanowskiParserDECIMAL      = 30
	kochanowskiParserID           = 31
	kochanowskiParserWS           = 32
)

// kochanowskiParser rules.
const (
	kochanowskiParserRULE_body         = 0
	kochanowskiParserRULE_statement    = 1
	kochanowskiParserRULE_var_create   = 2
	kochanowskiParserRULE_type         = 3
	kochanowskiParserRULE_var_assign   = 4
	kochanowskiParserRULE_read         = 5
	kochanowskiParserRULE_print        = 6
	kochanowskiParserRULE_expr         = 7
	kochanowskiParserRULE_expr_compare = 8
	kochanowskiParserRULE_expr_mod     = 9
	kochanowskiParserRULE_expr_bit     = 10
	kochanowskiParserRULE_expr_add     = 11
	kochanowskiParserRULE_expr_mult    = 12
	kochanowskiParserRULE_expr_power   = 13
	kochanowskiParserRULE_expr_paren   = 14
	kochanowskiParserRULE_unary        = 15
	kochanowskiParserRULE_value        = 16
)

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *kochanowskiParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, kochanowskiParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&898) != 0 {
		{
			p.SetState(34)
			p.Statement()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_create() IVar_createContext
	Var_assign() IVar_assignContext
	Print_() IPrintContext
	Read() IReadContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Var_create() IVar_createContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_createContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_createContext)
}

func (s *StatementContext) Var_assign() IVar_assignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *kochanowskiParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, kochanowskiParserRULE_statement)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserDEFINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Var_create()
		}

	case kochanowskiParserASSIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Var_assign()
		}

	case kochanowskiParserPRINT_WORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Print_()
		}

	case kochanowskiParserREAD_WORD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.Read()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_createContext is an interface to support dynamic dispatch.
type IVar_createContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	Type_() ITypeContext
	NAMED() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOT() antlr.TerminalNode
	WITH_VALUE() antlr.TerminalNode
	Expr() IExprContext

	// IsVar_createContext differentiates from other interfaces.
	IsVar_createContext()
}

type Var_createContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_createContext() *Var_createContext {
	var p = new(Var_createContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_var_create
	return p
}

func InitEmptyVar_createContext(p *Var_createContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_var_create
}

func (*Var_createContext) IsVar_createContext() {}

func NewVar_createContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_createContext {
	var p = new(Var_createContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_var_create

	return p
}

func (s *Var_createContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_createContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDEFINE, 0)
}

func (s *Var_createContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserVARIABLE, 0)
}

func (s *Var_createContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_createContext) NAMED() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserNAMED, 0)
}

func (s *Var_createContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Var_createContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Var_createContext) WITH_VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserWITH_VALUE, 0)
}

func (s *Var_createContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_createContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_createContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_createContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterVar_create(s)
	}
}

func (s *Var_createContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitVar_create(s)
	}
}

func (p *kochanowskiParser) Var_create() (localctx IVar_createContext) {
	localctx = NewVar_createContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, kochanowskiParserRULE_var_create)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Type_()
		}
		{
			p.SetState(49)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Type_()
		}
		{
			p.SetState(56)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Expr()
		}
		{
			p.SetState(60)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT32, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *kochanowskiParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kochanowskiParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(kochanowskiParserINT32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_assignContext is an interface to support dynamic dispatch.
type IVar_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode
	Expr() IExprContext
	DOT() antlr.TerminalNode

	// IsVar_assignContext differentiates from other interfaces.
	IsVar_assignContext()
}

type Var_assignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignContext() *Var_assignContext {
	var p = new(Var_assignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_var_assign
	return p
}

func InitEmptyVar_assignContext(p *Var_assignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_var_assign
}

func (*Var_assignContext) IsVar_assignContext() {}

func NewVar_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignContext {
	var p = new(Var_assignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_var_assign

	return p
}

func (s *Var_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserASSIGN, 0)
}

func (s *Var_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Var_assignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_assignContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Var_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterVar_assign(s)
	}
}

func (s *Var_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitVar_assign(s)
	}
}

func (p *kochanowskiParser) Var_assign() (localctx IVar_assignContext) {
	localctx = NewVar_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kochanowskiParserRULE_var_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Expr()
	}
	{
		p.SetState(69)
		p.Match(kochanowskiParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ_WORD() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_read
	return p
}

func InitEmptyReadContext(p *ReadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_read
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) READ_WORD() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserREAD_WORD, 0)
}

func (s *ReadContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *ReadContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *kochanowskiParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kochanowskiParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(kochanowskiParserREAD_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(kochanowskiParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT_WORD() antlr.TerminalNode
	Expr() IExprContext
	DOT() antlr.TerminalNode

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT_WORD() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserPRINT_WORD, 0)
}

func (s *PrintContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *kochanowskiParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, kochanowskiParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(kochanowskiParserPRINT_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Expr()
	}
	{
		p.SetState(77)
		p.Match(kochanowskiParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_compare() IExpr_compareContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Expr_compare() IExpr_compareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_compareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_compareContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *kochanowskiParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kochanowskiParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Expr_compare()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_compareContext is an interface to support dynamic dispatch.
type IExpr_compareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_mod() IExpr_modContext
	GREATER() antlr.TerminalNode
	Expr_compare() IExpr_compareContext
	GREATEREQUAL() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	LESSEQUAL() antlr.TerminalNode
	LESS() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode

	// IsExpr_compareContext differentiates from other interfaces.
	IsExpr_compareContext()
}

type Expr_compareContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_compareContext() *Expr_compareContext {
	var p = new(Expr_compareContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_compare
	return p
}

func InitEmptyExpr_compareContext(p *Expr_compareContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_compare
}

func (*Expr_compareContext) IsExpr_compareContext() {}

func NewExpr_compareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_compareContext {
	var p = new(Expr_compareContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_compare

	return p
}

func (s *Expr_compareContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_compareContext) Expr_mod() IExpr_modContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_modContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_modContext)
}

func (s *Expr_compareContext) GREATER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserGREATER, 0)
}

func (s *Expr_compareContext) Expr_compare() IExpr_compareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_compareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_compareContext)
}

func (s *Expr_compareContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserGREATEREQUAL, 0)
}

func (s *Expr_compareContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserEQUAL, 0)
}

func (s *Expr_compareContext) LESSEQUAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserLESSEQUAL, 0)
}

func (s *Expr_compareContext) LESS() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserLESS, 0)
}

func (s *Expr_compareContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserNOTEQUAL, 0)
}

func (s *Expr_compareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_compareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_compareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_compare(s)
	}
}

func (s *Expr_compareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_compare(s)
	}
}

func (p *kochanowskiParser) Expr_compare() (localctx IExpr_compareContext) {
	localctx = NewExpr_compareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kochanowskiParserRULE_expr_compare)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Expr_mod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Expr_mod()
		}
		{
			p.SetState(83)
			p.Match(kochanowskiParserGREATER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Expr_compare()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Expr_mod()
		}
		{
			p.SetState(87)
			p.Match(kochanowskiParserGREATEREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Expr_compare()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.Expr_mod()
		}
		{
			p.SetState(91)
			p.Match(kochanowskiParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Expr_compare()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Expr_mod()
		}
		{
			p.SetState(95)
			p.Match(kochanowskiParserLESSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Expr_compare()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.Expr_mod()
		}
		{
			p.SetState(99)
			p.Match(kochanowskiParserLESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Expr_compare()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.Expr_mod()
		}
		{
			p.SetState(103)
			p.Match(kochanowskiParserNOTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_modContext is an interface to support dynamic dispatch.
type IExpr_modContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_bit() IExpr_bitContext
	MODULO() antlr.TerminalNode
	Expr_mod() IExpr_modContext

	// IsExpr_modContext differentiates from other interfaces.
	IsExpr_modContext()
}

type Expr_modContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_modContext() *Expr_modContext {
	var p = new(Expr_modContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_mod
	return p
}

func InitEmptyExpr_modContext(p *Expr_modContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_mod
}

func (*Expr_modContext) IsExpr_modContext() {}

func NewExpr_modContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_modContext {
	var p = new(Expr_modContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_mod

	return p
}

func (s *Expr_modContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_modContext) Expr_bit() IExpr_bitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_bitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_bitContext)
}

func (s *Expr_modContext) MODULO() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserMODULO, 0)
}

func (s *Expr_modContext) Expr_mod() IExpr_modContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_modContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_modContext)
}

func (s *Expr_modContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_modContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_modContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_mod(s)
	}
}

func (s *Expr_modContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_mod(s)
	}
}

func (p *kochanowskiParser) Expr_mod() (localctx IExpr_modContext) {
	localctx = NewExpr_modContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kochanowskiParserRULE_expr_mod)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Expr_bit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Expr_bit()
		}
		{
			p.SetState(109)
			p.Match(kochanowskiParserMODULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Expr_mod()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_bitContext is an interface to support dynamic dispatch.
type IExpr_bitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_add() IExpr_addContext
	AND() antlr.TerminalNode
	Expr_bit() IExpr_bitContext
	OR() antlr.TerminalNode
	XOR() antlr.TerminalNode

	// IsExpr_bitContext differentiates from other interfaces.
	IsExpr_bitContext()
}

type Expr_bitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_bitContext() *Expr_bitContext {
	var p = new(Expr_bitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_bit
	return p
}

func InitEmptyExpr_bitContext(p *Expr_bitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_bit
}

func (*Expr_bitContext) IsExpr_bitContext() {}

func NewExpr_bitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_bitContext {
	var p = new(Expr_bitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_bit

	return p
}

func (s *Expr_bitContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_bitContext) Expr_add() IExpr_addContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_addContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_addContext)
}

func (s *Expr_bitContext) AND() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserAND, 0)
}

func (s *Expr_bitContext) Expr_bit() IExpr_bitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_bitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_bitContext)
}

func (s *Expr_bitContext) OR() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserOR, 0)
}

func (s *Expr_bitContext) XOR() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserXOR, 0)
}

func (s *Expr_bitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_bitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_bitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_bit(s)
	}
}

func (s *Expr_bitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_bit(s)
	}
}

func (p *kochanowskiParser) Expr_bit() (localctx IExpr_bitContext) {
	localctx = NewExpr_bitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, kochanowskiParserRULE_expr_bit)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Expr_add()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Expr_add()
		}
		{
			p.SetState(116)
			p.Match(kochanowskiParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Expr_bit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Expr_add()
		}
		{
			p.SetState(120)
			p.Match(kochanowskiParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Expr_bit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.Expr_add()
		}
		{
			p.SetState(124)
			p.Match(kochanowskiParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Expr_bit()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_addContext is an interface to support dynamic dispatch.
type IExpr_addContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_mult() IExpr_multContext
	PLUS() antlr.TerminalNode
	Expr_add() IExpr_addContext
	MINUS() antlr.TerminalNode

	// IsExpr_addContext differentiates from other interfaces.
	IsExpr_addContext()
}

type Expr_addContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_addContext() *Expr_addContext {
	var p = new(Expr_addContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_add
	return p
}

func InitEmptyExpr_addContext(p *Expr_addContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_add
}

func (*Expr_addContext) IsExpr_addContext() {}

func NewExpr_addContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_addContext {
	var p = new(Expr_addContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_add

	return p
}

func (s *Expr_addContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_addContext) Expr_mult() IExpr_multContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_multContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_multContext)
}

func (s *Expr_addContext) PLUS() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserPLUS, 0)
}

func (s *Expr_addContext) Expr_add() IExpr_addContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_addContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_addContext)
}

func (s *Expr_addContext) MINUS() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserMINUS, 0)
}

func (s *Expr_addContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_addContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_addContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_add(s)
	}
}

func (s *Expr_addContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_add(s)
	}
}

func (p *kochanowskiParser) Expr_add() (localctx IExpr_addContext) {
	localctx = NewExpr_addContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, kochanowskiParserRULE_expr_add)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Expr_mult()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Expr_mult()
		}
		{
			p.SetState(131)
			p.Match(kochanowskiParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Expr_add()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Expr_mult()
		}
		{
			p.SetState(135)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Expr_add()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_multContext is an interface to support dynamic dispatch.
type IExpr_multContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_power() IExpr_powerContext
	TIMES() antlr.TerminalNode
	Expr_mult() IExpr_multContext
	DIVIDE() antlr.TerminalNode

	// IsExpr_multContext differentiates from other interfaces.
	IsExpr_multContext()
}

type Expr_multContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_multContext() *Expr_multContext {
	var p = new(Expr_multContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_mult
	return p
}

func InitEmptyExpr_multContext(p *Expr_multContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_mult
}

func (*Expr_multContext) IsExpr_multContext() {}

func NewExpr_multContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_multContext {
	var p = new(Expr_multContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_mult

	return p
}

func (s *Expr_multContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_multContext) Expr_power() IExpr_powerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_powerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_powerContext)
}

func (s *Expr_multContext) TIMES() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserTIMES, 0)
}

func (s *Expr_multContext) Expr_mult() IExpr_multContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_multContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_multContext)
}

func (s *Expr_multContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDIVIDE, 0)
}

func (s *Expr_multContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_multContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_multContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_mult(s)
	}
}

func (s *Expr_multContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_mult(s)
	}
}

func (p *kochanowskiParser) Expr_mult() (localctx IExpr_multContext) {
	localctx = NewExpr_multContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, kochanowskiParserRULE_expr_mult)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Expr_power()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Expr_power()
		}
		{
			p.SetState(142)
			p.Match(kochanowskiParserTIMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Expr_mult()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Expr_power()
		}
		{
			p.SetState(146)
			p.Match(kochanowskiParserDIVIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Expr_mult()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_powerContext is an interface to support dynamic dispatch.
type IExpr_powerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_paren() IExpr_parenContext
	POWER() antlr.TerminalNode
	Expr_power() IExpr_powerContext

	// IsExpr_powerContext differentiates from other interfaces.
	IsExpr_powerContext()
}

type Expr_powerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_powerContext() *Expr_powerContext {
	var p = new(Expr_powerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_power
	return p
}

func InitEmptyExpr_powerContext(p *Expr_powerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_power
}

func (*Expr_powerContext) IsExpr_powerContext() {}

func NewExpr_powerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_powerContext {
	var p = new(Expr_powerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_power

	return p
}

func (s *Expr_powerContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_powerContext) Expr_paren() IExpr_parenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_parenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_parenContext)
}

func (s *Expr_powerContext) POWER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserPOWER, 0)
}

func (s *Expr_powerContext) Expr_power() IExpr_powerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_powerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_powerContext)
}

func (s *Expr_powerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_powerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_powerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_power(s)
	}
}

func (s *Expr_powerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_power(s)
	}
}

func (p *kochanowskiParser) Expr_power() (localctx IExpr_powerContext) {
	localctx = NewExpr_powerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, kochanowskiParserRULE_expr_power)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Expr_paren()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Expr_paren()
		}
		{
			p.SetState(153)
			p.Match(kochanowskiParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Expr_power()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_parenContext is an interface to support dynamic dispatch.
type IExpr_parenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	FIRST() antlr.TerminalNode
	Expr() IExprContext
	CALCULATE() antlr.TerminalNode

	// IsExpr_parenContext differentiates from other interfaces.
	IsExpr_parenContext()
}

type Expr_parenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_parenContext() *Expr_parenContext {
	var p = new(Expr_parenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_paren
	return p
}

func InitEmptyExpr_parenContext(p *Expr_parenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_paren
}

func (*Expr_parenContext) IsExpr_parenContext() {}

func NewExpr_parenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_parenContext {
	var p = new(Expr_parenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_paren

	return p
}

func (s *Expr_parenContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_parenContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *Expr_parenContext) FIRST() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserFIRST, 0)
}

func (s *Expr_parenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_parenContext) CALCULATE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserCALCULATE, 0)
}

func (s *Expr_parenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_parenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_parenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_paren(s)
	}
}

func (s *Expr_parenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_paren(s)
	}
}

func (p *kochanowskiParser) Expr_paren() (localctx IExpr_parenContext) {
	localctx = NewExpr_parenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kochanowskiParserRULE_expr_paren)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserMINUS, kochanowskiParserNOT, kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Unary()
		}

	case kochanowskiParserFIRST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(kochanowskiParserFIRST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Expr()
		}
		{
			p.SetState(161)
			p.Match(kochanowskiParserCALCULATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	MINUS() antlr.TerminalNode
	Expr() IExprContext
	NOT() antlr.TerminalNode

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserMINUS, 0)
}

func (s *UnaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserNOT, 0)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (p *kochanowskiParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, kochanowskiParserRULE_unary)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Value()
		}

	case kochanowskiParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Expr()
		}

	case kochanowskiParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Match(kochanowskiParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINTEGER, 0)
}

func (s *ValueContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDECIMAL, 0)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *kochanowskiParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, kochanowskiParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
