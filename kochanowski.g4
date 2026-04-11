grammar kochanowski;

body : statement*;

statement: var_create | var_assign | array_create | array_assign | print | read;

var_create
	: DEFINE VARIABLE type NAMED ID DOT
	| DEFINE VARIABLE type NAMED ID WITH_VALUE expr DOT;

type: INT32 | INT64 | F32 | F64;

array_create : DEFINE  array_type NAMED ID DOT;

array_type: INT32ARRAY | F32ARRAY;

array_assign: ASSIGN ARRAY ID UNDER expr VALUE expr DOT;

var_assign: ASSIGN ID expr DOT;

read: READ_WORD ID DOT;

print: PRINT_WORD expr DOT;

expr: expr_logic;

expr_logic
	: expr_compare
	| expr_compare logic_operator expr_logic;

logic_operator
	: LOGIC_AND
	| LOGIC_OR;

expr_compare
	: expr_mod 
	| expr_mod GREATER expr_compare
	| expr_mod GREATEREQUAL expr_compare
	| expr_mod EQUAL expr_compare
	| expr_mod LESSEQUAL expr_compare
	| expr_mod LESS expr_compare
	| expr_mod NOTEQUAL expr_compare; 

expr_mod
	: expr_bit 
	| expr_bit MODULO expr_mod;

expr_bit
	: expr_add 
	| expr_add AND expr_bit
	| expr_add OR expr_bit
	| expr_add XOR expr_bit;

expr_add
	: expr_mult
	| expr_mult PLUS expr_add
	| expr_mult MINUS expr_add;

expr_mult
	: expr_power
	| expr_power TIMES expr_mult
	| expr_power DIVIDE expr_mult;

expr_power
	: expr_paren
	| expr_paren POWER expr_power;

expr_paren
	: unary
	| FIRST expr CALCULATE;

unary
	: value
	| array_value
	| MINUS expr
	| NOT expr;

array_value: VALUE UNDER CELL expr ARRAY ID;

value
	: INTEGER
	| DECIMAL
	| ID;


ARRAY: 'tablicy';
UNDER: 'pod';
VALUE: 'wartość';
CELL: 'komórką';

DEFINE: 'Zdefiniuj';
VARIABLE: 'zmienną';
INT32: 'całkowitą';
INT64: 'całkowitą olbrzymiej wagi';
F32: 'zmiennoprzecinkową';
F64: 'zmiennoprzecinkową olbrzymiej precyzji';
INT32ARRAY: 'tablicę liczb całkowitych';
F32ARRAY: 'tablicę liczb zmiennoprzecinkowych';

NAMED: 'o nazwie';
WITH_VALUE: 'o wartości';
DOT: '.';

ASSIGN: 'Przypisz';
PRINT_WORD: 'Wypisz';
READ_WORD: 'Wczytaj';

LOGIC_AND: 'i jednocześnie';
LOGIC_OR: 'lub';
FROM: 'z';
GREATER: 'większe niż';
GREATEREQUAL: 'większe lub równe';
EQUAL: 'równe';
LESSEQUAL: 'mniejsze lub równe';
LESS: 'mniejsze niż';
NOTEQUAL: 'różne od';
MODULO: 'modulo';
AND: 'koniunkcja';
OR: 'alternatywa';
XOR: 'alternatywa wykluczająca';
PLUS: 'plus';
MINUS: 'minus';
TIMES: 'razy';
DIVIDE: 'podzielić przez';
POWER: 'do potęgi';
NOT: 'nie';
FIRST: 'wpierw';
CALCULATE: 'policz';

INTEGER: [0-9]+;
DECIMAL: [0-9]*'.'[0-9]+;
ID: [a-zA-Z0-9_ąćęłńóśźżĄĆĘŁŃÓŚŹŻ]+;

WS: [ \t\r\n]+ -> skip;
