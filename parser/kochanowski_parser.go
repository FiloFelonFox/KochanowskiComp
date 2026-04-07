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
		"'z'", "'wi\\u0119ksze'", "'wi\\u0119ksze lub r\\u00F3wne'", "'r\\u00F3wne'",
		"'mniejsze lub r\\u00F3wne'", "'mniejsze ni\\u017C'", "'modulo'", "'and'",
		"'or'", "'xor'", "'plus'", "'minus'", "'razy'", "'podzieli\\u0107 przez'",
		"'do pot\\u0119gi'", "'nie'", "'wpierw'", "'policz'",
	}
	staticData.SymbolicNames = []string{
		"", "DEFINE", "VARIABLE", "INT32", "NAMED", "WITH_VALUE", "DOT", "ASSIGN",
		"PRINT_WORD", "READ_WORD", "FROM", "GREATER", "GREATEREQUAL", "EQUAL",
		"LESSEQUAL", "LESS", "MODULO", "AND", "OR", "XOR", "PLUS", "MINUS",
		"TIMES", "DIVIDE", "POWER", "NOT", "FIRST", "CALCULATE", "INTEGER",
		"DECIMAL", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"body", "statement", "var_create", "type", "var_assign", "read", "print",
		"expr", "expr_compare", "expr_mod", "expr_bit", "expr_add", "expr_mult",
		"expr_power", "expr_paren", "unary", "num", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 169, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 96, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 103, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 118, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 129, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12,
		140, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 147, 8, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 154, 8, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 161, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 167, 8,
		17, 1, 17, 0, 0, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 0, 1, 1, 0, 28, 29, 172, 0, 39, 1, 0, 0, 0, 2, 46, 1, 0,
		0, 0, 4, 48, 1, 0, 0, 0, 6, 57, 1, 0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 64,
		1, 0, 0, 0, 12, 68, 1, 0, 0, 0, 14, 72, 1, 0, 0, 0, 16, 95, 1, 0, 0, 0,
		18, 102, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22, 128, 1, 0, 0, 0, 24, 139,
		1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 153, 1, 0, 0, 0, 30, 160, 1, 0, 0,
		0, 32, 162, 1, 0, 0, 0, 34, 166, 1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36,
		1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0,
		40, 1, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 47, 3, 4, 2, 0, 43, 47, 3, 8,
		4, 0, 44, 47, 3, 12, 6, 0, 45, 47, 3, 10, 5, 0, 46, 42, 1, 0, 0, 0, 46,
		43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 3, 1, 0, 0,
		0, 48, 49, 5, 1, 0, 0, 49, 50, 5, 2, 0, 0, 50, 51, 3, 6, 3, 0, 51, 52,
		5, 4, 0, 0, 52, 53, 5, 30, 0, 0, 53, 54, 5, 5, 0, 0, 54, 55, 3, 14, 7,
		0, 55, 56, 5, 6, 0, 0, 56, 5, 1, 0, 0, 0, 57, 58, 5, 3, 0, 0, 58, 7, 1,
		0, 0, 0, 59, 60, 5, 7, 0, 0, 60, 61, 5, 30, 0, 0, 61, 62, 3, 14, 7, 0,
		62, 63, 5, 6, 0, 0, 63, 9, 1, 0, 0, 0, 64, 65, 5, 9, 0, 0, 65, 66, 5, 30,
		0, 0, 66, 67, 5, 6, 0, 0, 67, 11, 1, 0, 0, 0, 68, 69, 5, 8, 0, 0, 69, 70,
		3, 14, 7, 0, 70, 71, 5, 6, 0, 0, 71, 13, 1, 0, 0, 0, 72, 73, 3, 16, 8,
		0, 73, 15, 1, 0, 0, 0, 74, 96, 3, 18, 9, 0, 75, 76, 3, 18, 9, 0, 76, 77,
		5, 11, 0, 0, 77, 78, 3, 16, 8, 0, 78, 96, 1, 0, 0, 0, 79, 80, 3, 18, 9,
		0, 80, 81, 5, 12, 0, 0, 81, 82, 3, 16, 8, 0, 82, 96, 1, 0, 0, 0, 83, 84,
		3, 18, 9, 0, 84, 85, 5, 13, 0, 0, 85, 86, 3, 16, 8, 0, 86, 96, 1, 0, 0,
		0, 87, 88, 3, 18, 9, 0, 88, 89, 5, 14, 0, 0, 89, 90, 3, 16, 8, 0, 90, 96,
		1, 0, 0, 0, 91, 92, 3, 18, 9, 0, 92, 93, 5, 15, 0, 0, 93, 94, 3, 16, 8,
		0, 94, 96, 1, 0, 0, 0, 95, 74, 1, 0, 0, 0, 95, 75, 1, 0, 0, 0, 95, 79,
		1, 0, 0, 0, 95, 83, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0,
		96, 17, 1, 0, 0, 0, 97, 103, 3, 20, 10, 0, 98, 99, 3, 20, 10, 0, 99, 100,
		5, 16, 0, 0, 100, 101, 3, 18, 9, 0, 101, 103, 1, 0, 0, 0, 102, 97, 1, 0,
		0, 0, 102, 98, 1, 0, 0, 0, 103, 19, 1, 0, 0, 0, 104, 118, 3, 22, 11, 0,
		105, 106, 3, 22, 11, 0, 106, 107, 5, 17, 0, 0, 107, 108, 3, 20, 10, 0,
		108, 118, 1, 0, 0, 0, 109, 110, 3, 22, 11, 0, 110, 111, 5, 18, 0, 0, 111,
		112, 3, 20, 10, 0, 112, 118, 1, 0, 0, 0, 113, 114, 3, 22, 11, 0, 114, 115,
		5, 19, 0, 0, 115, 116, 3, 20, 10, 0, 116, 118, 1, 0, 0, 0, 117, 104, 1,
		0, 0, 0, 117, 105, 1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 117, 113, 1, 0, 0,
		0, 118, 21, 1, 0, 0, 0, 119, 129, 3, 24, 12, 0, 120, 121, 3, 24, 12, 0,
		121, 122, 5, 20, 0, 0, 122, 123, 3, 22, 11, 0, 123, 129, 1, 0, 0, 0, 124,
		125, 3, 24, 12, 0, 125, 126, 5, 21, 0, 0, 126, 127, 3, 22, 11, 0, 127,
		129, 1, 0, 0, 0, 128, 119, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 128, 124,
		1, 0, 0, 0, 129, 23, 1, 0, 0, 0, 130, 140, 3, 26, 13, 0, 131, 132, 3, 26,
		13, 0, 132, 133, 5, 22, 0, 0, 133, 134, 3, 24, 12, 0, 134, 140, 1, 0, 0,
		0, 135, 136, 3, 26, 13, 0, 136, 137, 5, 23, 0, 0, 137, 138, 3, 24, 12,
		0, 138, 140, 1, 0, 0, 0, 139, 130, 1, 0, 0, 0, 139, 131, 1, 0, 0, 0, 139,
		135, 1, 0, 0, 0, 140, 25, 1, 0, 0, 0, 141, 147, 3, 28, 14, 0, 142, 143,
		3, 28, 14, 0, 143, 144, 5, 24, 0, 0, 144, 145, 3, 26, 13, 0, 145, 147,
		1, 0, 0, 0, 146, 141, 1, 0, 0, 0, 146, 142, 1, 0, 0, 0, 147, 27, 1, 0,
		0, 0, 148, 154, 3, 30, 15, 0, 149, 150, 5, 26, 0, 0, 150, 151, 3, 14, 7,
		0, 151, 152, 5, 27, 0, 0, 152, 154, 1, 0, 0, 0, 153, 148, 1, 0, 0, 0, 153,
		149, 1, 0, 0, 0, 154, 29, 1, 0, 0, 0, 155, 161, 3, 34, 17, 0, 156, 157,
		5, 21, 0, 0, 157, 161, 3, 14, 7, 0, 158, 159, 5, 25, 0, 0, 159, 161, 3,
		14, 7, 0, 160, 155, 1, 0, 0, 0, 160, 156, 1, 0, 0, 0, 160, 158, 1, 0, 0,
		0, 161, 31, 1, 0, 0, 0, 162, 163, 7, 0, 0, 0, 163, 33, 1, 0, 0, 0, 164,
		167, 3, 32, 16, 0, 165, 167, 5, 30, 0, 0, 166, 164, 1, 0, 0, 0, 166, 165,
		1, 0, 0, 0, 167, 35, 1, 0, 0, 0, 11, 39, 46, 95, 102, 117, 128, 139, 146,
		153, 160, 166,
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
	kochanowskiParserMODULO       = 16
	kochanowskiParserAND          = 17
	kochanowskiParserOR           = 18
	kochanowskiParserXOR          = 19
	kochanowskiParserPLUS         = 20
	kochanowskiParserMINUS        = 21
	kochanowskiParserTIMES        = 22
	kochanowskiParserDIVIDE       = 23
	kochanowskiParserPOWER        = 24
	kochanowskiParserNOT          = 25
	kochanowskiParserFIRST        = 26
	kochanowskiParserCALCULATE    = 27
	kochanowskiParserINTEGER      = 28
	kochanowskiParserDECIMAL      = 29
	kochanowskiParserID           = 30
	kochanowskiParserWS           = 31
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
	kochanowskiParserRULE_num          = 16
	kochanowskiParserRULE_value        = 17
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&898) != 0 {
		{
			p.SetState(36)
			p.Statement()
		}

		p.SetState(41)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserDEFINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Var_create()
		}

	case kochanowskiParserASSIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Var_assign()
		}

	case kochanowskiParserPRINT_WORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Print_()
		}

	case kochanowskiParserREAD_WORD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
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
	WITH_VALUE() antlr.TerminalNode
	Expr() IExprContext
	DOT() antlr.TerminalNode

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

func (s *Var_createContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(kochanowskiParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(kochanowskiParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Type_()
	}
	{
		p.SetState(51)
		p.Match(kochanowskiParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(kochanowskiParserWITH_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Expr()
	}
	{
		p.SetState(55)
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
		p.SetState(57)
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
		p.SetState(59)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Expr()
	}
	{
		p.SetState(62)
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
		p.SetState(64)
		p.Match(kochanowskiParserREAD_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
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
		p.SetState(68)
		p.Match(kochanowskiParserPRINT_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Expr()
	}
	{
		p.SetState(70)
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
		p.SetState(72)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Expr_mod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Expr_mod()
		}
		{
			p.SetState(76)
			p.Match(kochanowskiParserGREATER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Expr_compare()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Expr_mod()
		}
		{
			p.SetState(80)
			p.Match(kochanowskiParserGREATEREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Expr_compare()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Expr_mod()
		}
		{
			p.SetState(84)
			p.Match(kochanowskiParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Expr_compare()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(87)
			p.Expr_mod()
		}
		{
			p.SetState(88)
			p.Match(kochanowskiParserLESSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Expr_compare()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.Expr_mod()
		}
		{
			p.SetState(92)
			p.Match(kochanowskiParserLESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Expr_compare()
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Expr_bit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Expr_bit()
		}
		{
			p.SetState(99)
			p.Match(kochanowskiParserMODULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Expr_add()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Expr_add()
		}
		{
			p.SetState(106)
			p.Match(kochanowskiParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Expr_bit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Expr_add()
		}
		{
			p.SetState(110)
			p.Match(kochanowskiParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Expr_bit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Expr_add()
		}
		{
			p.SetState(114)
			p.Match(kochanowskiParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Expr_mult()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Expr_mult()
		}
		{
			p.SetState(121)
			p.Match(kochanowskiParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Expr_add()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Expr_mult()
		}
		{
			p.SetState(125)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Expr_power()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Expr_power()
		}
		{
			p.SetState(132)
			p.Match(kochanowskiParserTIMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Expr_mult()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.Expr_power()
		}
		{
			p.SetState(136)
			p.Match(kochanowskiParserDIVIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Expr_paren()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Expr_paren()
		}
		{
			p.SetState(143)
			p.Match(kochanowskiParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserMINUS, kochanowskiParserNOT, kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Unary()
		}

	case kochanowskiParserFIRST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(kochanowskiParserFIRST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Expr()
		}
		{
			p.SetState(151)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Value()
		}

	case kochanowskiParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Expr()
		}

	case kochanowskiParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(kochanowskiParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
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

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_num
	return p
}

func InitEmptyNumContext(p *NumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_num
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINTEGER, 0)
}

func (s *NumContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDECIMAL, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitNum(s)
	}
}

func (p *kochanowskiParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, kochanowskiParserRULE_num)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !(_la == kochanowskiParserINTEGER || _la == kochanowskiParserDECIMAL) {
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Num() INumContext
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

func (s *ValueContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
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
	p.EnterRule(localctx, 34, kochanowskiParserRULE_value)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserINTEGER, kochanowskiParserDECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Num()
		}

	case kochanowskiParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(kochanowskiParserID)
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
