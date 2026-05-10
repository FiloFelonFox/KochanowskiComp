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
		"", "'{'", "'}'", "'napis'", "'na'", "'macierz liczb ca\\u0142kowitych'",
		"'macierz liczb zmiennoprzecinkowych'", "'macierzy'", "'pod kolumn\\u0105'",
		"'wierszem'", "'tablicy'", "'pod'", "'warto\\u015B\\u0107'", "'kom\\u00F3rk\\u0105'",
		"'Zdefiniuj'", "'zmienn\\u0105'", "'ca\\u0142kowit\\u0105'", "'ca\\u0142kowit\\u0105 olbrzymiej wagi'",
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
		"", "START_BLOCK", "END_BLOCK", "STRING", "BY", "INT32MATRIX", "F32MATRIX",
		"MATRIX", "UNDER_COLUMN", "ROW", "ARRAY", "UNDER", "VALUE", "CELL",
		"DEFINE", "VARIABLE", "INT32", "INT64", "F32", "F64", "INT32ARRAY",
		"F32ARRAY", "WITH_SIZE", "NAMED", "WITH_VALUE", "DOT", "ASSIGN", "PRINT_WORD",
		"READ_WORD", "LOGIC_AND", "LOGIC_OR", "FROM", "GREATER", "GREATEREQUAL",
		"EQUAL", "LESSEQUAL", "LESS", "NOTEQUAL", "MODULO", "AND", "OR", "XOR",
		"PLUS", "MINUS", "TIMES", "DIVIDE", "POWER", "NOT", "FIRST", "CALCULATE",
		"STRING_LITERAL", "INTEGER", "DECIMAL", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "body", "statement", "block", "var_create", "type", "array_create",
		"array_type", "array_assign", "matrix_create", "matrix_type", "matrix_assign",
		"var_assign", "read", "print", "expr", "expr_logic", "logic_operator",
		"expr_compare", "expr_mod", "expr_bit", "expr_add", "expr_mult", "expr_power",
		"expr_paren", "unary", "string_value", "array_value", "matrix_value",
		"value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 302, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1, 1, 5, 1,
		64, 8, 1, 10, 1, 12, 1, 67, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 100, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 122, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 168, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		185, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 214, 8,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 221, 8, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 236, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 247, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 258, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 3, 23, 265, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 272,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 281, 8,
		25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 0,
		0, 30, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 0, 5, 1, 0, 16, 19, 2,
		0, 3, 3, 20, 21, 1, 0, 5, 6, 1, 0, 29, 30, 1, 0, 51, 53, 304, 0, 60, 1,
		0, 0, 0, 2, 65, 1, 0, 0, 0, 4, 77, 1, 0, 0, 0, 6, 79, 1, 0, 0, 0, 8, 99,
		1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 121, 1, 0, 0, 0, 14, 123, 1, 0, 0,
		0, 16, 125, 1, 0, 0, 0, 18, 134, 1, 0, 0, 0, 20, 144, 1, 0, 0, 0, 22, 146,
		1, 0, 0, 0, 24, 167, 1, 0, 0, 0, 26, 169, 1, 0, 0, 0, 28, 173, 1, 0, 0,
		0, 30, 177, 1, 0, 0, 0, 32, 184, 1, 0, 0, 0, 34, 186, 1, 0, 0, 0, 36, 213,
		1, 0, 0, 0, 38, 220, 1, 0, 0, 0, 40, 235, 1, 0, 0, 0, 42, 246, 1, 0, 0,
		0, 44, 257, 1, 0, 0, 0, 46, 264, 1, 0, 0, 0, 48, 271, 1, 0, 0, 0, 50, 280,
		1, 0, 0, 0, 52, 282, 1, 0, 0, 0, 54, 284, 1, 0, 0, 0, 56, 291, 1, 0, 0,
		0, 58, 299, 1, 0, 0, 0, 60, 61, 3, 2, 1, 0, 61, 1, 1, 0, 0, 0, 62, 64,
		3, 4, 2, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 3, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 78, 3, 8,
		4, 0, 69, 78, 3, 24, 12, 0, 70, 78, 3, 12, 6, 0, 71, 78, 3, 16, 8, 0, 72,
		78, 3, 18, 9, 0, 73, 78, 3, 22, 11, 0, 74, 78, 3, 28, 14, 0, 75, 78, 3,
		26, 13, 0, 76, 78, 3, 6, 3, 0, 77, 68, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0,
		77, 70, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73, 1,
		0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78,
		5, 1, 0, 0, 0, 79, 80, 5, 1, 0, 0, 80, 81, 3, 2, 1, 0, 81, 82, 5, 2, 0,
		0, 82, 7, 1, 0, 0, 0, 83, 84, 5, 14, 0, 0, 84, 85, 5, 15, 0, 0, 85, 86,
		3, 10, 5, 0, 86, 87, 5, 23, 0, 0, 87, 88, 5, 53, 0, 0, 88, 89, 5, 25, 0,
		0, 89, 100, 1, 0, 0, 0, 90, 91, 5, 14, 0, 0, 91, 92, 5, 15, 0, 0, 92, 93,
		3, 10, 5, 0, 93, 94, 5, 23, 0, 0, 94, 95, 5, 53, 0, 0, 95, 96, 5, 24, 0,
		0, 96, 97, 3, 30, 15, 0, 97, 98, 5, 25, 0, 0, 98, 100, 1, 0, 0, 0, 99,
		83, 1, 0, 0, 0, 99, 90, 1, 0, 0, 0, 100, 9, 1, 0, 0, 0, 101, 102, 7, 0,
		0, 0, 102, 11, 1, 0, 0, 0, 103, 104, 5, 14, 0, 0, 104, 105, 3, 14, 7, 0,
		105, 106, 5, 23, 0, 0, 106, 107, 5, 53, 0, 0, 107, 108, 5, 22, 0, 0, 108,
		109, 3, 30, 15, 0, 109, 110, 5, 25, 0, 0, 110, 122, 1, 0, 0, 0, 111, 112,
		5, 14, 0, 0, 112, 113, 3, 14, 7, 0, 113, 114, 5, 23, 0, 0, 114, 115, 5,
		53, 0, 0, 115, 116, 5, 22, 0, 0, 116, 117, 3, 30, 15, 0, 117, 118, 5, 24,
		0, 0, 118, 119, 3, 52, 26, 0, 119, 120, 5, 25, 0, 0, 120, 122, 1, 0, 0,
		0, 121, 103, 1, 0, 0, 0, 121, 111, 1, 0, 0, 0, 122, 13, 1, 0, 0, 0, 123,
		124, 7, 1, 0, 0, 124, 15, 1, 0, 0, 0, 125, 126, 5, 26, 0, 0, 126, 127,
		5, 10, 0, 0, 127, 128, 5, 53, 0, 0, 128, 129, 5, 11, 0, 0, 129, 130, 3,
		30, 15, 0, 130, 131, 5, 12, 0, 0, 131, 132, 3, 30, 15, 0, 132, 133, 5,
		25, 0, 0, 133, 17, 1, 0, 0, 0, 134, 135, 5, 14, 0, 0, 135, 136, 3, 20,
		10, 0, 136, 137, 5, 23, 0, 0, 137, 138, 5, 53, 0, 0, 138, 139, 5, 22, 0,
		0, 139, 140, 3, 30, 15, 0, 140, 141, 5, 4, 0, 0, 141, 142, 3, 30, 15, 0,
		142, 143, 5, 25, 0, 0, 143, 19, 1, 0, 0, 0, 144, 145, 7, 2, 0, 0, 145,
		21, 1, 0, 0, 0, 146, 147, 5, 26, 0, 0, 147, 148, 5, 7, 0, 0, 148, 149,
		5, 53, 0, 0, 149, 150, 5, 8, 0, 0, 150, 151, 3, 30, 15, 0, 151, 152, 5,
		9, 0, 0, 152, 153, 3, 30, 15, 0, 153, 154, 5, 12, 0, 0, 154, 155, 3, 30,
		15, 0, 155, 156, 5, 25, 0, 0, 156, 23, 1, 0, 0, 0, 157, 158, 5, 26, 0,
		0, 158, 159, 5, 53, 0, 0, 159, 160, 3, 30, 15, 0, 160, 161, 5, 25, 0, 0,
		161, 168, 1, 0, 0, 0, 162, 163, 5, 26, 0, 0, 163, 164, 5, 53, 0, 0, 164,
		165, 3, 52, 26, 0, 165, 166, 5, 25, 0, 0, 166, 168, 1, 0, 0, 0, 167, 157,
		1, 0, 0, 0, 167, 162, 1, 0, 0, 0, 168, 25, 1, 0, 0, 0, 169, 170, 5, 28,
		0, 0, 170, 171, 5, 53, 0, 0, 171, 172, 5, 25, 0, 0, 172, 27, 1, 0, 0, 0,
		173, 174, 5, 27, 0, 0, 174, 175, 3, 30, 15, 0, 175, 176, 5, 25, 0, 0, 176,
		29, 1, 0, 0, 0, 177, 178, 3, 32, 16, 0, 178, 31, 1, 0, 0, 0, 179, 185,
		3, 36, 18, 0, 180, 181, 3, 36, 18, 0, 181, 182, 3, 34, 17, 0, 182, 183,
		3, 32, 16, 0, 183, 185, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 184, 180, 1,
		0, 0, 0, 185, 33, 1, 0, 0, 0, 186, 187, 7, 3, 0, 0, 187, 35, 1, 0, 0, 0,
		188, 214, 3, 38, 19, 0, 189, 190, 3, 38, 19, 0, 190, 191, 5, 32, 0, 0,
		191, 192, 3, 36, 18, 0, 192, 214, 1, 0, 0, 0, 193, 194, 3, 38, 19, 0, 194,
		195, 5, 33, 0, 0, 195, 196, 3, 36, 18, 0, 196, 214, 1, 0, 0, 0, 197, 198,
		3, 38, 19, 0, 198, 199, 5, 34, 0, 0, 199, 200, 3, 36, 18, 0, 200, 214,
		1, 0, 0, 0, 201, 202, 3, 38, 19, 0, 202, 203, 5, 35, 0, 0, 203, 204, 3,
		36, 18, 0, 204, 214, 1, 0, 0, 0, 205, 206, 3, 38, 19, 0, 206, 207, 5, 36,
		0, 0, 207, 208, 3, 36, 18, 0, 208, 214, 1, 0, 0, 0, 209, 210, 3, 38, 19,
		0, 210, 211, 5, 37, 0, 0, 211, 212, 3, 36, 18, 0, 212, 214, 1, 0, 0, 0,
		213, 188, 1, 0, 0, 0, 213, 189, 1, 0, 0, 0, 213, 193, 1, 0, 0, 0, 213,
		197, 1, 0, 0, 0, 213, 201, 1, 0, 0, 0, 213, 205, 1, 0, 0, 0, 213, 209,
		1, 0, 0, 0, 214, 37, 1, 0, 0, 0, 215, 221, 3, 40, 20, 0, 216, 217, 3, 40,
		20, 0, 217, 218, 5, 38, 0, 0, 218, 219, 3, 38, 19, 0, 219, 221, 1, 0, 0,
		0, 220, 215, 1, 0, 0, 0, 220, 216, 1, 0, 0, 0, 221, 39, 1, 0, 0, 0, 222,
		236, 3, 42, 21, 0, 223, 224, 3, 42, 21, 0, 224, 225, 5, 39, 0, 0, 225,
		226, 3, 40, 20, 0, 226, 236, 1, 0, 0, 0, 227, 228, 3, 42, 21, 0, 228, 229,
		5, 40, 0, 0, 229, 230, 3, 40, 20, 0, 230, 236, 1, 0, 0, 0, 231, 232, 3,
		42, 21, 0, 232, 233, 5, 41, 0, 0, 233, 234, 3, 40, 20, 0, 234, 236, 1,
		0, 0, 0, 235, 222, 1, 0, 0, 0, 235, 223, 1, 0, 0, 0, 235, 227, 1, 0, 0,
		0, 235, 231, 1, 0, 0, 0, 236, 41, 1, 0, 0, 0, 237, 247, 3, 44, 22, 0, 238,
		239, 3, 44, 22, 0, 239, 240, 5, 42, 0, 0, 240, 241, 3, 42, 21, 0, 241,
		247, 1, 0, 0, 0, 242, 243, 3, 44, 22, 0, 243, 244, 5, 43, 0, 0, 244, 245,
		3, 42, 21, 0, 245, 247, 1, 0, 0, 0, 246, 237, 1, 0, 0, 0, 246, 238, 1,
		0, 0, 0, 246, 242, 1, 0, 0, 0, 247, 43, 1, 0, 0, 0, 248, 258, 3, 46, 23,
		0, 249, 250, 3, 46, 23, 0, 250, 251, 5, 44, 0, 0, 251, 252, 3, 44, 22,
		0, 252, 258, 1, 0, 0, 0, 253, 254, 3, 46, 23, 0, 254, 255, 5, 45, 0, 0,
		255, 256, 3, 44, 22, 0, 256, 258, 1, 0, 0, 0, 257, 248, 1, 0, 0, 0, 257,
		249, 1, 0, 0, 0, 257, 253, 1, 0, 0, 0, 258, 45, 1, 0, 0, 0, 259, 265, 3,
		48, 24, 0, 260, 261, 3, 48, 24, 0, 261, 262, 5, 46, 0, 0, 262, 263, 3,
		46, 23, 0, 263, 265, 1, 0, 0, 0, 264, 259, 1, 0, 0, 0, 264, 260, 1, 0,
		0, 0, 265, 47, 1, 0, 0, 0, 266, 272, 3, 50, 25, 0, 267, 268, 5, 48, 0,
		0, 268, 269, 3, 30, 15, 0, 269, 270, 5, 49, 0, 0, 270, 272, 1, 0, 0, 0,
		271, 266, 1, 0, 0, 0, 271, 267, 1, 0, 0, 0, 272, 49, 1, 0, 0, 0, 273, 281,
		3, 58, 29, 0, 274, 281, 3, 56, 28, 0, 275, 281, 3, 54, 27, 0, 276, 277,
		5, 43, 0, 0, 277, 281, 3, 30, 15, 0, 278, 279, 5, 47, 0, 0, 279, 281, 3,
		30, 15, 0, 280, 273, 1, 0, 0, 0, 280, 274, 1, 0, 0, 0, 280, 275, 1, 0,
		0, 0, 280, 276, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 51, 1, 0, 0, 0,
		282, 283, 5, 50, 0, 0, 283, 53, 1, 0, 0, 0, 284, 285, 5, 12, 0, 0, 285,
		286, 5, 11, 0, 0, 286, 287, 5, 13, 0, 0, 287, 288, 3, 30, 15, 0, 288, 289,
		5, 10, 0, 0, 289, 290, 5, 53, 0, 0, 290, 55, 1, 0, 0, 0, 291, 292, 5, 12,
		0, 0, 292, 293, 5, 8, 0, 0, 293, 294, 3, 30, 15, 0, 294, 295, 5, 9, 0,
		0, 295, 296, 3, 30, 15, 0, 296, 297, 5, 7, 0, 0, 297, 298, 5, 53, 0, 0,
		298, 57, 1, 0, 0, 0, 299, 300, 7, 4, 0, 0, 300, 59, 1, 0, 0, 0, 14, 65,
		77, 99, 121, 167, 184, 213, 220, 235, 246, 257, 264, 271, 280,
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
	kochanowskiParserEOF            = antlr.TokenEOF
	kochanowskiParserSTART_BLOCK    = 1
	kochanowskiParserEND_BLOCK      = 2
	kochanowskiParserSTRING         = 3
	kochanowskiParserBY             = 4
	kochanowskiParserINT32MATRIX    = 5
	kochanowskiParserF32MATRIX      = 6
	kochanowskiParserMATRIX         = 7
	kochanowskiParserUNDER_COLUMN   = 8
	kochanowskiParserROW            = 9
	kochanowskiParserARRAY          = 10
	kochanowskiParserUNDER          = 11
	kochanowskiParserVALUE          = 12
	kochanowskiParserCELL           = 13
	kochanowskiParserDEFINE         = 14
	kochanowskiParserVARIABLE       = 15
	kochanowskiParserINT32          = 16
	kochanowskiParserINT64          = 17
	kochanowskiParserF32            = 18
	kochanowskiParserF64            = 19
	kochanowskiParserINT32ARRAY     = 20
	kochanowskiParserF32ARRAY       = 21
	kochanowskiParserWITH_SIZE      = 22
	kochanowskiParserNAMED          = 23
	kochanowskiParserWITH_VALUE     = 24
	kochanowskiParserDOT            = 25
	kochanowskiParserASSIGN         = 26
	kochanowskiParserPRINT_WORD     = 27
	kochanowskiParserREAD_WORD      = 28
	kochanowskiParserLOGIC_AND      = 29
	kochanowskiParserLOGIC_OR       = 30
	kochanowskiParserFROM           = 31
	kochanowskiParserGREATER        = 32
	kochanowskiParserGREATEREQUAL   = 33
	kochanowskiParserEQUAL          = 34
	kochanowskiParserLESSEQUAL      = 35
	kochanowskiParserLESS           = 36
	kochanowskiParserNOTEQUAL       = 37
	kochanowskiParserMODULO         = 38
	kochanowskiParserAND            = 39
	kochanowskiParserOR             = 40
	kochanowskiParserXOR            = 41
	kochanowskiParserPLUS           = 42
	kochanowskiParserMINUS          = 43
	kochanowskiParserTIMES          = 44
	kochanowskiParserDIVIDE         = 45
	kochanowskiParserPOWER          = 46
	kochanowskiParserNOT            = 47
	kochanowskiParserFIRST          = 48
	kochanowskiParserCALCULATE      = 49
	kochanowskiParserSTRING_LITERAL = 50
	kochanowskiParserINTEGER        = 51
	kochanowskiParserDECIMAL        = 52
	kochanowskiParserID             = 53
	kochanowskiParserWS             = 54
)

// kochanowskiParser rules.
const (
	kochanowskiParserRULE_prog           = 0
	kochanowskiParserRULE_body           = 1
	kochanowskiParserRULE_statement      = 2
	kochanowskiParserRULE_block          = 3
	kochanowskiParserRULE_var_create     = 4
	kochanowskiParserRULE_type           = 5
	kochanowskiParserRULE_array_create   = 6
	kochanowskiParserRULE_array_type     = 7
	kochanowskiParserRULE_array_assign   = 8
	kochanowskiParserRULE_matrix_create  = 9
	kochanowskiParserRULE_matrix_type    = 10
	kochanowskiParserRULE_matrix_assign  = 11
	kochanowskiParserRULE_var_assign     = 12
	kochanowskiParserRULE_read           = 13
	kochanowskiParserRULE_print          = 14
	kochanowskiParserRULE_expr           = 15
	kochanowskiParserRULE_expr_logic     = 16
	kochanowskiParserRULE_logic_operator = 17
	kochanowskiParserRULE_expr_compare   = 18
	kochanowskiParserRULE_expr_mod       = 19
	kochanowskiParserRULE_expr_bit       = 20
	kochanowskiParserRULE_expr_add       = 21
	kochanowskiParserRULE_expr_mult      = 22
	kochanowskiParserRULE_expr_power     = 23
	kochanowskiParserRULE_expr_paren     = 24
	kochanowskiParserRULE_unary          = 25
	kochanowskiParserRULE_string_value   = 26
	kochanowskiParserRULE_array_value    = 27
	kochanowskiParserRULE_matrix_value   = 28
	kochanowskiParserRULE_value          = 29
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Body() IBodyContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *kochanowskiParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, kochanowskiParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Body()
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
	p.EnterRule(localctx, 2, kochanowskiParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469778434) != 0 {
		{
			p.SetState(62)
			p.Statement()
		}

		p.SetState(67)
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
	Block() IBlockContext

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

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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
	p.EnterRule(localctx, 4, kochanowskiParserRULE_statement)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Var_create()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Var_assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Array_create()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Array_assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.Matrix_create()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.Matrix_assign()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(74)
			p.Print_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.Read()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.Block()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	START_BLOCK() antlr.TerminalNode
	Body() IBodyContext
	END_BLOCK() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) START_BLOCK() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserSTART_BLOCK, 0)
}

func (s *BlockContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *BlockContext) END_BLOCK() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserEND_BLOCK, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *kochanowskiParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kochanowskiParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(kochanowskiParserSTART_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Body()
	}
	{
		p.SetState(81)
		p.Match(kochanowskiParserEND_BLOCK)
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
	p.EnterRule(localctx, 8, kochanowskiParserRULE_var_create)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Type_()
		}
		{
			p.SetState(86)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
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
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Type_()
		}
		{
			p.SetState(93)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Expr()
		}
		{
			p.SetState(97)
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
	p.EnterRule(localctx, 10, kochanowskiParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0) {
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
	WITH_VALUE() antlr.TerminalNode
	String_value() IString_valueContext

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

func (s *Array_createContext) WITH_VALUE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserWITH_VALUE, 0)
}

func (s *Array_createContext) String_value() IString_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_valueContext)
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
	p.EnterRule(localctx, 12, kochanowskiParserRULE_array_create)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Array_type()
		}
		{
			p.SetState(105)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(kochanowskiParserWITH_SIZE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Expr()
		}
		{
			p.SetState(109)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Array_type()
		}
		{
			p.SetState(113)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(kochanowskiParserWITH_SIZE)
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
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.String_value()
		}
		{
			p.SetState(119)
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

// IArray_typeContext is an interface to support dynamic dispatch.
type IArray_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32ARRAY() antlr.TerminalNode
	F32ARRAY() antlr.TerminalNode
	STRING() antlr.TerminalNode

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

func (s *Array_typeContext) STRING() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserSTRING, 0)
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
	p.EnterRule(localctx, 14, kochanowskiParserRULE_array_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3145736) != 0) {
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
	p.EnterRule(localctx, 16, kochanowskiParserRULE_array_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(kochanowskiParserUNDER)
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
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Expr()
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 18, kochanowskiParserRULE_matrix_create)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(kochanowskiParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Matrix_type()
	}
	{
		p.SetState(136)
		p.Match(kochanowskiParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(kochanowskiParserWITH_SIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Expr()
	}
	{
		p.SetState(140)
		p.Match(kochanowskiParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Expr()
	}
	{
		p.SetState(142)
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
	p.EnterRule(localctx, 20, kochanowskiParserRULE_matrix_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
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
	p.EnterRule(localctx, 22, kochanowskiParserRULE_matrix_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(kochanowskiParserUNDER_COLUMN)
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
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Expr()
	}
	{
		p.SetState(153)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Expr()
	}
	{
		p.SetState(155)
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
	String_value() IString_valueContext

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

func (s *Var_assignContext) String_value() IString_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_valueContext)
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
	p.EnterRule(localctx, 24, kochanowskiParserRULE_var_assign)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(kochanowskiParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Expr()
		}
		{
			p.SetState(160)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(kochanowskiParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.String_value()
		}
		{
			p.SetState(165)
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
	p.EnterRule(localctx, 26, kochanowskiParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(kochanowskiParserREAD_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
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
	p.EnterRule(localctx, 28, kochanowskiParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(kochanowskiParserPRINT_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Expr()
	}
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 30, kochanowskiParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
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
	p.EnterRule(localctx, 32, kochanowskiParserRULE_expr_logic)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Expr_compare()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Expr_compare()
		}
		{
			p.SetState(181)
			p.Logic_operator()
		}
		{
			p.SetState(182)
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
	p.EnterRule(localctx, 34, kochanowskiParserRULE_logic_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
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
	p.EnterRule(localctx, 36, kochanowskiParserRULE_expr_compare)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Expr_mod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Expr_mod()
		}
		{
			p.SetState(190)
			p.Match(kochanowskiParserGREATER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Expr_compare()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Expr_mod()
		}
		{
			p.SetState(194)
			p.Match(kochanowskiParserGREATEREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Expr_compare()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)
			p.Expr_mod()
		}
		{
			p.SetState(198)
			p.Match(kochanowskiParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Expr_compare()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)
			p.Expr_mod()
		}
		{
			p.SetState(202)
			p.Match(kochanowskiParserLESSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Expr_compare()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(205)
			p.Expr_mod()
		}
		{
			p.SetState(206)
			p.Match(kochanowskiParserLESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Expr_compare()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(209)
			p.Expr_mod()
		}
		{
			p.SetState(210)
			p.Match(kochanowskiParserNOTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
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
	p.EnterRule(localctx, 38, kochanowskiParserRULE_expr_mod)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Expr_bit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Expr_bit()
		}
		{
			p.SetState(217)
			p.Match(kochanowskiParserMODULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 40, kochanowskiParserRULE_expr_bit)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Expr_add()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Expr_add()
		}
		{
			p.SetState(224)
			p.Match(kochanowskiParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Expr_bit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.Expr_add()
		}
		{
			p.SetState(228)
			p.Match(kochanowskiParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Expr_bit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.Expr_add()
		}
		{
			p.SetState(232)
			p.Match(kochanowskiParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
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
	p.EnterRule(localctx, 42, kochanowskiParserRULE_expr_add)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Expr_mult()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Expr_mult()
		}
		{
			p.SetState(239)
			p.Match(kochanowskiParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Expr_add()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Expr_mult()
		}
		{
			p.SetState(243)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
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
	p.EnterRule(localctx, 44, kochanowskiParserRULE_expr_mult)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Expr_power()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.Expr_power()
		}
		{
			p.SetState(250)
			p.Match(kochanowskiParserTIMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.Expr_mult()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)
			p.Expr_power()
		}
		{
			p.SetState(254)
			p.Match(kochanowskiParserDIVIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
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
	p.EnterRule(localctx, 46, kochanowskiParserRULE_expr_power)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Expr_paren()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Expr_paren()
		}
		{
			p.SetState(261)
			p.Match(kochanowskiParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
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
	p.EnterRule(localctx, 48, kochanowskiParserRULE_expr_paren)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserVALUE, kochanowskiParserMINUS, kochanowskiParserNOT, kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Unary()
		}

	case kochanowskiParserFIRST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(kochanowskiParserFIRST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Expr()
		}
		{
			p.SetState(269)
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
	p.EnterRule(localctx, 50, kochanowskiParserRULE_unary)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Matrix_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(275)
			p.Array_value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(276)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Expr()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(278)
			p.Match(kochanowskiParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
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

// IString_valueContext is an interface to support dynamic dispatch.
type IString_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsString_valueContext differentiates from other interfaces.
	IsString_valueContext()
}

type String_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_valueContext() *String_valueContext {
	var p = new(String_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_string_value
	return p
}

func InitEmptyString_valueContext(p *String_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_string_value
}

func (*String_valueContext) IsString_valueContext() {}

func NewString_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_valueContext {
	var p = new(String_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_string_value

	return p
}

func (s *String_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *String_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserSTRING_LITERAL, 0)
}

func (s *String_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterString_value(s)
	}
}

func (s *String_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitString_value(s)
	}
}

func (p *kochanowskiParser) String_value() (localctx IString_valueContext) {
	localctx = NewString_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, kochanowskiParserRULE_string_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(kochanowskiParserSTRING_LITERAL)
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
	p.EnterRule(localctx, 54, kochanowskiParserRULE_array_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(kochanowskiParserUNDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(kochanowskiParserCELL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Expr()
	}
	{
		p.SetState(288)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
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
	p.EnterRule(localctx, 56, kochanowskiParserRULE_matrix_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(kochanowskiParserUNDER_COLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Expr()
	}
	{
		p.SetState(294)
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Expr()
	}
	{
		p.SetState(296)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
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
	p.EnterRule(localctx, 58, kochanowskiParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15762598695796736) != 0) {
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
