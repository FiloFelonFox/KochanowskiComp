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
		"", "'na'", "'macierz liczb ca\\u0142kowitych'", "'macierz liczb zmiennoprzecinkowych'",
		"'macierzy'", "'pod kolumn\\u0105'", "'wierszem'", "'tablicy'", "'pod'",
		"'warto\\u015B\\u0107'", "'kom\\u00F3rk\\u0105'", "'Zdefiniuj'", "'zmienn\\u0105'",
		"'ca\\u0142kowit\\u0105'", "'ca\\u0142kowit\\u0105 olbrzymiej wagi'",
		"'zmiennoprzecinkow\\u0105'", "'zmiennoprzecinkow\\u0105 olbrzymiej precyzji'",
		"'tablic\\u0119 liczb ca\\u0142kowitych'", "'tablic\\u0119 liczb zmiennoprzecinkowych'",
		"'o rozmiarze'", "'o nazwie'", "'o warto\\u015Bci'", "'.'", "'Przypisz'",
		"'Wypisz'", "'Wczytaj'", "'i jednocze\\u015Bnie'", "'lub'", "'z'", "'wi\\u0119ksze ni\\u017C'",
		"'wi\\u0119ksze lub r\\u00F3wne'", "'r\\u00F3wne'", "'mniejsze lub r\\u00F3wne'",
		"'mniejsze ni\\u017C'", "'r\\u00F3\\u017Cne od'", "'modulo'", "'koniunkcja'",
		"'alternatywa'", "'alternatywa wykluczaj\\u0105ca'", "'plus'", "'minus'",
		"'razy'", "'podzieli\\u0107 przez'", "'do pot\\u0119gi'", "'nie'", "'wpierw'",
		"'policz'",
	}
	staticData.SymbolicNames = []string{
		"", "BY", "INT32MATRIX", "F32MATRIX", "MATRIX", "UNDER_COLUMN", "ROW",
		"ARRAY", "UNDER", "VALUE", "CELL", "DEFINE", "VARIABLE", "INT32", "INT64",
		"F32", "F64", "INT32ARRAY", "F32ARRAY", "WITH_SIZE", "NAMED", "WITH_VALUE",
		"DOT", "ASSIGN", "PRINT_WORD", "READ_WORD", "LOGIC_AND", "LOGIC_OR",
		"FROM", "GREATER", "GREATEREQUAL", "EQUAL", "LESSEQUAL", "LESS", "NOTEQUAL",
		"MODULO", "AND", "OR", "XOR", "PLUS", "MINUS", "TIMES", "DIVIDE", "POWER",
		"NOT", "FIRST", "CALCULATE", "INTEGER", "DECIMAL", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"body", "statement", "var_create", "type", "array_create", "array_type",
		"array_assign", "matrix_create", "matrix_type", "matrix_assign", "var_assign",
		"read", "print", "expr", "expr_logic", "logic_operator", "expr_compare",
		"expr_mod", "expr_bit", "expr_add", "expr_mult", "expr_power", "expr_paren",
		"unary", "array_value", "matrix_value", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 268, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 5, 0, 56, 8, 0, 10, 0, 12, 0, 59, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 69, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 87, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 153, 8, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 182, 8, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 189, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 204,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 215, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 3, 20, 226, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 233,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 240, 8, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 249, 8, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 0, 0, 27, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 0, 5, 1, 0, 13, 16, 1, 0, 17, 18, 1, 0, 2, 3, 1, 0, 26, 27, 1,
		0, 47, 49, 270, 0, 57, 1, 0, 0, 0, 2, 68, 1, 0, 0, 0, 4, 86, 1, 0, 0, 0,
		6, 88, 1, 0, 0, 0, 8, 90, 1, 0, 0, 0, 10, 98, 1, 0, 0, 0, 12, 100, 1, 0,
		0, 0, 14, 109, 1, 0, 0, 0, 16, 119, 1, 0, 0, 0, 18, 121, 1, 0, 0, 0, 20,
		132, 1, 0, 0, 0, 22, 137, 1, 0, 0, 0, 24, 141, 1, 0, 0, 0, 26, 145, 1,
		0, 0, 0, 28, 152, 1, 0, 0, 0, 30, 154, 1, 0, 0, 0, 32, 181, 1, 0, 0, 0,
		34, 188, 1, 0, 0, 0, 36, 203, 1, 0, 0, 0, 38, 214, 1, 0, 0, 0, 40, 225,
		1, 0, 0, 0, 42, 232, 1, 0, 0, 0, 44, 239, 1, 0, 0, 0, 46, 248, 1, 0, 0,
		0, 48, 250, 1, 0, 0, 0, 50, 257, 1, 0, 0, 0, 52, 265, 1, 0, 0, 0, 54, 56,
		3, 2, 1, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0,
		57, 58, 1, 0, 0, 0, 58, 1, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 69, 3, 4,
		2, 0, 61, 69, 3, 20, 10, 0, 62, 69, 3, 8, 4, 0, 63, 69, 3, 12, 6, 0, 64,
		69, 3, 14, 7, 0, 65, 69, 3, 18, 9, 0, 66, 69, 3, 24, 12, 0, 67, 69, 3,
		22, 11, 0, 68, 60, 1, 0, 0, 0, 68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0,
		68, 63, 1, 0, 0, 0, 68, 64, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1,
		0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 3, 1, 0, 0, 0, 70, 71, 5, 11, 0, 0, 71,
		72, 5, 12, 0, 0, 72, 73, 3, 6, 3, 0, 73, 74, 5, 20, 0, 0, 74, 75, 5, 49,
		0, 0, 75, 76, 5, 22, 0, 0, 76, 87, 1, 0, 0, 0, 77, 78, 5, 11, 0, 0, 78,
		79, 5, 12, 0, 0, 79, 80, 3, 6, 3, 0, 80, 81, 5, 20, 0, 0, 81, 82, 5, 49,
		0, 0, 82, 83, 5, 21, 0, 0, 83, 84, 3, 26, 13, 0, 84, 85, 5, 22, 0, 0, 85,
		87, 1, 0, 0, 0, 86, 70, 1, 0, 0, 0, 86, 77, 1, 0, 0, 0, 87, 5, 1, 0, 0,
		0, 88, 89, 7, 0, 0, 0, 89, 7, 1, 0, 0, 0, 90, 91, 5, 11, 0, 0, 91, 92,
		3, 10, 5, 0, 92, 93, 5, 20, 0, 0, 93, 94, 5, 49, 0, 0, 94, 95, 5, 19, 0,
		0, 95, 96, 3, 26, 13, 0, 96, 97, 5, 22, 0, 0, 97, 9, 1, 0, 0, 0, 98, 99,
		7, 1, 0, 0, 99, 11, 1, 0, 0, 0, 100, 101, 5, 23, 0, 0, 101, 102, 5, 7,
		0, 0, 102, 103, 5, 49, 0, 0, 103, 104, 5, 8, 0, 0, 104, 105, 3, 26, 13,
		0, 105, 106, 5, 9, 0, 0, 106, 107, 3, 26, 13, 0, 107, 108, 5, 22, 0, 0,
		108, 13, 1, 0, 0, 0, 109, 110, 5, 11, 0, 0, 110, 111, 3, 16, 8, 0, 111,
		112, 5, 20, 0, 0, 112, 113, 5, 49, 0, 0, 113, 114, 5, 19, 0, 0, 114, 115,
		3, 26, 13, 0, 115, 116, 5, 1, 0, 0, 116, 117, 3, 26, 13, 0, 117, 118, 5,
		22, 0, 0, 118, 15, 1, 0, 0, 0, 119, 120, 7, 2, 0, 0, 120, 17, 1, 0, 0,
		0, 121, 122, 5, 23, 0, 0, 122, 123, 5, 4, 0, 0, 123, 124, 5, 49, 0, 0,
		124, 125, 5, 5, 0, 0, 125, 126, 3, 26, 13, 0, 126, 127, 5, 6, 0, 0, 127,
		128, 3, 26, 13, 0, 128, 129, 5, 9, 0, 0, 129, 130, 3, 26, 13, 0, 130, 131,
		5, 22, 0, 0, 131, 19, 1, 0, 0, 0, 132, 133, 5, 23, 0, 0, 133, 134, 5, 49,
		0, 0, 134, 135, 3, 26, 13, 0, 135, 136, 5, 22, 0, 0, 136, 21, 1, 0, 0,
		0, 137, 138, 5, 25, 0, 0, 138, 139, 5, 49, 0, 0, 139, 140, 5, 22, 0, 0,
		140, 23, 1, 0, 0, 0, 141, 142, 5, 24, 0, 0, 142, 143, 3, 26, 13, 0, 143,
		144, 5, 22, 0, 0, 144, 25, 1, 0, 0, 0, 145, 146, 3, 28, 14, 0, 146, 27,
		1, 0, 0, 0, 147, 153, 3, 32, 16, 0, 148, 149, 3, 32, 16, 0, 149, 150, 3,
		30, 15, 0, 150, 151, 3, 28, 14, 0, 151, 153, 1, 0, 0, 0, 152, 147, 1, 0,
		0, 0, 152, 148, 1, 0, 0, 0, 153, 29, 1, 0, 0, 0, 154, 155, 7, 3, 0, 0,
		155, 31, 1, 0, 0, 0, 156, 182, 3, 34, 17, 0, 157, 158, 3, 34, 17, 0, 158,
		159, 5, 29, 0, 0, 159, 160, 3, 32, 16, 0, 160, 182, 1, 0, 0, 0, 161, 162,
		3, 34, 17, 0, 162, 163, 5, 30, 0, 0, 163, 164, 3, 32, 16, 0, 164, 182,
		1, 0, 0, 0, 165, 166, 3, 34, 17, 0, 166, 167, 5, 31, 0, 0, 167, 168, 3,
		32, 16, 0, 168, 182, 1, 0, 0, 0, 169, 170, 3, 34, 17, 0, 170, 171, 5, 32,
		0, 0, 171, 172, 3, 32, 16, 0, 172, 182, 1, 0, 0, 0, 173, 174, 3, 34, 17,
		0, 174, 175, 5, 33, 0, 0, 175, 176, 3, 32, 16, 0, 176, 182, 1, 0, 0, 0,
		177, 178, 3, 34, 17, 0, 178, 179, 5, 34, 0, 0, 179, 180, 3, 32, 16, 0,
		180, 182, 1, 0, 0, 0, 181, 156, 1, 0, 0, 0, 181, 157, 1, 0, 0, 0, 181,
		161, 1, 0, 0, 0, 181, 165, 1, 0, 0, 0, 181, 169, 1, 0, 0, 0, 181, 173,
		1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 182, 33, 1, 0, 0, 0, 183, 189, 3, 36,
		18, 0, 184, 185, 3, 36, 18, 0, 185, 186, 5, 35, 0, 0, 186, 187, 3, 34,
		17, 0, 187, 189, 1, 0, 0, 0, 188, 183, 1, 0, 0, 0, 188, 184, 1, 0, 0, 0,
		189, 35, 1, 0, 0, 0, 190, 204, 3, 38, 19, 0, 191, 192, 3, 38, 19, 0, 192,
		193, 5, 36, 0, 0, 193, 194, 3, 36, 18, 0, 194, 204, 1, 0, 0, 0, 195, 196,
		3, 38, 19, 0, 196, 197, 5, 37, 0, 0, 197, 198, 3, 36, 18, 0, 198, 204,
		1, 0, 0, 0, 199, 200, 3, 38, 19, 0, 200, 201, 5, 38, 0, 0, 201, 202, 3,
		36, 18, 0, 202, 204, 1, 0, 0, 0, 203, 190, 1, 0, 0, 0, 203, 191, 1, 0,
		0, 0, 203, 195, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 204, 37, 1, 0, 0, 0,
		205, 215, 3, 40, 20, 0, 206, 207, 3, 40, 20, 0, 207, 208, 5, 39, 0, 0,
		208, 209, 3, 38, 19, 0, 209, 215, 1, 0, 0, 0, 210, 211, 3, 40, 20, 0, 211,
		212, 5, 40, 0, 0, 212, 213, 3, 38, 19, 0, 213, 215, 1, 0, 0, 0, 214, 205,
		1, 0, 0, 0, 214, 206, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 215, 39, 1, 0,
		0, 0, 216, 226, 3, 42, 21, 0, 217, 218, 3, 42, 21, 0, 218, 219, 5, 41,
		0, 0, 219, 220, 3, 40, 20, 0, 220, 226, 1, 0, 0, 0, 221, 222, 3, 42, 21,
		0, 222, 223, 5, 42, 0, 0, 223, 224, 3, 40, 20, 0, 224, 226, 1, 0, 0, 0,
		225, 216, 1, 0, 0, 0, 225, 217, 1, 0, 0, 0, 225, 221, 1, 0, 0, 0, 226,
		41, 1, 0, 0, 0, 227, 233, 3, 44, 22, 0, 228, 229, 3, 44, 22, 0, 229, 230,
		5, 43, 0, 0, 230, 231, 3, 42, 21, 0, 231, 233, 1, 0, 0, 0, 232, 227, 1,
		0, 0, 0, 232, 228, 1, 0, 0, 0, 233, 43, 1, 0, 0, 0, 234, 240, 3, 46, 23,
		0, 235, 236, 5, 45, 0, 0, 236, 237, 3, 26, 13, 0, 237, 238, 5, 46, 0, 0,
		238, 240, 1, 0, 0, 0, 239, 234, 1, 0, 0, 0, 239, 235, 1, 0, 0, 0, 240,
		45, 1, 0, 0, 0, 241, 249, 3, 52, 26, 0, 242, 249, 3, 50, 25, 0, 243, 249,
		3, 48, 24, 0, 244, 245, 5, 40, 0, 0, 245, 249, 3, 26, 13, 0, 246, 247,
		5, 44, 0, 0, 247, 249, 3, 26, 13, 0, 248, 241, 1, 0, 0, 0, 248, 242, 1,
		0, 0, 0, 248, 243, 1, 0, 0, 0, 248, 244, 1, 0, 0, 0, 248, 246, 1, 0, 0,
		0, 249, 47, 1, 0, 0, 0, 250, 251, 5, 9, 0, 0, 251, 252, 5, 8, 0, 0, 252,
		253, 5, 10, 0, 0, 253, 254, 3, 26, 13, 0, 254, 255, 5, 7, 0, 0, 255, 256,
		5, 49, 0, 0, 256, 49, 1, 0, 0, 0, 257, 258, 5, 9, 0, 0, 258, 259, 5, 5,
		0, 0, 259, 260, 3, 26, 13, 0, 260, 261, 5, 6, 0, 0, 261, 262, 3, 26, 13,
		0, 262, 263, 5, 4, 0, 0, 263, 264, 5, 49, 0, 0, 264, 51, 1, 0, 0, 0, 265,
		266, 7, 4, 0, 0, 266, 53, 1, 0, 0, 0, 12, 57, 68, 86, 152, 181, 188, 203,
		214, 225, 232, 239, 248,
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
	kochanowskiParserBY           = 1
	kochanowskiParserINT32MATRIX  = 2
	kochanowskiParserF32MATRIX    = 3
	kochanowskiParserMATRIX       = 4
	kochanowskiParserUNDER_COLUMN = 5
	kochanowskiParserROW          = 6
	kochanowskiParserARRAY        = 7
	kochanowskiParserUNDER        = 8
	kochanowskiParserVALUE        = 9
	kochanowskiParserCELL         = 10
	kochanowskiParserDEFINE       = 11
	kochanowskiParserVARIABLE     = 12
	kochanowskiParserINT32        = 13
	kochanowskiParserINT64        = 14
	kochanowskiParserF32          = 15
	kochanowskiParserF64          = 16
	kochanowskiParserINT32ARRAY   = 17
	kochanowskiParserF32ARRAY     = 18
	kochanowskiParserWITH_SIZE    = 19
	kochanowskiParserNAMED        = 20
	kochanowskiParserWITH_VALUE   = 21
	kochanowskiParserDOT          = 22
	kochanowskiParserASSIGN       = 23
	kochanowskiParserPRINT_WORD   = 24
	kochanowskiParserREAD_WORD    = 25
	kochanowskiParserLOGIC_AND    = 26
	kochanowskiParserLOGIC_OR     = 27
	kochanowskiParserFROM         = 28
	kochanowskiParserGREATER      = 29
	kochanowskiParserGREATEREQUAL = 30
	kochanowskiParserEQUAL        = 31
	kochanowskiParserLESSEQUAL    = 32
	kochanowskiParserLESS         = 33
	kochanowskiParserNOTEQUAL     = 34
	kochanowskiParserMODULO       = 35
	kochanowskiParserAND          = 36
	kochanowskiParserOR           = 37
	kochanowskiParserXOR          = 38
	kochanowskiParserPLUS         = 39
	kochanowskiParserMINUS        = 40
	kochanowskiParserTIMES        = 41
	kochanowskiParserDIVIDE       = 42
	kochanowskiParserPOWER        = 43
	kochanowskiParserNOT          = 44
	kochanowskiParserFIRST        = 45
	kochanowskiParserCALCULATE    = 46
	kochanowskiParserINTEGER      = 47
	kochanowskiParserDECIMAL      = 48
	kochanowskiParserID           = 49
	kochanowskiParserWS           = 50
)

// kochanowskiParser rules.
const (
	kochanowskiParserRULE_body           = 0
	kochanowskiParserRULE_statement      = 1
	kochanowskiParserRULE_var_create     = 2
	kochanowskiParserRULE_type           = 3
	kochanowskiParserRULE_array_create   = 4
	kochanowskiParserRULE_array_type     = 5
	kochanowskiParserRULE_array_assign   = 6
	kochanowskiParserRULE_matrix_create  = 7
	kochanowskiParserRULE_matrix_type    = 8
	kochanowskiParserRULE_matrix_assign  = 9
	kochanowskiParserRULE_var_assign     = 10
	kochanowskiParserRULE_read           = 11
	kochanowskiParserRULE_print          = 12
	kochanowskiParserRULE_expr           = 13
	kochanowskiParserRULE_expr_logic     = 14
	kochanowskiParserRULE_logic_operator = 15
	kochanowskiParserRULE_expr_compare   = 16
	kochanowskiParserRULE_expr_mod       = 17
	kochanowskiParserRULE_expr_bit       = 18
	kochanowskiParserRULE_expr_add       = 19
	kochanowskiParserRULE_expr_mult      = 20
	kochanowskiParserRULE_expr_power     = 21
	kochanowskiParserRULE_expr_paren     = 22
	kochanowskiParserRULE_unary          = 23
	kochanowskiParserRULE_array_value    = 24
	kochanowskiParserRULE_matrix_value   = 25
	kochanowskiParserRULE_value          = 26
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&58722304) != 0 {
		{
			p.SetState(54)
			p.Statement()
		}

		p.SetState(59)
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
	Array_create() IArray_createContext
	Array_assign() IArray_assignContext
	Matrix_create() IMatrix_createContext
	Matrix_assign() IMatrix_assignContext
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

func (s *StatementContext) Array_create() IArray_createContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_createContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_createContext)
}

func (s *StatementContext) Array_assign() IArray_assignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_assignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_assignContext)
}

func (s *StatementContext) Matrix_create() IMatrix_createContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_createContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_createContext)
}

func (s *StatementContext) Matrix_assign() IMatrix_assignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_assignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_assignContext)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Var_create()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Var_assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Array_create()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Array_assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(64)
			p.Matrix_create()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Matrix_assign()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(66)
			p.Print_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(67)
			p.Read()
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Type_()
		}
		{
			p.SetState(73)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Type_()
		}
		{
			p.SetState(80)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Expr()
		}
		{
			p.SetState(84)
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
	INT64() antlr.TerminalNode
	F32() antlr.TerminalNode
	F64() antlr.TerminalNode

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

func (s *TypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT64, 0)
}

func (s *TypeContext) F32() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF32, 0)
}

func (s *TypeContext) F64() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF64, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122880) != 0) {
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

// IArray_createContext is an interface to support dynamic dispatch.
type IArray_createContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE() antlr.TerminalNode
	Array_type() IArray_typeContext
	NAMED() antlr.TerminalNode
	ID() antlr.TerminalNode
	WITH_SIZE() antlr.TerminalNode
	Expr() IExprContext
	DOT() antlr.TerminalNode

	// IsArray_createContext differentiates from other interfaces.
	IsArray_createContext()
}

type Array_createContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_createContext() *Array_createContext {
	var p = new(Array_createContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_create
	return p
}

func InitEmptyArray_createContext(p *Array_createContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_create
}

func (*Array_createContext) IsArray_createContext() {}

func NewArray_createContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_createContext {
	var p = new(Array_createContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_array_create

	return p
}

func (s *Array_createContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_createContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDEFINE, 0)
}

func (s *Array_createContext) Array_type() IArray_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_typeContext)
}

func (s *Array_createContext) NAMED() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserNAMED, 0)
}

func (s *Array_createContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Array_createContext) WITH_SIZE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserWITH_SIZE, 0)
}

func (s *Array_createContext) Expr() IExprContext {
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

func (s *Array_createContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Array_createContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_createContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_createContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterArray_create(s)
	}
}

func (s *Array_createContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitArray_create(s)
	}
}

func (p *kochanowskiParser) Array_create() (localctx IArray_createContext) {
	localctx = NewArray_createContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kochanowskiParserRULE_array_create)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(kochanowskiParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Array_type()
	}
	{
		p.SetState(92)
		p.Match(kochanowskiParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(kochanowskiParserWITH_SIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Expr()
	}
	{
		p.SetState(96)
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

// IArray_typeContext is an interface to support dynamic dispatch.
type IArray_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32ARRAY() antlr.TerminalNode
	F32ARRAY() antlr.TerminalNode

	// IsArray_typeContext differentiates from other interfaces.
	IsArray_typeContext()
}

type Array_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_typeContext() *Array_typeContext {
	var p = new(Array_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_type
	return p
}

func InitEmptyArray_typeContext(p *Array_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_type
}

func (*Array_typeContext) IsArray_typeContext() {}

func NewArray_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_typeContext {
	var p = new(Array_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_array_type

	return p
}

func (s *Array_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_typeContext) INT32ARRAY() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT32ARRAY, 0)
}

func (s *Array_typeContext) F32ARRAY() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF32ARRAY, 0)
}

func (s *Array_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterArray_type(s)
	}
}

func (s *Array_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitArray_type(s)
	}
}

func (p *kochanowskiParser) Array_type() (localctx IArray_typeContext) {
	localctx = NewArray_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kochanowskiParserRULE_array_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !(_la == kochanowskiParserINT32ARRAY || _la == kochanowskiParserF32ARRAY) {
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

// IArray_assignContext is an interface to support dynamic dispatch.
type IArray_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ARRAY() antlr.TerminalNode
	ID() antlr.TerminalNode
	UNDER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	VALUE() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsArray_assignContext differentiates from other interfaces.
	IsArray_assignContext()
}

type Array_assignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_assignContext() *Array_assignContext {
	var p = new(Array_assignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_assign
	return p
}

func InitEmptyArray_assignContext(p *Array_assignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_assign
}

func (*Array_assignContext) IsArray_assignContext() {}

func NewArray_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_assignContext {
	var p = new(Array_assignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_array_assign

	return p
}

func (s *Array_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserASSIGN, 0)
}

func (s *Array_assignContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserARRAY, 0)
}

func (s *Array_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Array_assignContext) UNDER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserUNDER, 0)
}

func (s *Array_assignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Array_assignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Array_assignContext) VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserVALUE, 0)
}

func (s *Array_assignContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Array_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterArray_assign(s)
	}
}

func (s *Array_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitArray_assign(s)
	}
}

func (p *kochanowskiParser) Array_assign() (localctx IArray_assignContext) {
	localctx = NewArray_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, kochanowskiParserRULE_array_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(kochanowskiParserUNDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Expr()
	}
	{
		p.SetState(105)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Expr()
	}
	{
		p.SetState(107)
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

// IMatrix_createContext is an interface to support dynamic dispatch.
type IMatrix_createContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext
	NAMED() antlr.TerminalNode
	ID() antlr.TerminalNode
	WITH_SIZE() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	BY() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsMatrix_createContext differentiates from other interfaces.
	IsMatrix_createContext()
}

type Matrix_createContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_createContext() *Matrix_createContext {
	var p = new(Matrix_createContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_create
	return p
}

func InitEmptyMatrix_createContext(p *Matrix_createContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_create
}

func (*Matrix_createContext) IsMatrix_createContext() {}

func NewMatrix_createContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_createContext {
	var p = new(Matrix_createContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_matrix_create

	return p
}

func (s *Matrix_createContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_createContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDEFINE, 0)
}

func (s *Matrix_createContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *Matrix_createContext) NAMED() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserNAMED, 0)
}

func (s *Matrix_createContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Matrix_createContext) WITH_SIZE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserWITH_SIZE, 0)
}

func (s *Matrix_createContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Matrix_createContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Matrix_createContext) BY() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserBY, 0)
}

func (s *Matrix_createContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Matrix_createContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_createContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_createContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterMatrix_create(s)
	}
}

func (s *Matrix_createContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitMatrix_create(s)
	}
}

func (p *kochanowskiParser) Matrix_create() (localctx IMatrix_createContext) {
	localctx = NewMatrix_createContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kochanowskiParserRULE_matrix_create)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(kochanowskiParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Matrix_type()
	}
	{
		p.SetState(111)
		p.Match(kochanowskiParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(kochanowskiParserWITH_SIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Expr()
	}
	{
		p.SetState(115)
		p.Match(kochanowskiParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Expr()
	}
	{
		p.SetState(117)
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

// IMatrix_typeContext is an interface to support dynamic dispatch.
type IMatrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32MATRIX() antlr.TerminalNode
	F32MATRIX() antlr.TerminalNode

	// IsMatrix_typeContext differentiates from other interfaces.
	IsMatrix_typeContext()
}

type Matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_typeContext() *Matrix_typeContext {
	var p = new(Matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_type
	return p
}

func InitEmptyMatrix_typeContext(p *Matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_type
}

func (*Matrix_typeContext) IsMatrix_typeContext() {}

func NewMatrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_typeContext {
	var p = new(Matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_matrix_type

	return p
}

func (s *Matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_typeContext) INT32MATRIX() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT32MATRIX, 0)
}

func (s *Matrix_typeContext) F32MATRIX() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF32MATRIX, 0)
}

func (s *Matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterMatrix_type(s)
	}
}

func (s *Matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitMatrix_type(s)
	}
}

func (p *kochanowskiParser) Matrix_type() (localctx IMatrix_typeContext) {
	localctx = NewMatrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kochanowskiParserRULE_matrix_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !(_la == kochanowskiParserINT32MATRIX || _la == kochanowskiParserF32MATRIX) {
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

// IMatrix_assignContext is an interface to support dynamic dispatch.
type IMatrix_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	MATRIX() antlr.TerminalNode
	ID() antlr.TerminalNode
	UNDER_COLUMN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ROW() antlr.TerminalNode
	VALUE() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsMatrix_assignContext differentiates from other interfaces.
	IsMatrix_assignContext()
}

type Matrix_assignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_assignContext() *Matrix_assignContext {
	var p = new(Matrix_assignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_assign
	return p
}

func InitEmptyMatrix_assignContext(p *Matrix_assignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_assign
}

func (*Matrix_assignContext) IsMatrix_assignContext() {}

func NewMatrix_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_assignContext {
	var p = new(Matrix_assignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_matrix_assign

	return p
}

func (s *Matrix_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserASSIGN, 0)
}

func (s *Matrix_assignContext) MATRIX() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserMATRIX, 0)
}

func (s *Matrix_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Matrix_assignContext) UNDER_COLUMN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserUNDER_COLUMN, 0)
}

func (s *Matrix_assignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Matrix_assignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Matrix_assignContext) ROW() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserROW, 0)
}

func (s *Matrix_assignContext) VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserVALUE, 0)
}

func (s *Matrix_assignContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *Matrix_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterMatrix_assign(s)
	}
}

func (s *Matrix_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitMatrix_assign(s)
	}
}

func (p *kochanowskiParser) Matrix_assign() (localctx IMatrix_assignContext) {
	localctx = NewMatrix_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kochanowskiParserRULE_matrix_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(kochanowskiParserUNDER_COLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Expr()
	}
	{
		p.SetState(126)
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Expr()
	}
	{
		p.SetState(128)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Expr()
	}
	{
		p.SetState(130)
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
	p.EnterRule(localctx, 20, kochanowskiParserRULE_var_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Expr()
	}
	{
		p.SetState(135)
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
	p.EnterRule(localctx, 22, kochanowskiParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(kochanowskiParserREAD_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
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
	p.EnterRule(localctx, 24, kochanowskiParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(kochanowskiParserPRINT_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Expr()
	}
	{
		p.SetState(143)
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
	Expr_logic() IExpr_logicContext

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

func (s *ExprContext) Expr_logic() IExpr_logicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_logicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_logicContext)
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
	p.EnterRule(localctx, 26, kochanowskiParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Expr_logic()
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

// IExpr_logicContext is an interface to support dynamic dispatch.
type IExpr_logicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_compare() IExpr_compareContext
	Logic_operator() ILogic_operatorContext
	Expr_logic() IExpr_logicContext

	// IsExpr_logicContext differentiates from other interfaces.
	IsExpr_logicContext()
}

type Expr_logicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_logicContext() *Expr_logicContext {
	var p = new(Expr_logicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_logic
	return p
}

func InitEmptyExpr_logicContext(p *Expr_logicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_expr_logic
}

func (*Expr_logicContext) IsExpr_logicContext() {}

func NewExpr_logicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_logicContext {
	var p = new(Expr_logicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_expr_logic

	return p
}

func (s *Expr_logicContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_logicContext) Expr_compare() IExpr_compareContext {
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

func (s *Expr_logicContext) Logic_operator() ILogic_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogic_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogic_operatorContext)
}

func (s *Expr_logicContext) Expr_logic() IExpr_logicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_logicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_logicContext)
}

func (s *Expr_logicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_logicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_logicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterExpr_logic(s)
	}
}

func (s *Expr_logicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitExpr_logic(s)
	}
}

func (p *kochanowskiParser) Expr_logic() (localctx IExpr_logicContext) {
	localctx = NewExpr_logicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kochanowskiParserRULE_expr_logic)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Expr_compare()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Expr_compare()
		}
		{
			p.SetState(149)
			p.Logic_operator()
		}
		{
			p.SetState(150)
			p.Expr_logic()
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

// ILogic_operatorContext is an interface to support dynamic dispatch.
type ILogic_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOGIC_AND() antlr.TerminalNode
	LOGIC_OR() antlr.TerminalNode

	// IsLogic_operatorContext differentiates from other interfaces.
	IsLogic_operatorContext()
}

type Logic_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogic_operatorContext() *Logic_operatorContext {
	var p = new(Logic_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_logic_operator
	return p
}

func InitEmptyLogic_operatorContext(p *Logic_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_logic_operator
}

func (*Logic_operatorContext) IsLogic_operatorContext() {}

func NewLogic_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logic_operatorContext {
	var p = new(Logic_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_logic_operator

	return p
}

func (s *Logic_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Logic_operatorContext) LOGIC_AND() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserLOGIC_AND, 0)
}

func (s *Logic_operatorContext) LOGIC_OR() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserLOGIC_OR, 0)
}

func (s *Logic_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logic_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logic_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterLogic_operator(s)
	}
}

func (s *Logic_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitLogic_operator(s)
	}
}

func (p *kochanowskiParser) Logic_operator() (localctx ILogic_operatorContext) {
	localctx = NewLogic_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, kochanowskiParserRULE_logic_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		_la = p.GetTokenStream().LA(1)

		if !(_la == kochanowskiParserLOGIC_AND || _la == kochanowskiParserLOGIC_OR) {
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
	p.EnterRule(localctx, 32, kochanowskiParserRULE_expr_compare)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Expr_mod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Expr_mod()
		}
		{
			p.SetState(158)
			p.Match(kochanowskiParserGREATER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Expr_compare()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Expr_mod()
		}
		{
			p.SetState(162)
			p.Match(kochanowskiParserGREATEREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Expr_compare()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(165)
			p.Expr_mod()
		}
		{
			p.SetState(166)
			p.Match(kochanowskiParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Expr_compare()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Expr_mod()
		}
		{
			p.SetState(170)
			p.Match(kochanowskiParserLESSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Expr_compare()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(173)
			p.Expr_mod()
		}
		{
			p.SetState(174)
			p.Match(kochanowskiParserLESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Expr_compare()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(177)
			p.Expr_mod()
		}
		{
			p.SetState(178)
			p.Match(kochanowskiParserNOTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
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
	p.EnterRule(localctx, 34, kochanowskiParserRULE_expr_mod)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Expr_bit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Expr_bit()
		}
		{
			p.SetState(185)
			p.Match(kochanowskiParserMODULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
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
	p.EnterRule(localctx, 36, kochanowskiParserRULE_expr_bit)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Expr_add()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Expr_add()
		}
		{
			p.SetState(192)
			p.Match(kochanowskiParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Expr_bit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Expr_add()
		}
		{
			p.SetState(196)
			p.Match(kochanowskiParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Expr_bit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.Expr_add()
		}
		{
			p.SetState(200)
			p.Match(kochanowskiParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
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
	p.EnterRule(localctx, 38, kochanowskiParserRULE_expr_add)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Expr_mult()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Expr_mult()
		}
		{
			p.SetState(207)
			p.Match(kochanowskiParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Expr_add()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)
			p.Expr_mult()
		}
		{
			p.SetState(211)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
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
	p.EnterRule(localctx, 40, kochanowskiParserRULE_expr_mult)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Expr_power()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Expr_power()
		}
		{
			p.SetState(218)
			p.Match(kochanowskiParserTIMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Expr_mult()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.Expr_power()
		}
		{
			p.SetState(222)
			p.Match(kochanowskiParserDIVIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
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
	p.EnterRule(localctx, 42, kochanowskiParserRULE_expr_power)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Expr_paren()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Expr_paren()
		}
		{
			p.SetState(229)
			p.Match(kochanowskiParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
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
	p.EnterRule(localctx, 44, kochanowskiParserRULE_expr_paren)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserVALUE, kochanowskiParserMINUS, kochanowskiParserNOT, kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Unary()
		}

	case kochanowskiParserFIRST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(kochanowskiParserFIRST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Expr()
		}
		{
			p.SetState(237)
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
	Matrix_value() IMatrix_valueContext
	Array_value() IArray_valueContext
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

func (s *UnaryContext) Matrix_value() IMatrix_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_valueContext)
}

func (s *UnaryContext) Array_value() IArray_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_valueContext)
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
	p.EnterRule(localctx, 46, kochanowskiParserRULE_unary)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Matrix_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.Array_value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Expr()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(246)
			p.Match(kochanowskiParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Expr()
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

// IArray_valueContext is an interface to support dynamic dispatch.
type IArray_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE() antlr.TerminalNode
	UNDER() antlr.TerminalNode
	CELL() antlr.TerminalNode
	Expr() IExprContext
	ARRAY() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsArray_valueContext differentiates from other interfaces.
	IsArray_valueContext()
}

type Array_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_valueContext() *Array_valueContext {
	var p = new(Array_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_value
	return p
}

func InitEmptyArray_valueContext(p *Array_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_array_value
}

func (*Array_valueContext) IsArray_valueContext() {}

func NewArray_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_valueContext {
	var p = new(Array_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_array_value

	return p
}

func (s *Array_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_valueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserVALUE, 0)
}

func (s *Array_valueContext) UNDER() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserUNDER, 0)
}

func (s *Array_valueContext) CELL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserCELL, 0)
}

func (s *Array_valueContext) Expr() IExprContext {
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

func (s *Array_valueContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserARRAY, 0)
}

func (s *Array_valueContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Array_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterArray_value(s)
	}
}

func (s *Array_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitArray_value(s)
	}
}

func (p *kochanowskiParser) Array_value() (localctx IArray_valueContext) {
	localctx = NewArray_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, kochanowskiParserRULE_array_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Match(kochanowskiParserUNDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(kochanowskiParserCELL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Expr()
	}
	{
		p.SetState(254)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(kochanowskiParserID)
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

// IMatrix_valueContext is an interface to support dynamic dispatch.
type IMatrix_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE() antlr.TerminalNode
	UNDER_COLUMN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ROW() antlr.TerminalNode
	MATRIX() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsMatrix_valueContext differentiates from other interfaces.
	IsMatrix_valueContext()
}

type Matrix_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_valueContext() *Matrix_valueContext {
	var p = new(Matrix_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_value
	return p
}

func InitEmptyMatrix_valueContext(p *Matrix_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_matrix_value
}

func (*Matrix_valueContext) IsMatrix_valueContext() {}

func NewMatrix_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_valueContext {
	var p = new(Matrix_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_matrix_value

	return p
}

func (s *Matrix_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_valueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserVALUE, 0)
}

func (s *Matrix_valueContext) UNDER_COLUMN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserUNDER_COLUMN, 0)
}

func (s *Matrix_valueContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Matrix_valueContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Matrix_valueContext) ROW() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserROW, 0)
}

func (s *Matrix_valueContext) MATRIX() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserMATRIX, 0)
}

func (s *Matrix_valueContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Matrix_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterMatrix_value(s)
	}
}

func (s *Matrix_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitMatrix_value(s)
	}
}

func (p *kochanowskiParser) Matrix_value() (localctx IMatrix_valueContext) {
	localctx = NewMatrix_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, kochanowskiParserRULE_matrix_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(kochanowskiParserUNDER_COLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Expr()
	}
	{
		p.SetState(260)
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Expr()
	}
	{
		p.SetState(262)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(kochanowskiParserID)
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
	p.EnterRule(localctx, 52, kochanowskiParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0) {
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
