grammar kochanowski;

body : statement*;

statement: var_create | var_assign | print | read;

var_create
	: DEFINE VARIABLE type NAMED ID DOT
	| DEFINE VARIABLE type NAMED ID WITH_VALUE expr DOT;

type: INT32;

var_assign: ASSIGN ID expr DOT;

read: READ_WORD ID DOT;

print: PRINT_WORD expr DOT;

expr: expr_compare;

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
	| MINUS expr
	| NOT expr;

value
	: INTEGER
	| DECIMAL
	| ID;


DEFINE: 'Zdefiniuj';
VARIABLE: 'zmienną';
INT32: 'całkowitą';
NAMED: 'o nazwie';
WITH_VALUE: 'o wartości';
DOT: '.';

ASSIGN: 'Przypisz';
PRINT_WORD: 'Wypisz';
READ_WORD: 'Wczytaj';

FROM: 'z';
GREATER: 'większe niż';
GREATEREQUAL: 'większe lub równe';
EQUAL: 'równe';
LESSEQUAL: 'mniejsze lub równe';
LESS: 'mniejsze niż';
NOTEQUAL: 'różne od';
MODULO: 'modulo';
AND: 'i';
OR: 'lub';
XOR: 'wykluczając';
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
ID: [a-zA-Z]+;

WS: [ \t\r\n]+ -> skip;
