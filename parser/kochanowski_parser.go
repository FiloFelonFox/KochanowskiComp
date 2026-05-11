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
		"", "'('", "')'", "','", "'return'", "'funkcja'", "'Je\\u015Bli'", "'to'",
		"'w przeciwnym wypadku'", "'Powtarzaj dop\\u00F3ki'", "'{'", "'}'",
		"'napis'", "'na'", "'macierz liczb ca\\u0142kowitych'", "'macierz liczb zmiennoprzecinkowych'",
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
		"", "", "", "", "", "FUNCTION", "IF", "THEN", "ELSE", "WHILE", "START_BLOCK",
		"END_BLOCK", "STRING", "BY", "INT32MATRIX", "F32MATRIX", "MATRIX", "UNDER_COLUMN",
		"ROW", "ARRAY", "UNDER", "VALUE", "CELL", "DEFINE", "VARIABLE", "INT32",
		"INT64", "F32", "F64", "INT32ARRAY", "F32ARRAY", "WITH_SIZE", "NAMED",
		"WITH_VALUE", "DOT", "ASSIGN", "PRINT_WORD", "READ_WORD", "LOGIC_AND",
		"LOGIC_OR", "FROM", "GREATER", "GREATEREQUAL", "EQUAL", "LESSEQUAL",
		"LESS", "NOTEQUAL", "MODULO", "AND", "OR", "XOR", "PLUS", "MINUS", "TIMES",
		"DIVIDE", "POWER", "NOT", "FIRST", "CALCULATE", "STRING_LITERAL", "INTEGER",
		"DECIMAL", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "body", "function_decl", "param_list", "param", "func_type",
		"statement", "if", "conditional_body", "if_expr", "if_body", "else_body",
		"while", "while_body", "return", "block", "var_create", "type", "array_create",
		"array_type", "array_assign", "matrix_create", "matrix_type", "matrix_assign",
		"var_assign", "read", "print", "expr", "expr_logic", "logic_operator",
		"expr_compare", "expr_mod", "expr_bit", "expr_add", "expr_mult", "expr_power",
		"expr_paren", "unary", "function_call", "call_arguments", "string_value",
		"array_value", "matrix_value", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 405, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 94, 8, 1,
		10, 1, 12, 1, 97, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 104, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 112, 8, 3, 10, 3, 12, 3, 115,
		9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 134, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 3, 8, 141, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 164, 8, 15, 10, 15, 12, 15, 167, 9,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 187, 8,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 209, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 255, 8, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 272, 8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		3, 30, 301, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 308, 8, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 3, 32, 323, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 334, 8, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 345, 8, 34, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 3, 35, 352, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 3, 36, 359, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 3, 37, 369, 8, 37, 1, 38, 1, 38, 1, 38, 3, 38, 374, 8, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 5, 39, 381, 8, 39, 10, 39, 12, 39, 384, 9,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 0,
		0, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 0, 5, 1, 0, 25, 28, 2, 0, 12, 12, 29, 30,
		1, 0, 14, 15, 1, 0, 38, 39, 1, 0, 60, 62, 404, 0, 88, 1, 0, 0, 0, 2, 95,
		1, 0, 0, 0, 4, 98, 1, 0, 0, 0, 6, 108, 1, 0, 0, 0, 8, 116, 1, 0, 0, 0,
		10, 119, 1, 0, 0, 0, 12, 133, 1, 0, 0, 0, 14, 135, 1, 0, 0, 0, 16, 138,
		1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 148, 1, 0, 0,
		0, 24, 151, 1, 0, 0, 0, 26, 155, 1, 0, 0, 0, 28, 157, 1, 0, 0, 0, 30, 161,
		1, 0, 0, 0, 32, 186, 1, 0, 0, 0, 34, 188, 1, 0, 0, 0, 36, 208, 1, 0, 0,
		0, 38, 210, 1, 0, 0, 0, 40, 212, 1, 0, 0, 0, 42, 221, 1, 0, 0, 0, 44, 231,
		1, 0, 0, 0, 46, 233, 1, 0, 0, 0, 48, 254, 1, 0, 0, 0, 50, 256, 1, 0, 0,
		0, 52, 260, 1, 0, 0, 0, 54, 264, 1, 0, 0, 0, 56, 271, 1, 0, 0, 0, 58, 273,
		1, 0, 0, 0, 60, 300, 1, 0, 0, 0, 62, 307, 1, 0, 0, 0, 64, 322, 1, 0, 0,
		0, 66, 333, 1, 0, 0, 0, 68, 344, 1, 0, 0, 0, 70, 351, 1, 0, 0, 0, 72, 358,
		1, 0, 0, 0, 74, 368, 1, 0, 0, 0, 76, 370, 1, 0, 0, 0, 78, 377, 1, 0, 0,
		0, 80, 385, 1, 0, 0, 0, 82, 387, 1, 0, 0, 0, 84, 394, 1, 0, 0, 0, 86, 402,
		1, 0, 0, 0, 88, 89, 3, 2, 1, 0, 89, 90, 5, 0, 0, 1, 90, 1, 1, 0, 0, 0,
		91, 94, 3, 12, 6, 0, 92, 94, 3, 4, 2, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1,
		0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		3, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99, 5, 5, 0, 0, 99, 100, 3, 10,
		5, 0, 100, 101, 5, 62, 0, 0, 101, 103, 5, 1, 0, 0, 102, 104, 3, 6, 3, 0,
		103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105,
		106, 5, 2, 0, 0, 106, 107, 3, 30, 15, 0, 107, 5, 1, 0, 0, 0, 108, 113,
		3, 8, 4, 0, 109, 110, 5, 3, 0, 0, 110, 112, 3, 8, 4, 0, 111, 109, 1, 0,
		0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 7, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 3, 10, 5, 0, 117, 118,
		5, 62, 0, 0, 118, 9, 1, 0, 0, 0, 119, 120, 7, 0, 0, 0, 120, 11, 1, 0, 0,
		0, 121, 134, 3, 32, 16, 0, 122, 134, 3, 48, 24, 0, 123, 134, 3, 36, 18,
		0, 124, 134, 3, 40, 20, 0, 125, 134, 3, 42, 21, 0, 126, 134, 3, 46, 23,
		0, 127, 134, 3, 52, 26, 0, 128, 134, 3, 50, 25, 0, 129, 134, 3, 30, 15,
		0, 130, 134, 3, 28, 14, 0, 131, 134, 3, 14, 7, 0, 132, 134, 3, 24, 12,
		0, 133, 121, 1, 0, 0, 0, 133, 122, 1, 0, 0, 0, 133, 123, 1, 0, 0, 0, 133,
		124, 1, 0, 0, 0, 133, 125, 1, 0, 0, 0, 133, 126, 1, 0, 0, 0, 133, 127,
		1, 0, 0, 0, 133, 128, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 133, 130, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134, 13, 1, 0, 0, 0,
		135, 136, 3, 18, 9, 0, 136, 137, 3, 16, 8, 0, 137, 15, 1, 0, 0, 0, 138,
		140, 3, 20, 10, 0, 139, 141, 3, 22, 11, 0, 140, 139, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 17, 1, 0, 0, 0, 142, 143, 5, 6, 0, 0, 143, 144, 3, 54,
		27, 0, 144, 19, 1, 0, 0, 0, 145, 146, 5, 7, 0, 0, 146, 147, 3, 30, 15,
		0, 147, 21, 1, 0, 0, 0, 148, 149, 5, 8, 0, 0, 149, 150, 3, 30, 15, 0, 150,
		23, 1, 0, 0, 0, 151, 152, 5, 9, 0, 0, 152, 153, 3, 54, 27, 0, 153, 154,
		3, 26, 13, 0, 154, 25, 1, 0, 0, 0, 155, 156, 3, 30, 15, 0, 156, 27, 1,
		0, 0, 0, 157, 158, 5, 4, 0, 0, 158, 159, 3, 54, 27, 0, 159, 160, 5, 34,
		0, 0, 160, 29, 1, 0, 0, 0, 161, 165, 5, 10, 0, 0, 162, 164, 3, 12, 6, 0,
		163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169,
		5, 11, 0, 0, 169, 31, 1, 0, 0, 0, 170, 171, 5, 23, 0, 0, 171, 172, 5, 24,
		0, 0, 172, 173, 3, 34, 17, 0, 173, 174, 5, 32, 0, 0, 174, 175, 5, 62, 0,
		0, 175, 176, 5, 34, 0, 0, 176, 187, 1, 0, 0, 0, 177, 178, 5, 23, 0, 0,
		178, 179, 5, 24, 0, 0, 179, 180, 3, 34, 17, 0, 180, 181, 5, 32, 0, 0, 181,
		182, 5, 62, 0, 0, 182, 183, 5, 33, 0, 0, 183, 184, 3, 54, 27, 0, 184, 185,
		5, 34, 0, 0, 185, 187, 1, 0, 0, 0, 186, 170, 1, 0, 0, 0, 186, 177, 1, 0,
		0, 0, 187, 33, 1, 0, 0, 0, 188, 189, 7, 0, 0, 0, 189, 35, 1, 0, 0, 0, 190,
		191, 5, 23, 0, 0, 191, 192, 3, 38, 19, 0, 192, 193, 5, 32, 0, 0, 193, 194,
		5, 62, 0, 0, 194, 195, 5, 31, 0, 0, 195, 196, 3, 54, 27, 0, 196, 197, 5,
		34, 0, 0, 197, 209, 1, 0, 0, 0, 198, 199, 5, 23, 0, 0, 199, 200, 3, 38,
		19, 0, 200, 201, 5, 32, 0, 0, 201, 202, 5, 62, 0, 0, 202, 203, 5, 31, 0,
		0, 203, 204, 3, 54, 27, 0, 204, 205, 5, 33, 0, 0, 205, 206, 3, 80, 40,
		0, 206, 207, 5, 34, 0, 0, 207, 209, 1, 0, 0, 0, 208, 190, 1, 0, 0, 0, 208,
		198, 1, 0, 0, 0, 209, 37, 1, 0, 0, 0, 210, 211, 7, 1, 0, 0, 211, 39, 1,
		0, 0, 0, 212, 213, 5, 35, 0, 0, 213, 214, 5, 19, 0, 0, 214, 215, 5, 62,
		0, 0, 215, 216, 5, 20, 0, 0, 216, 217, 3, 54, 27, 0, 217, 218, 5, 21, 0,
		0, 218, 219, 3, 54, 27, 0, 219, 220, 5, 34, 0, 0, 220, 41, 1, 0, 0, 0,
		221, 222, 5, 23, 0, 0, 222, 223, 3, 44, 22, 0, 223, 224, 5, 32, 0, 0, 224,
		225, 5, 62, 0, 0, 225, 226, 5, 31, 0, 0, 226, 227, 3, 54, 27, 0, 227, 228,
		5, 13, 0, 0, 228, 229, 3, 54, 27, 0, 229, 230, 5, 34, 0, 0, 230, 43, 1,
		0, 0, 0, 231, 232, 7, 2, 0, 0, 232, 45, 1, 0, 0, 0, 233, 234, 5, 35, 0,
		0, 234, 235, 5, 16, 0, 0, 235, 236, 5, 62, 0, 0, 236, 237, 5, 17, 0, 0,
		237, 238, 3, 54, 27, 0, 238, 239, 5, 18, 0, 0, 239, 240, 3, 54, 27, 0,
		240, 241, 5, 21, 0, 0, 241, 242, 3, 54, 27, 0, 242, 243, 5, 34, 0, 0, 243,
		47, 1, 0, 0, 0, 244, 245, 5, 35, 0, 0, 245, 246, 5, 62, 0, 0, 246, 247,
		3, 54, 27, 0, 247, 248, 5, 34, 0, 0, 248, 255, 1, 0, 0, 0, 249, 250, 5,
		35, 0, 0, 250, 251, 5, 62, 0, 0, 251, 252, 3, 80, 40, 0, 252, 253, 5, 34,
		0, 0, 253, 255, 1, 0, 0, 0, 254, 244, 1, 0, 0, 0, 254, 249, 1, 0, 0, 0,
		255, 49, 1, 0, 0, 0, 256, 257, 5, 37, 0, 0, 257, 258, 5, 62, 0, 0, 258,
		259, 5, 34, 0, 0, 259, 51, 1, 0, 0, 0, 260, 261, 5, 36, 0, 0, 261, 262,
		3, 54, 27, 0, 262, 263, 5, 34, 0, 0, 263, 53, 1, 0, 0, 0, 264, 265, 3,
		56, 28, 0, 265, 55, 1, 0, 0, 0, 266, 272, 3, 60, 30, 0, 267, 268, 3, 60,
		30, 0, 268, 269, 3, 58, 29, 0, 269, 270, 3, 56, 28, 0, 270, 272, 1, 0,
		0, 0, 271, 266, 1, 0, 0, 0, 271, 267, 1, 0, 0, 0, 272, 57, 1, 0, 0, 0,
		273, 274, 7, 3, 0, 0, 274, 59, 1, 0, 0, 0, 275, 301, 3, 62, 31, 0, 276,
		277, 3, 62, 31, 0, 277, 278, 5, 41, 0, 0, 278, 279, 3, 60, 30, 0, 279,
		301, 1, 0, 0, 0, 280, 281, 3, 62, 31, 0, 281, 282, 5, 42, 0, 0, 282, 283,
		3, 60, 30, 0, 283, 301, 1, 0, 0, 0, 284, 285, 3, 62, 31, 0, 285, 286, 5,
		43, 0, 0, 286, 287, 3, 60, 30, 0, 287, 301, 1, 0, 0, 0, 288, 289, 3, 62,
		31, 0, 289, 290, 5, 44, 0, 0, 290, 291, 3, 60, 30, 0, 291, 301, 1, 0, 0,
		0, 292, 293, 3, 62, 31, 0, 293, 294, 5, 45, 0, 0, 294, 295, 3, 60, 30,
		0, 295, 301, 1, 0, 0, 0, 296, 297, 3, 62, 31, 0, 297, 298, 5, 46, 0, 0,
		298, 299, 3, 60, 30, 0, 299, 301, 1, 0, 0, 0, 300, 275, 1, 0, 0, 0, 300,
		276, 1, 0, 0, 0, 300, 280, 1, 0, 0, 0, 300, 284, 1, 0, 0, 0, 300, 288,
		1, 0, 0, 0, 300, 292, 1, 0, 0, 0, 300, 296, 1, 0, 0, 0, 301, 61, 1, 0,
		0, 0, 302, 308, 3, 64, 32, 0, 303, 304, 3, 64, 32, 0, 304, 305, 5, 47,
		0, 0, 305, 306, 3, 62, 31, 0, 306, 308, 1, 0, 0, 0, 307, 302, 1, 0, 0,
		0, 307, 303, 1, 0, 0, 0, 308, 63, 1, 0, 0, 0, 309, 323, 3, 66, 33, 0, 310,
		311, 3, 66, 33, 0, 311, 312, 5, 48, 0, 0, 312, 313, 3, 64, 32, 0, 313,
		323, 1, 0, 0, 0, 314, 315, 3, 66, 33, 0, 315, 316, 5, 49, 0, 0, 316, 317,
		3, 64, 32, 0, 317, 323, 1, 0, 0, 0, 318, 319, 3, 66, 33, 0, 319, 320, 5,
		50, 0, 0, 320, 321, 3, 64, 32, 0, 321, 323, 1, 0, 0, 0, 322, 309, 1, 0,
		0, 0, 322, 310, 1, 0, 0, 0, 322, 314, 1, 0, 0, 0, 322, 318, 1, 0, 0, 0,
		323, 65, 1, 0, 0, 0, 324, 334, 3, 68, 34, 0, 325, 326, 3, 68, 34, 0, 326,
		327, 5, 51, 0, 0, 327, 328, 3, 66, 33, 0, 328, 334, 1, 0, 0, 0, 329, 330,
		3, 68, 34, 0, 330, 331, 5, 52, 0, 0, 331, 332, 3, 66, 33, 0, 332, 334,
		1, 0, 0, 0, 333, 324, 1, 0, 0, 0, 333, 325, 1, 0, 0, 0, 333, 329, 1, 0,
		0, 0, 334, 67, 1, 0, 0, 0, 335, 345, 3, 70, 35, 0, 336, 337, 3, 70, 35,
		0, 337, 338, 5, 53, 0, 0, 338, 339, 3, 68, 34, 0, 339, 345, 1, 0, 0, 0,
		340, 341, 3, 70, 35, 0, 341, 342, 5, 54, 0, 0, 342, 343, 3, 68, 34, 0,
		343, 345, 1, 0, 0, 0, 344, 335, 1, 0, 0, 0, 344, 336, 1, 0, 0, 0, 344,
		340, 1, 0, 0, 0, 345, 69, 1, 0, 0, 0, 346, 352, 3, 72, 36, 0, 347, 348,
		3, 72, 36, 0, 348, 349, 5, 55, 0, 0, 349, 350, 3, 70, 35, 0, 350, 352,
		1, 0, 0, 0, 351, 346, 1, 0, 0, 0, 351, 347, 1, 0, 0, 0, 352, 71, 1, 0,
		0, 0, 353, 359, 3, 74, 37, 0, 354, 355, 5, 57, 0, 0, 355, 356, 3, 54, 27,
		0, 356, 357, 5, 58, 0, 0, 357, 359, 1, 0, 0, 0, 358, 353, 1, 0, 0, 0, 358,
		354, 1, 0, 0, 0, 359, 73, 1, 0, 0, 0, 360, 369, 3, 86, 43, 0, 361, 369,
		3, 84, 42, 0, 362, 369, 3, 82, 41, 0, 363, 369, 3, 76, 38, 0, 364, 365,
		5, 52, 0, 0, 365, 369, 3, 54, 27, 0, 366, 367, 5, 56, 0, 0, 367, 369, 3,
		54, 27, 0, 368, 360, 1, 0, 0, 0, 368, 361, 1, 0, 0, 0, 368, 362, 1, 0,
		0, 0, 368, 363, 1, 0, 0, 0, 368, 364, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0,
		369, 75, 1, 0, 0, 0, 370, 371, 5, 62, 0, 0, 371, 373, 5, 1, 0, 0, 372,
		374, 3, 78, 39, 0, 373, 372, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375,
		1, 0, 0, 0, 375, 376, 5, 2, 0, 0, 376, 77, 1, 0, 0, 0, 377, 382, 3, 54,
		27, 0, 378, 379, 5, 3, 0, 0, 379, 381, 3, 54, 27, 0, 380, 378, 1, 0, 0,
		0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383,
		79, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 386, 5, 59, 0, 0, 386, 81, 1,
		0, 0, 0, 387, 388, 5, 21, 0, 0, 388, 389, 5, 20, 0, 0, 389, 390, 5, 22,
		0, 0, 390, 391, 3, 54, 27, 0, 391, 392, 5, 19, 0, 0, 392, 393, 5, 62, 0,
		0, 393, 83, 1, 0, 0, 0, 394, 395, 5, 21, 0, 0, 395, 396, 5, 17, 0, 0, 396,
		397, 3, 54, 27, 0, 397, 398, 5, 18, 0, 0, 398, 399, 3, 54, 27, 0, 399,
		400, 5, 16, 0, 0, 400, 401, 5, 62, 0, 0, 401, 85, 1, 0, 0, 0, 402, 403,
		7, 4, 0, 0, 403, 87, 1, 0, 0, 0, 21, 93, 95, 103, 113, 133, 140, 165, 186,
		208, 254, 271, 300, 307, 322, 333, 344, 351, 358, 368, 373, 382,
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
	kochanowskiParserT__0           = 1
	kochanowskiParserT__1           = 2
	kochanowskiParserT__2           = 3
	kochanowskiParserT__3           = 4
	kochanowskiParserFUNCTION       = 5
	kochanowskiParserIF             = 6
	kochanowskiParserTHEN           = 7
	kochanowskiParserELSE           = 8
	kochanowskiParserWHILE          = 9
	kochanowskiParserSTART_BLOCK    = 10
	kochanowskiParserEND_BLOCK      = 11
	kochanowskiParserSTRING         = 12
	kochanowskiParserBY             = 13
	kochanowskiParserINT32MATRIX    = 14
	kochanowskiParserF32MATRIX      = 15
	kochanowskiParserMATRIX         = 16
	kochanowskiParserUNDER_COLUMN   = 17
	kochanowskiParserROW            = 18
	kochanowskiParserARRAY          = 19
	kochanowskiParserUNDER          = 20
	kochanowskiParserVALUE          = 21
	kochanowskiParserCELL           = 22
	kochanowskiParserDEFINE         = 23
	kochanowskiParserVARIABLE       = 24
	kochanowskiParserINT32          = 25
	kochanowskiParserINT64          = 26
	kochanowskiParserF32            = 27
	kochanowskiParserF64            = 28
	kochanowskiParserINT32ARRAY     = 29
	kochanowskiParserF32ARRAY       = 30
	kochanowskiParserWITH_SIZE      = 31
	kochanowskiParserNAMED          = 32
	kochanowskiParserWITH_VALUE     = 33
	kochanowskiParserDOT            = 34
	kochanowskiParserASSIGN         = 35
	kochanowskiParserPRINT_WORD     = 36
	kochanowskiParserREAD_WORD      = 37
	kochanowskiParserLOGIC_AND      = 38
	kochanowskiParserLOGIC_OR       = 39
	kochanowskiParserFROM           = 40
	kochanowskiParserGREATER        = 41
	kochanowskiParserGREATEREQUAL   = 42
	kochanowskiParserEQUAL          = 43
	kochanowskiParserLESSEQUAL      = 44
	kochanowskiParserLESS           = 45
	kochanowskiParserNOTEQUAL       = 46
	kochanowskiParserMODULO         = 47
	kochanowskiParserAND            = 48
	kochanowskiParserOR             = 49
	kochanowskiParserXOR            = 50
	kochanowskiParserPLUS           = 51
	kochanowskiParserMINUS          = 52
	kochanowskiParserTIMES          = 53
	kochanowskiParserDIVIDE         = 54
	kochanowskiParserPOWER          = 55
	kochanowskiParserNOT            = 56
	kochanowskiParserFIRST          = 57
	kochanowskiParserCALCULATE      = 58
	kochanowskiParserSTRING_LITERAL = 59
	kochanowskiParserINTEGER        = 60
	kochanowskiParserDECIMAL        = 61
	kochanowskiParserID             = 62
	kochanowskiParserWS             = 63
)

// kochanowskiParser rules.
const (
	kochanowskiParserRULE_prog             = 0
	kochanowskiParserRULE_body             = 1
	kochanowskiParserRULE_function_decl    = 2
	kochanowskiParserRULE_param_list       = 3
	kochanowskiParserRULE_param            = 4
	kochanowskiParserRULE_func_type        = 5
	kochanowskiParserRULE_statement        = 6
	kochanowskiParserRULE_if               = 7
	kochanowskiParserRULE_conditional_body = 8
	kochanowskiParserRULE_if_expr          = 9
	kochanowskiParserRULE_if_body          = 10
	kochanowskiParserRULE_else_body        = 11
	kochanowskiParserRULE_while            = 12
	kochanowskiParserRULE_while_body       = 13
	kochanowskiParserRULE_return           = 14
	kochanowskiParserRULE_block            = 15
	kochanowskiParserRULE_var_create       = 16
	kochanowskiParserRULE_type             = 17
	kochanowskiParserRULE_array_create     = 18
	kochanowskiParserRULE_array_type       = 19
	kochanowskiParserRULE_array_assign     = 20
	kochanowskiParserRULE_matrix_create    = 21
	kochanowskiParserRULE_matrix_type      = 22
	kochanowskiParserRULE_matrix_assign    = 23
	kochanowskiParserRULE_var_assign       = 24
	kochanowskiParserRULE_read             = 25
	kochanowskiParserRULE_print            = 26
	kochanowskiParserRULE_expr             = 27
	kochanowskiParserRULE_expr_logic       = 28
	kochanowskiParserRULE_logic_operator   = 29
	kochanowskiParserRULE_expr_compare     = 30
	kochanowskiParserRULE_expr_mod         = 31
	kochanowskiParserRULE_expr_bit         = 32
	kochanowskiParserRULE_expr_add         = 33
	kochanowskiParserRULE_expr_mult        = 34
	kochanowskiParserRULE_expr_power       = 35
	kochanowskiParserRULE_expr_paren       = 36
	kochanowskiParserRULE_unary            = 37
	kochanowskiParserRULE_function_call    = 38
	kochanowskiParserRULE_call_arguments   = 39
	kochanowskiParserRULE_string_value     = 40
	kochanowskiParserRULE_array_value      = 41
	kochanowskiParserRULE_matrix_value     = 42
	kochanowskiParserRULE_value            = 43
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Body() IBodyContext
	EOF() antlr.TerminalNode

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

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserEOF, 0)
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
		p.SetState(88)
		p.Body()
	}
	{
		p.SetState(89)
		p.Match(kochanowskiParserEOF)
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

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFunction_decl() []IFunction_declContext
	Function_decl(i int) IFunction_declContext

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

func (s *BodyContext) AllFunction_decl() []IFunction_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_declContext); ok {
			len++
		}
	}

	tst := make([]IFunction_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_declContext); ok {
			tst[i] = t.(IFunction_declContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Function_decl(i int) IFunction_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declContext); ok {
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

	return t.(IFunction_declContext)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240526558832) != 0 {
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case kochanowskiParserT__3, kochanowskiParserIF, kochanowskiParserWHILE, kochanowskiParserSTART_BLOCK, kochanowskiParserDEFINE, kochanowskiParserASSIGN, kochanowskiParserPRINT_WORD, kochanowskiParserREAD_WORD:
			{
				p.SetState(91)
				p.Statement()
			}

		case kochanowskiParserFUNCTION:
			{
				p.SetState(92)
				p.Function_decl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(97)
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

// IFunction_declContext is an interface to support dynamic dispatch.
type IFunction_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	Func_type() IFunc_typeContext
	ID() antlr.TerminalNode
	Block() IBlockContext
	Param_list() IParam_listContext

	// IsFunction_declContext differentiates from other interfaces.
	IsFunction_declContext()
}

type Function_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declContext() *Function_declContext {
	var p = new(Function_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_function_decl
	return p
}

func InitEmptyFunction_declContext(p *Function_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_function_decl
}

func (*Function_declContext) IsFunction_declContext() {}

func NewFunction_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declContext {
	var p = new(Function_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_function_decl

	return p
}

func (s *Function_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserFUNCTION, 0)
}

func (s *Function_declContext) Func_type() IFunc_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_typeContext)
}

func (s *Function_declContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Function_declContext) Block() IBlockContext {
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

func (s *Function_declContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Function_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterFunction_decl(s)
	}
}

func (s *Function_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitFunction_decl(s)
	}
}

func (p *kochanowskiParser) Function_decl() (localctx IFunction_declContext) {
	localctx = NewFunction_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, kochanowskiParserRULE_function_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(kochanowskiParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Func_type()
	}
	{
		p.SetState(100)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(kochanowskiParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0 {
		{
			p.SetState(102)
			p.Param_list()
		}

	}
	{
		p.SetState(105)
		p.Match(kochanowskiParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Block()
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

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *Param_listContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterParam_list(s)
	}
}

func (s *Param_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitParam_list(s)
	}
}

func (p *kochanowskiParser) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kochanowskiParserRULE_param_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Param()
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kochanowskiParserT__2 {
		{
			p.SetState(109)
			p.Match(kochanowskiParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Param()
		}

		p.SetState(115)
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Func_type() IFunc_typeContext
	ID() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Func_type() IFunc_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_typeContext)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *kochanowskiParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kochanowskiParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Func_type()
	}
	{
		p.SetState(117)
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

// IFunc_typeContext is an interface to support dynamic dispatch.
type IFunc_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	F32() antlr.TerminalNode
	F64() antlr.TerminalNode

	// IsFunc_typeContext differentiates from other interfaces.
	IsFunc_typeContext()
}

type Func_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_typeContext() *Func_typeContext {
	var p = new(Func_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_func_type
	return p
}

func InitEmptyFunc_typeContext(p *Func_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_func_type
}

func (*Func_typeContext) IsFunc_typeContext() {}

func NewFunc_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_typeContext {
	var p = new(Func_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_func_type

	return p
}

func (s *Func_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_typeContext) INT32() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT32, 0)
}

func (s *Func_typeContext) INT64() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserINT64, 0)
}

func (s *Func_typeContext) F32() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF32, 0)
}

func (s *Func_typeContext) F64() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserF64, 0)
}

func (s *Func_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterFunc_type(s)
	}
}

func (s *Func_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitFunc_type(s)
	}
}

func (p *kochanowskiParser) Func_type() (localctx IFunc_typeContext) {
	localctx = NewFunc_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kochanowskiParserRULE_func_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
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
	Return_() IReturnContext
	If_() IIfContext
	While() IWhileContext

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

func (s *StatementContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatementContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *StatementContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
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
	p.EnterRule(localctx, 12, kochanowskiParserRULE_statement)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Var_create()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Var_assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Array_create()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.Array_assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.Matrix_create()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.Matrix_assign()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(127)
			p.Print_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(128)
			p.Read()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(129)
			p.Block()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(130)
			p.Return_()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(131)
			p.If_()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(132)
			p.While()
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_expr() IIf_exprContext
	Conditional_body() IConditional_bodyContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) If_expr() IIf_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_exprContext)
}

func (s *IfContext) Conditional_body() IConditional_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditional_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditional_bodyContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *kochanowskiParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kochanowskiParserRULE_if)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.If_expr()
	}
	{
		p.SetState(136)
		p.Conditional_body()
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

// IConditional_bodyContext is an interface to support dynamic dispatch.
type IConditional_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_body() IIf_bodyContext
	Else_body() IElse_bodyContext

	// IsConditional_bodyContext differentiates from other interfaces.
	IsConditional_bodyContext()
}

type Conditional_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional_bodyContext() *Conditional_bodyContext {
	var p = new(Conditional_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_conditional_body
	return p
}

func InitEmptyConditional_bodyContext(p *Conditional_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_conditional_body
}

func (*Conditional_bodyContext) IsConditional_bodyContext() {}

func NewConditional_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional_bodyContext {
	var p = new(Conditional_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_conditional_body

	return p
}

func (s *Conditional_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Conditional_bodyContext) If_body() IIf_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_bodyContext)
}

func (s *Conditional_bodyContext) Else_body() IElse_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_bodyContext)
}

func (s *Conditional_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterConditional_body(s)
	}
}

func (s *Conditional_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitConditional_body(s)
	}
}

func (p *kochanowskiParser) Conditional_body() (localctx IConditional_bodyContext) {
	localctx = NewConditional_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kochanowskiParserRULE_conditional_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.If_body()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kochanowskiParserELSE {
		{
			p.SetState(139)
			p.Else_body()
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

// IIf_exprContext is an interface to support dynamic dispatch.
type IIf_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext

	// IsIf_exprContext differentiates from other interfaces.
	IsIf_exprContext()
}

type If_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_exprContext() *If_exprContext {
	var p = new(If_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if_expr
	return p
}

func InitEmptyIf_exprContext(p *If_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if_expr
}

func (*If_exprContext) IsIf_exprContext() {}

func NewIf_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_exprContext {
	var p = new(If_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_if_expr

	return p
}

func (s *If_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *If_exprContext) IF() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserIF, 0)
}

func (s *If_exprContext) Expr() IExprContext {
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

func (s *If_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterIf_expr(s)
	}
}

func (s *If_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitIf_expr(s)
	}
}

func (p *kochanowskiParser) If_expr() (localctx IIf_exprContext) {
	localctx = NewIf_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kochanowskiParserRULE_if_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(kochanowskiParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Expr()
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

// IIf_bodyContext is an interface to support dynamic dispatch.
type IIf_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THEN() antlr.TerminalNode
	Block() IBlockContext

	// IsIf_bodyContext differentiates from other interfaces.
	IsIf_bodyContext()
}

type If_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_bodyContext() *If_bodyContext {
	var p = new(If_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if_body
	return p
}

func InitEmptyIf_bodyContext(p *If_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_if_body
}

func (*If_bodyContext) IsIf_bodyContext() {}

func NewIf_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_bodyContext {
	var p = new(If_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_if_body

	return p
}

func (s *If_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *If_bodyContext) THEN() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserTHEN, 0)
}

func (s *If_bodyContext) Block() IBlockContext {
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

func (s *If_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterIf_body(s)
	}
}

func (s *If_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitIf_body(s)
	}
}

func (p *kochanowskiParser) If_body() (localctx IIf_bodyContext) {
	localctx = NewIf_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, kochanowskiParserRULE_if_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(kochanowskiParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Block()
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

// IElse_bodyContext is an interface to support dynamic dispatch.
type IElse_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElse_bodyContext differentiates from other interfaces.
	IsElse_bodyContext()
}

type Else_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_bodyContext() *Else_bodyContext {
	var p = new(Else_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_else_body
	return p
}

func InitEmptyElse_bodyContext(p *Else_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_else_body
}

func (*Else_bodyContext) IsElse_bodyContext() {}

func NewElse_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_bodyContext {
	var p = new(Else_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_else_body

	return p
}

func (s *Else_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_bodyContext) ELSE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserELSE, 0)
}

func (s *Else_bodyContext) Block() IBlockContext {
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

func (s *Else_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterElse_body(s)
	}
}

func (s *Else_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitElse_body(s)
	}
}

func (p *kochanowskiParser) Else_body() (localctx IElse_bodyContext) {
	localctx = NewElse_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, kochanowskiParserRULE_else_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(kochanowskiParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Block()
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

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	While_body() IWhile_bodyContext

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_while
	return p
}

func InitEmptyWhileContext(p *WhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_while
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserWHILE, 0)
}

func (s *WhileContext) Expr() IExprContext {
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

func (s *WhileContext) While_body() IWhile_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_bodyContext)
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (p *kochanowskiParser) While() (localctx IWhileContext) {
	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, kochanowskiParserRULE_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(kochanowskiParserWHILE)
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
		p.While_body()
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

// IWhile_bodyContext is an interface to support dynamic dispatch.
type IWhile_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsWhile_bodyContext differentiates from other interfaces.
	IsWhile_bodyContext()
}

type While_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_bodyContext() *While_bodyContext {
	var p = new(While_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_while_body
	return p
}

func InitEmptyWhile_bodyContext(p *While_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_while_body
}

func (*While_bodyContext) IsWhile_bodyContext() {}

func NewWhile_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_bodyContext {
	var p = new(While_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_while_body

	return p
}

func (s *While_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *While_bodyContext) Block() IBlockContext {
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

func (s *While_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterWhile_body(s)
	}
}

func (s *While_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitWhile_body(s)
	}
}

func (p *kochanowskiParser) While_body() (localctx IWhile_bodyContext) {
	localctx = NewWhile_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, kochanowskiParserRULE_while_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Block()
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

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	DOT() antlr.TerminalNode

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) Expr() IExprContext {
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

func (s *ReturnContext) DOT() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserDOT, 0)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *kochanowskiParser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kochanowskiParserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(kochanowskiParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Expr()
	}
	{
		p.SetState(159)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	START_BLOCK() antlr.TerminalNode
	END_BLOCK() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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

func (s *BlockContext) END_BLOCK() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserEND_BLOCK, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
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

func (s *BlockContext) Statement(i int) IStatementContext {
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
	p.EnterRule(localctx, 30, kochanowskiParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(kochanowskiParserSTART_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240526558800) != 0 {
		{
			p.SetState(162)
			p.Statement()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
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
	p.EnterRule(localctx, 32, kochanowskiParserRULE_var_create)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Type_()
		}
		{
			p.SetState(173)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(kochanowskiParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Type_()
		}
		{
			p.SetState(180)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Expr()
		}
		{
			p.SetState(184)
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
	p.EnterRule(localctx, 34, kochanowskiParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
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
	p.EnterRule(localctx, 36, kochanowskiParserRULE_array_create)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Array_type()
		}
		{
			p.SetState(192)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(kochanowskiParserWITH_SIZE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Expr()
		}
		{
			p.SetState(196)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(kochanowskiParserDEFINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Array_type()
		}
		{
			p.SetState(200)
			p.Match(kochanowskiParserNAMED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(kochanowskiParserWITH_SIZE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Expr()
		}
		{
			p.SetState(204)
			p.Match(kochanowskiParserWITH_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.String_value()
		}
		{
			p.SetState(206)
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
	p.EnterRule(localctx, 38, kochanowskiParserRULE_array_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1610616832) != 0) {
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
	p.EnterRule(localctx, 40, kochanowskiParserRULE_array_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(kochanowskiParserUNDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Expr()
	}
	{
		p.SetState(217)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Expr()
	}
	{
		p.SetState(219)
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
	p.EnterRule(localctx, 42, kochanowskiParserRULE_matrix_create)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(kochanowskiParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Matrix_type()
	}
	{
		p.SetState(223)
		p.Match(kochanowskiParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(kochanowskiParserWITH_SIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Expr()
	}
	{
		p.SetState(227)
		p.Match(kochanowskiParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Expr()
	}
	{
		p.SetState(229)
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
	p.EnterRule(localctx, 44, kochanowskiParserRULE_matrix_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
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
	p.EnterRule(localctx, 46, kochanowskiParserRULE_matrix_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(kochanowskiParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(kochanowskiParserUNDER_COLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Expr()
	}
	{
		p.SetState(238)
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Expr()
	}
	{
		p.SetState(240)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Expr()
	}
	{
		p.SetState(242)
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
	p.EnterRule(localctx, 48, kochanowskiParserRULE_var_assign)
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.Match(kochanowskiParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Expr()
		}
		{
			p.SetState(247)
			p.Match(kochanowskiParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.Match(kochanowskiParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(kochanowskiParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.String_value()
		}
		{
			p.SetState(252)
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
	p.EnterRule(localctx, 50, kochanowskiParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(kochanowskiParserREAD_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 52, kochanowskiParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(kochanowskiParserPRINT_WORD)
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
	p.EnterRule(localctx, 54, kochanowskiParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
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
	p.EnterRule(localctx, 56, kochanowskiParserRULE_expr_logic)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Expr_compare()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Expr_compare()
		}
		{
			p.SetState(268)
			p.Logic_operator()
		}
		{
			p.SetState(269)
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
	p.EnterRule(localctx, 58, kochanowskiParserRULE_logic_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
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
	p.EnterRule(localctx, 60, kochanowskiParserRULE_expr_compare)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Expr_mod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Expr_mod()
		}
		{
			p.SetState(277)
			p.Match(kochanowskiParserGREATER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Expr_compare()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(280)
			p.Expr_mod()
		}
		{
			p.SetState(281)
			p.Match(kochanowskiParserGREATEREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Expr_compare()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(284)
			p.Expr_mod()
		}
		{
			p.SetState(285)
			p.Match(kochanowskiParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Expr_compare()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(288)
			p.Expr_mod()
		}
		{
			p.SetState(289)
			p.Match(kochanowskiParserLESSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Expr_compare()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(292)
			p.Expr_mod()
		}
		{
			p.SetState(293)
			p.Match(kochanowskiParserLESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Expr_compare()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(296)
			p.Expr_mod()
		}
		{
			p.SetState(297)
			p.Match(kochanowskiParserNOTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
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
	p.EnterRule(localctx, 62, kochanowskiParserRULE_expr_mod)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Expr_bit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.Expr_bit()
		}
		{
			p.SetState(304)
			p.Match(kochanowskiParserMODULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
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
	p.EnterRule(localctx, 64, kochanowskiParserRULE_expr_bit)
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Expr_add()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Expr_add()
		}
		{
			p.SetState(311)
			p.Match(kochanowskiParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Expr_bit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(314)
			p.Expr_add()
		}
		{
			p.SetState(315)
			p.Match(kochanowskiParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Expr_bit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(318)
			p.Expr_add()
		}
		{
			p.SetState(319)
			p.Match(kochanowskiParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
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
	p.EnterRule(localctx, 66, kochanowskiParserRULE_expr_add)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Expr_mult()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.Expr_mult()
		}
		{
			p.SetState(326)
			p.Match(kochanowskiParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Expr_add()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(329)
			p.Expr_mult()
		}
		{
			p.SetState(330)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
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
	p.EnterRule(localctx, 68, kochanowskiParserRULE_expr_mult)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Expr_power()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Expr_power()
		}
		{
			p.SetState(337)
			p.Match(kochanowskiParserTIMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Expr_mult()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Expr_power()
		}
		{
			p.SetState(341)
			p.Match(kochanowskiParserDIVIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
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
	p.EnterRule(localctx, 70, kochanowskiParserRULE_expr_power)
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)
			p.Expr_paren()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Expr_paren()
		}
		{
			p.SetState(348)
			p.Match(kochanowskiParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
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
	p.EnterRule(localctx, 72, kochanowskiParserRULE_expr_paren)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kochanowskiParserVALUE, kochanowskiParserMINUS, kochanowskiParserNOT, kochanowskiParserINTEGER, kochanowskiParserDECIMAL, kochanowskiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Unary()
		}

	case kochanowskiParserFIRST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(kochanowskiParserFIRST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Expr()
		}
		{
			p.SetState(356)
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
	Function_call() IFunction_callContext
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

func (s *UnaryContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
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
	p.EnterRule(localctx, 74, kochanowskiParserRULE_unary)
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(361)
			p.Matrix_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(362)
			p.Array_value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(363)
			p.Function_call()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(364)
			p.Match(kochanowskiParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Expr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(366)
			p.Match(kochanowskiParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
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

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Call_arguments() ICall_argumentsContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) ID() antlr.TerminalNode {
	return s.GetToken(kochanowskiParserID, 0)
}

func (s *Function_callContext) Call_arguments() ICall_argumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_argumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_argumentsContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *kochanowskiParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, kochanowskiParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(kochanowskiParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(kochanowskiParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8291126913991180288) != 0 {
		{
			p.SetState(372)
			p.Call_arguments()
		}

	}
	{
		p.SetState(375)
		p.Match(kochanowskiParserT__1)
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

// ICall_argumentsContext is an interface to support dynamic dispatch.
type ICall_argumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsCall_argumentsContext differentiates from other interfaces.
	IsCall_argumentsContext()
}

type Call_argumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_argumentsContext() *Call_argumentsContext {
	var p = new(Call_argumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_call_arguments
	return p
}

func InitEmptyCall_argumentsContext(p *Call_argumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kochanowskiParserRULE_call_arguments
}

func (*Call_argumentsContext) IsCall_argumentsContext() {}

func NewCall_argumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_argumentsContext {
	var p = new(Call_argumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kochanowskiParserRULE_call_arguments

	return p
}

func (s *Call_argumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_argumentsContext) AllExpr() []IExprContext {
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

func (s *Call_argumentsContext) Expr(i int) IExprContext {
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

func (s *Call_argumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_argumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_argumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.EnterCall_arguments(s)
	}
}

func (s *Call_argumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kochanowskiListener); ok {
		listenerT.ExitCall_arguments(s)
	}
}

func (p *kochanowskiParser) Call_arguments() (localctx ICall_argumentsContext) {
	localctx = NewCall_argumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, kochanowskiParserRULE_call_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Expr()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kochanowskiParserT__2 {
		{
			p.SetState(378)
			p.Match(kochanowskiParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Expr()
		}

		p.SetState(384)
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
	p.EnterRule(localctx, 80, kochanowskiParserRULE_string_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
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
	p.EnterRule(localctx, 82, kochanowskiParserRULE_array_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(kochanowskiParserUNDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.Match(kochanowskiParserCELL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Expr()
	}
	{
		p.SetState(391)
		p.Match(kochanowskiParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 84, kochanowskiParserRULE_matrix_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(kochanowskiParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(kochanowskiParserUNDER_COLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Expr()
	}
	{
		p.SetState(397)
		p.Match(kochanowskiParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Expr()
	}
	{
		p.SetState(399)
		p.Match(kochanowskiParserMATRIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
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
	p.EnterRule(localctx, 86, kochanowskiParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070450532247928832) != 0) {
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
