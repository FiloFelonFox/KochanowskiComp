package main

import (
	"KochanowskiComp/parser"
	"fmt"
)

var typeAlignMap = map[string]string{
	"i1":    "align 1",
	"i32":   "align 4",
	"i64":   "align 8",
	"float": "align 4",
	"double": "align 8",
}

type t_var struct {
	_value string
	_type  string
}

type t_array struct {
	_value string
	_type  string
}

type t_matrix struct {
	_value string
	_type  string
	_colLen t_var
}

var prog string = ""

var variables = make(map[string]t_var)
var arrays = make(map[string]t_array)
var matrices = make(map[string]t_matrix)
var varCount int = 0

var stack VariableStack

type logicFrame struct {
	operator   string
	shortLabel string
	rhsLabel   string
	endLabel   string
}

var logicStack []logicFrame
var logicCount int = 0

type kochanowskiListener struct {
	*parser.BasekochanowskiListener
}

func nextVar() string {
	varCount++
	return "%" + fmt.Sprint(varCount)
}

func isIntType(_type string) bool {
	return _type == "i1" || _type == "i32" || _type == "i64"
}

func isFloatType(_type string) bool {
	return _type == "float" || _type == "double"
}

func intRank(_type string) int {
	switch _type {
	case "i1":
		return 1
	case "i32":
		return 32
	case "i64":
		return 64
	default:
		return 0
	}
}

func castType(v t_var, targetType string) t_var {
	if v._type == targetType {
		return v
	}

	dst := nextVar()

	if isIntType(v._type) && isIntType(targetType) {
		if intRank(v._type) < intRank(targetType) {
			prog += dst + " = sext " + v._type + " " + v._value + " to " + targetType + "\n"
		} else {
			prog += dst + " = trunc " + v._type + " " + v._value + " to " + targetType + "\n"
		}
		return t_var{dst, targetType}
	}

	if isIntType(v._type) && isFloatType(targetType) {
		prog += dst + " = sitofp " + v._type + " " + v._value + " to " + targetType + "\n"
		return t_var{dst, targetType}
	}

	if isFloatType(v._type) && isIntType(targetType) {
		prog += dst + " = fptosi " + v._type + " " + v._value + " to " + targetType + "\n"
		return t_var{dst, targetType}
	}

	if isFloatType(v._type) && isFloatType(targetType) {
		if v._type == "float" && targetType == "double" {
			prog += dst + " = fpext float " + v._value + " to double\n"
		} else if v._type == "double" && targetType == "float" {
			prog += dst + " = fptrunc double " + v._value + " to float\n"
		}
		return t_var{dst, targetType}
	}

	fmt.Println("Nie można przekonwertować typu " + v._type + " na " + targetType)
	panic(1)
}

func castLastTo(targetType string) {
	v := castType(stack.peek(), targetType)
	stack.updateLast(v._value, v._type)
}

func castSecondTo(targetType string) {
	v := castType(stack.peekSecond(), targetType)
	stack.updateSecond(v._value, v._type)
}

func matchLastTwoTypes() {
	last := stack.peek()
	secondLast := stack.peekSecond()
	if last._type == secondLast._type {
		return
	}

	common := "i1"
	switch {
	case last._type  == "double" || secondLast._type == "double":
		common = "double"
	case last._type  == "float" || secondLast._type == "float":
		common = "float"
	case last._type  == "i64" || secondLast._type == "i64":
		common = "i64"
	case last._type  == "i32" || secondLast._type == "i32":
		common = "i32"
	}
	
	castLastTo(common)
	castSecondTo(common)
}

func lastTwoToFloat() {
	castLastTo("float")
	castSecondTo("float")
}

func convertLastToi1() {
	castLastTo("i1")
}

func castIntToI64(v t_var) t_var {
	if v._type == "i64" {
		return v
	}
	if isIntType(v._type) {
		return castType(v, "i64")
	}
	fmt.Println("Nie można przekonwertować typu " + v._type + " na i32")
	panic(1)
}

/*func emitArrayElemPtr(a t_array, indexVar t_var) string {
	i := castIntToI32(indexVar)
	ptr := nextVar()
	prog += ptr + " = getelementptr inbounds [" + fmt.Sprint(a._size) + " x " + a._type + "], ptr " + a._value + ", i32 0, i32 " + i._value + "\n"
	return ptr
}*/

// BODY
func (l *kochanowskiListener) EnterBody(ctx *parser.BodyContext) {
	prog += "declare i32 @printf(ptr noundef, ...)\n"
	prog += "declare i32 @__isoc99_scanf(ptr noundef, ...)\n"
	prog += "declare i32 @atoi(ptr noundef)\n"
	prog += "declare double @atof(ptr noundef)\n"
	prog += "declare float @llvm.pow.f32(float, float)\n"
	prog += "declare double @llvm.pow.f64(double, double)\n"
	prog += "declare ptr @malloc(i64)\n"
	prog += "declare void @free(ptr)\n"
	prog += "\n"
	prog += "@.str.i1 = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n"
	prog += "@.str.i32 = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n"
	prog += "@.str.i64 = private unnamed_addr constant [5 x i8] c\"%ld\\0A\\00\", align 1\n"
	prog += "@.str.double = private unnamed_addr constant [5 x i8] c\"%lg\\0A\\00\", align 1\n"
	prog +="\n"
	prog += "@.scan.str = private unnamed_addr constant [6 x i8] c\"%255s\\00\", align 1\n"
	prog += "@.read.buf = internal global [256 x i8] zeroinitializer, align 1\n"
	prog += "\n"
	prog += "define dso_local i32 @main() {\n"
}

func (l *kochanowskiListener) ExitBody(ctx *parser.BodyContext) {
	for arr := range arrays {
		prog += "call void @free(ptr " + arrays[arr]._value + ")\n"
	}
	for mat := range matrices {
		prog += "call void @free(ptr " + matrices[mat]._value + ")\n"
	}
	prog += "ret i32 0\n}\n"
}

// PRINT
func (l *kochanowskiListener) EnterPrint(ctx *parser.PrintContext) {
}

func (l *kochanowskiListener) ExitPrint(ctx *parser.PrintContext) {
	v := stack.pop()
	if v._type == "float" {
		v = castType(v, "double")
	}
	ret := nextVar()
	prog += ret + " = call i32 (ptr, ...) @printf(ptr noundef @.str." + v._type + ", " + v._type + " noundef " + v._value + ")\n"
}

// READ
func (l *kochanowskiListener) EnterRead(ctx *parser.ReadContext) {

}

func (l *kochanowskiListener) ExitRead(ctx *parser.ReadContext) {
	variable := variables[ctx.ID().GetText()]
	ptr := nextVar()
	prog +=  ptr + " = getelementptr inbounds [256 x i8], ptr @.read.buf, i32 0, i32 0\n"
	scan := nextVar()
	prog += scan + " = call i32 (ptr, ...) @__isoc99_scanf(ptr noundef @.scan.str, ptr noundef " + ptr + ")\n"
	switch variable._type {
	case "i1":
		val := nextVar()
		prog += val + " = call i32 @atoi(ptr noundef " + ptr + ")\n"
		converted := nextVar() 
		prog += converted + "trunc i32 " + val + " to i1\n"
		prog += "store i1 " + converted + ", ptr " + variable._value + ", align 1\n"
	case "i32":
		val := nextVar()
		prog += val + " = call i32 @atoi(ptr noundef " + ptr + ")\n"
		prog += "store i32 " + val + ", ptr " + variable._value + ", align 4\n"
	case "i64":
		val := nextVar()
		prog += val + " = call i32 @atoi(ptr noundef " + ptr + ")\n"
		converted := nextVar()
		prog += converted + " = sext i32 " + val + " to i64\n"
		prog += "store i64 " + converted + ", ptr " + variable._value + ", align 8\n"
	case "float":
		val := nextVar()
		prog += val + " = call double @atof(ptr noundef " + ptr + ")\n"
		converted := nextVar()
		prog += converted + " = fptrunc double " + val + " to float\n"
		prog += "store float " + converted + ", ptr " + variable._value + ", align 4\n"
	case "double":
		val := nextVar()
		prog += val + " = call double @atof(ptr noundef " + ptr + ")\n"
		prog += "store double " + val + ", ptr " + variable._value + ", align 8\n"
	default:
		fmt.Println("Nie można wczytać tego typu")
		panic(1)
	}
}

// VAR_ASSIGN
func (l *kochanowskiListener) EnterVar_assign(ctx *parser.Var_assignContext) {
}

func (l *kochanowskiListener) ExitVar_assign(ctx *parser.Var_assignContext) {
	variable := variables[ctx.ID().GetText()]
	value := stack.peek()
	if variable._type != value._type { //TODO: Better type mismatch error handling
		switch {
		case isIntType(variable._type) && isIntType(value._type):
			if intRank(variable._type) >= intRank(value._type) {
				castLastTo(variable._type)
			} else {
				fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
				panic(1)
			}
		case isFloatType(variable._type) && isFloatType(value._type):
			if variable._type == "double" && value._type == "float" {
				castLastTo(variable._type)
			} else if variable._type == "float" && stack.peek()._type == "double" {
				fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
				panic(1)
			}
		case isIntType(variable._type) && isFloatType(value._type):
			fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
			panic(1)
		case isFloatType(variable._type) && isIntType(value._type):
			castLastTo(variable._type)
		}
	}
	v := stack.pop()
	prog += "store " + v._type + " " + v._value + ",ptr " + variable._value + "\n"
}

// VAR_CREATE
func (l *kochanowskiListener) EnterVar_create(ctx *parser.Var_createContext) {
	//TODO: check for existing variables
	v := nextVar()
	prog += v + " = alloca "
	variables[ctx.ID().GetText()] = t_var{"%" + fmt.Sprint(varCount), ""}
}

func (l *kochanowskiListener) ExitVar_create(ctx *parser.Var_createContext) {
	variable := variables[ctx.ID().GetText()]
	if ctx.Expr() != nil {
		variable._type = stack.peekSecond()._type 
		variables[ctx.ID().GetText()] = variable
		value := stack.peek()
		fmt.Println("Przypisywana wartość: " + value._value + " typu " + value._type)
		fmt.Println("Zmienna: " + variable._value + " typu " + variable._type)
		if variable._type != value._type { //TODO: Better type mismatch error handling
			switch {
			case isIntType(variable._type) && isIntType(value._type):
				if intRank(variable._type) >= intRank(value._type) {
					castLastTo(variable._type)
				} else {
					fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
					panic(1)
				}
			case isFloatType(variable._type) && isFloatType(value._type):
				if variable._type == "double" && value._type == "float" {
					castLastTo(variable._type)
				} else if variable._type == "float" && stack.peek()._type == "double" {
					fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
					panic(1)
				}
			case isIntType(variable._type) && isFloatType(value._type):
				fmt.Println("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type)
				panic(1)
			case isFloatType(variable._type) && isIntType(value._type):
				castLastTo(variable._type)
			}
		}
		v := stack.pop()
		prog += "store " + v._type + " " + v._value + ", ptr " + variable._value + "\n"
	} else {
	variable._type = stack.peek()._type 
	variables[ctx.ID().GetText()] = variable
	}
	stack.pop()
	fmt.Println("Zadeklarowano zmienną " + ctx.ID().GetText() + " typu " + variables[ctx.ID().GetText()]._type)
}

// TYPE
func (l *kochanowskiListener) EnterType(ctx *parser.TypeContext) {
	switch ctx.GetText() {
	case "całkowitą":
		prog += "i32, align 4\n"
		stack.push("i32", "i32")
	case "całkowitą olbrzymiej wagi":
		prog += "i64, align 8\n"
		stack.push("i64", "i64")
	case "zmiennoprzecinkową":
		prog += "float, align 4\n"
		stack.push("float", "float")
	case "zmiennoprzecinkową olbrzymiej precyzji":
		prog += "double, align 8\n"
		stack.push("double", "double")
	}
}

func (l *kochanowskiListener) ExitType(ctx *parser.TypeContext) {
	//PASS
}

// ARRAY_CREATE
func (l *kochanowskiListener) EnterArray_create(ctx *parser.Array_createContext) {
}

func (l *kochanowskiListener) ExitArray_create(ctx *parser.Array_createContext) {
	name := ctx.ID().GetText()

	elemType := ""
	elemSizeBytes := 0
	switch ctx.Array_type().GetText() {
	case "tablicę liczb całkowitych":
		elemType = "i32"
		elemSizeBytes = 4
	case "tablicę liczb zmiennoprzecinkowych":
		elemType = "float"
		elemSizeBytes = 4
	}
	
	sizeVar := stack.pop()
	if !isIntType(sizeVar._type) {
		fmt.Println("Błąd typów: rozmiar tablicy musi być typu całkowitego, a otrzymał " + sizeVar._type)
		panic(1)
	}

	size64 := castIntToI64(sizeVar)

	bytes := nextVar()
	prog += bytes + " = mul i64 " + size64._value + ", " + fmt.Sprint(elemSizeBytes) + "\n"

	base := nextVar()
	prog += base + " = call ptr @malloc(i64 " + bytes + ")\n"
	arrays[name] = t_array{base, elemType}
}

// ARRAY_TYPE
func (l *kochanowskiListener) EnterArray_type(ctx *parser.Array_typeContext) {
}

func (l *kochanowskiListener) ExitArray_type(ctx *parser.Array_typeContext) {
}

// ARRAY_ASSIGN
func (l *kochanowskiListener) EnterArray_assign(ctx *parser.Array_assignContext) {
}

func (l *kochanowskiListener) ExitArray_assign(ctx *parser.Array_assignContext) {
	arr := arrays[ctx.ID().GetText()]

	value := stack.pop()
	index := stack.pop()

	if !isIntType(index._type) {
		fmt.Println("Błąd typów: indeks tablicy musi być typu całkowitego, a otrzymał " + index._type)
		panic(1)
	}
	if isIntType(arr._type) && isFloatType(value._type) {
		fmt.Println("Błąd typów: nie można przypisać wartości typu " + value._type + " do tablicy typu " + arr._type)
		panic(1)
	}

	if value._type != arr._type {
		value = castType(value, arr._type)
	}

	index64 := castIntToI64(index)

	elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + arr._type + ", ptr " + arr._value + ", i64 " + index64._value + "\n"
	prog += "store " + arr._type + " " + value._value + ", ptr " + elemPtr + ", " + typeAlignMap[arr._type] + "\n"
}

// ARRAY_VALUE
func (l *kochanowskiListener) EnterArray_value(ctx *parser.Array_valueContext) {
}

func (l *kochanowskiListener) ExitArray_value(ctx *parser.Array_valueContext) {
	arr := arrays[ctx.ID().GetText()]

	index := stack.pop()
	
	if !isIntType(index._type) {
		fmt.Println("Błąd typów: indeks tablicy musi być typu całkowitego, a otrzymał " + index._type)
		panic(1)
	}

	index64 := castIntToI64(index)

	elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + arr._type + ", ptr " + arr._value + ", i64 " + index64._value + "\n"
	value := nextVar()
	prog += value + " = load " + arr._type + ", ptr " + elemPtr + ", " + typeAlignMap[arr._type] + "\n"
	stack.push(value, arr._type)
}

// MATRIX_CREATE
func (l *kochanowskiListener) EnterMatrix_create(ctx *parser.Matrix_createContext) {
}

func (l *kochanowskiListener) ExitMatrix_create(ctx *parser.Matrix_createContext) {
	    name := ctx.ID().GetText()

    elemType := ""
    elemSizeBytes := 0
    switch ctx.Matrix_type().GetText() {
    case "macierz liczb całkowitych":
        elemType = "i32"
        elemSizeBytes = 4
    case "macierz liczb zmiennoprzecinkowych":
        elemType = "float"
        elemSizeBytes = 4
    }
    
    rows := stack.pop()
    cols := stack.pop()
    
    if !isIntType(rows._type) || !isIntType(cols._type) {
        fmt.Println("Błąd typów: wymiary macierzy muszą być typu całkowitego")
        panic(1)
    }

    rows64 := castIntToI64(rows)
    cols64 := castIntToI64(cols)

	colLen := nextVar()
	prog += colLen + " = add i64 " + cols64._value + ", 0\n"

    // Calculate total elements: rows * cols
    totalElems := nextVar()
    prog += totalElems + " = mul i64 " + rows64._value + ", " + cols64._value + "\n"

    // Calculate total bytes needed
    bytes := nextVar()
    prog += bytes + " = mul i64 " + totalElems + ", " + fmt.Sprint(elemSizeBytes) + "\n"

    base := nextVar()
    prog += base + " = call ptr @malloc(i64 " + bytes + ")\n"
    matrices[name] = t_matrix{base, elemType, t_var{colLen, "i64"}}
}

// MATRIX_VALUE
func (l *kochanowskiListener) EnterMatrix_value(ctx *parser.Matrix_valueContext) {
}

func (l *kochanowskiListener) ExitMatrix_value(ctx *parser.Matrix_valueContext) {
	mat := matrices[ctx.ID().GetText()]

    col := stack.pop()
    row := stack.pop()
    
    if !isIntType(row._type) || !isIntType(col._type) {
        fmt.Println("Błąd typów: indeksy macierzy muszą być typu całkowitego")
        panic(1)
    }

    row64 := castIntToI64(row)
    col64 := castIntToI64(col)

    // Calculate linear index: row * cols + col
    // Note: you'll need to track column count separately or pass it
	colCount := mat._colLen
	help := nextVar()
	prog += help + " = mul i64 " + row64._value + ", " + colCount._value + "\n"
	linearIndex := nextVar()
	prog += linearIndex + " = add i64 " + help + ", " + col64._value + "\n"

    elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + mat._type + ", ptr " + mat._value + ", i64 " + linearIndex + "\n"
	value := nextVar()
	prog += value + " = load " + mat._type + ", ptr " + elemPtr + ", " + typeAlignMap[mat._type] + "\n"
	stack.push(value, mat._type)
}

// MATRIX_ASSIGN
func (l *kochanowskiListener) EnterMatrix_assign(ctx *parser.Matrix_assignContext) {
}

func (l *kochanowskiListener) ExitMatrix_assign(ctx *parser.Matrix_assignContext) {
	mat := matrices[ctx.ID().GetText()]

    value := stack.pop()
    col := stack.pop()
    row := stack.pop()

    if !isIntType(row._type) || !isIntType(col._type) {
        fmt.Println("Błąd typów: indeksy macierzy muszą być typu całkowitego")
        panic(1)
    }

    if isIntType(mat._type) && isFloatType(value._type) {
        fmt.Println("Błąd typów: nie można przypisać wartości typu " + value._type + " do macierzy typu " + mat._type)
        panic(1)
    }

    if value._type != mat._type {
        value = castType(value, mat._type)
    }

    row64 := castIntToI64(row)
    col64 := castIntToI64(col)

	colCount := mat._colLen
	help := nextVar()
	prog += help + " = mul i64 " + row64._value + ", " + colCount._value + "\n"
	linearIndex := nextVar()
	prog += linearIndex + " = add i64 " + help + ", " + col64._value + "\n"

    elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + mat._type + ", ptr " + mat._value + ", i64 " + linearIndex + "\n"
	prog += "store " + mat._type + " " + value._value + ", ptr " + elemPtr + ", " + typeAlignMap[mat._type] + "\n"
}

// MATRIX_TYPE
func (l *kochanowskiListener) EnterMatrix_type(ctx *parser.Matrix_typeContext) {
}

func (l *kochanowskiListener) ExitMatrix_type(ctx *parser.Matrix_typeContext) {
}

// EXPR
func (l *kochanowskiListener) EnterExpr(ctx *parser.ExprContext) {
}

func (l *kochanowskiListener) ExitExpr(ctx *parser.ExprContext) {
}

// EXPR_LOGIC
func (l *kochanowskiListener) EnterExpr_logic(ctx *parser.Expr_logicContext) {
	if ctx.Expr_logic() != nil {
		varCount++
		frame := logicFrame{
			rhsLabel:   fmt.Sprintf("logic_rhs_%d", varCount),
			endLabel:   fmt.Sprintf("logic_end_%d", varCount),
			shortLabel: fmt.Sprintf("logic_short_%d", varCount),
		}
		logicStack = append(logicStack, frame)
		logicCount++
	}
}

func (l *kochanowskiListener) ExitExpr_logic(ctx *parser.Expr_logicContext) {
	if ctx.Expr_logic() != nil {
		convertLastToi1()
		v := nextVar()
		prog += "br label %" + logicStack[logicCount-1].endLabel + "\n"
		prog += logicStack[logicCount-1].shortLabel + ":\n"
		prog += "br label %" + logicStack[logicCount-1].endLabel + "\n"
		prog += logicStack[logicCount-1].endLabel + ":\n"
		if logicStack[logicCount-1].operator == "and" {
			if logicCount < len(logicStack) && len(logicStack) > 1 {
				prog += v + " = phi i1 [ 0, %" + logicStack[logicCount-1].shortLabel + " ], [ " + stack.peek()._value + ", %" + logicStack[logicCount].endLabel + " ]\n"
			} else {
				prog += v + " = phi i1 [ 0, %" + logicStack[logicCount-1].shortLabel + " ], [ " + stack.peek()._value + ", %" + logicStack[logicCount-1].rhsLabel + " ]\n"
			}
		} else if logicStack[logicCount-1].operator == "or" {
			if logicCount < len(logicStack) && len(logicStack) > 1 {
				prog += v + " = phi i1 [ 1, %" + logicStack[logicCount-1].shortLabel + " ], [ " + stack.peek()._value + ", %" + logicStack[logicCount].endLabel + " ]\n"
			} else {
				prog += v + " = phi i1 [ 1, %" + logicStack[logicCount-1].shortLabel + " ], [ " + stack.peek()._value + ", %" + logicStack[logicCount-1].rhsLabel + " ]\n"
			}
		}
		logicCount--
		stack.pop()
		stack.pop()
		stack.push("%"+fmt.Sprint(varCount), "i1")
	}
}

// LOGIC_OPERATOR
func (l *kochanowskiListener) EnterLogic_operator(ctx *parser.Logic_operatorContext) {
}

func (l *kochanowskiListener) ExitLogic_operator(ctx *parser.Logic_operatorContext) {
	convertLastToi1()
	if ctx.LOGIC_AND() != nil {
		logicStack[len(logicStack)-1].operator = "and"
		prog += "br i1 " + stack.peek()._value + ", label %" + logicStack[len(logicStack)-1].rhsLabel + ", label %" + logicStack[len(logicStack)-1].shortLabel + "\n"
	} else if ctx.LOGIC_OR() != nil {
		logicStack[len(logicStack)-1].operator = "or"
		prog += "br i1 " + stack.peek()._value + ", label %" + logicStack[len(logicStack)-1].shortLabel + ", label %" + logicStack[len(logicStack)-1].rhsLabel + "\n"
	}
	prog += logicStack[len(logicStack)-1].rhsLabel + ":\n"
}

// EXPR_COMPARE
func (l *kochanowskiListener) EnterExpr_compare(ctx *parser.Expr_compareContext) {
}

func (l *kochanowskiListener) ExitExpr_compare(ctx *parser.Expr_compareContext) {
	if ctx.Expr_compare() != nil {
		matchLastTwoTypes()
		v := nextVar()
		first := stack.pop()
		second := stack.pop()
		prog += v + " = "
		if isFloatType(first._type) {
			if ctx.GREATER() != nil {
				prog += "fcmp ogt"
			} else if ctx.LESS() != nil {
				prog += "fcmp olt"
			} else if ctx.EQUAL() != nil {
				prog += "fcmp oeq"
			} else if ctx.NOTEQUAL() != nil {
				prog += "fcmp one"
			} else if ctx.GREATEREQUAL() != nil {
				prog += "fcmp oge"
			} else if ctx.LESSEQUAL() != nil {
				prog += "fcmp ole"
			}
		} else {
			if ctx.GREATER() != nil {
				prog += "icmp sgt"
			} else if ctx.LESS() != nil {
				prog += "icmp slt"
			} else if ctx.EQUAL() != nil {
				prog += "icmp eq"
			} else if ctx.NOTEQUAL() != nil {
				prog += "icmp ne"
			} else if ctx.GREATEREQUAL() != nil {
				prog += "icmp sge"
			} else if ctx.LESSEQUAL() != nil {
				prog += "icmp sle"
			}
		}
		prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
		stack.push(v, "i1")
	}
}

// EXPR_MOD
func (l *kochanowskiListener) EnterExpr_mod(ctx *parser.Expr_modContext) {
}

func (l *kochanowskiListener) ExitExpr_mod(ctx *parser.Expr_modContext) {
	if ctx.Expr_mod() != nil {
		matchLastTwoTypes()
		varCount++
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			prog += "frem"
		} else {
			prog += "srem"
		}
		prog += " " + first._type + " " + second._value + ", " + first._value + "\n"
		stack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_BIT
func (l *kochanowskiListener) EnterExpr_bit(ctx *parser.Expr_bitContext) {
}

func (l *kochanowskiListener) ExitExpr_bit(ctx *parser.Expr_bitContext) {
	if ctx.Expr_bit() != nil {
		matchLastTwoTypes()
		if !isIntType(stack.peek()._type) {
			fmt.Println("Błąd typów: operator bitowy wymaga typu całkowitego, a otrzymał " + stack.peek()._type)
			panic(1)
		}
		varCount++
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if ctx.AND() != nil {
			prog += "and"
		} else if ctx.OR() != nil {
			prog += "or"
		} else if ctx.XOR() != nil {
			prog += "xor"
		}
		prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
		stack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_ADD
func (l *kochanowskiListener) EnterExpr_add(ctx *parser.Expr_addContext) {
}

func (l *kochanowskiListener) ExitExpr_add(ctx *parser.Expr_addContext) {
	if ctx.Expr_add() != nil {
		matchLastTwoTypes()
		varCount++
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			prog += "f"
		}
		if ctx.PLUS() != nil {
			prog += "add"
		} else if ctx.MINUS() != nil {
			prog += "sub"
		}
		prog += " " + first._type + " " + second._value + ", " + first._value + "\n"
		stack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

//EXPR_MULT

func (l *kochanowskiListener) EnterExpr_mult(ctx *parser.Expr_multContext) {
}

func (l *kochanowskiListener) ExitExpr_mult(ctx *parser.Expr_multContext) {
	if ctx.Expr_mult() != nil {
		matchLastTwoTypes()
		varCount++
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			prog += "f"
		}
		if ctx.TIMES() != nil {
			prog += "mul"
		} else if ctx.DIVIDE() != nil {
			prog += "div"
		}
		prog += " " + first._type + " " + second._value + ", " + first._value + "\n"
		stack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_POWER
func (l *kochanowskiListener) EnterExpr_power(ctx *parser.Expr_powerContext) {
}

func (l *kochanowskiListener) ExitExpr_power(ctx *parser.Expr_powerContext) {
	if ctx.Expr_power() != nil {
			matchLastTwoTypes()
			if stack.peek()._type == "i1" || stack.peek()._type == "i32" || stack.peek()._type == "i64" {
				castLastTo("float")
				castSecondTo("float")
			}
			varCount++
			first := stack.pop()
			second := stack.pop()
			if first._type == "double" {
				prog += "%" + fmt.Sprint(varCount) + " = call double @llvm.pow.f64( double " + second._value + ", double " + first._value + ")\n"
				stack.push("%"+fmt.Sprint(varCount), "double")
			} else {
				prog += "%" + fmt.Sprint(varCount) + " = call float @llvm.pow.f32( float " + second._value + ", float " + first._value + ")\n"
				stack.push("%"+fmt.Sprint(varCount), "float")
			} 
	}
}

// EXPR_UNARY
func (l *kochanowskiListener) EnterUnary(ctx *parser.UnaryContext) {
}

func (l *kochanowskiListener) ExitUnary(ctx *parser.UnaryContext) {
	if ctx.Expr() != nil {
		varCount++
		first := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			if ctx.MINUS() != nil {
				prog += "fneg float" + first._value + "\n"
				stack.push("%"+fmt.Sprint(varCount), "float")
			} else if ctx.NOT() != nil {
				prog += "fcmp oeq " + first._type + " " + first._value + ", 0.0\n"
				stack.push("%"+fmt.Sprint(varCount), "i1")
			}
		} else {
			if ctx.MINUS() != nil {
				prog += "sub " + first._type + " 0, " + first._value + "\n"
				stack.push("%"+fmt.Sprint(varCount), first._type)
			} else if ctx.NOT() != nil {
				prog += "icmp eq " + first._type + " " + first._value + ", 0\n"
				stack.push("%"+fmt.Sprint(varCount), "i1")
			}
		}
	}
}

// VALUE
func (l *kochanowskiListener) EnterValue(ctx *parser.ValueContext) {
	if ctx.INTEGER() != nil {
		stack.push(fmt.Sprint(ctx.INTEGER().GetText()), "i32")
	} else if ctx.DECIMAL() != nil {
		stack.push(fmt.Sprint(ctx.DECIMAL().GetText()), "double")
	} else if ctx.ID() != nil {
		varCount++
		prog += "%" + fmt.Sprint(varCount) + " = load " + variables[ctx.ID().GetText()]._type + ", ptr " + variables[ctx.ID().GetText()]._value + ", " + typeAlignMap[variables[ctx.ID().GetText()]._type] + "\n"
		stack.push("%"+fmt.Sprint(varCount), variables[ctx.ID().GetText()]._type)
	}
}

func (l *kochanowskiListener) ExitValue(ctx *parser.ValueContext) {
}

func (l *kochanowskiListener) progprint() string {
	return prog
}
