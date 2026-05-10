package main

import (
	"KochanowskiComp/parser"
	"fmt"
)

type kochanowskiListener struct {
	*parser.BasekochanowskiListener
}

//PROG
func (l* kochanowskiListener) EnterProg(ctx *parser.ProgContext) {
	initEnviroment()
}

func (l* kochanowskiListener) ExitProg(ctx *parser.ProgContext) {

}

//BLOCK
func (l* kochanowskiListener) EnterBlock(ctx *parser.BlockContext) {
	upscopeEnviroment()
}

func (l* kochanowskiListener) ExitBlock(ctx *parser.BlockContext) {
	downscopeEnviroment()
}

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
	prog += "declare void @llvm.memcpy.p0i8.p0i8.i64(ptr align 1, ptr align 1, i64, i1)\n"
	prog += "\n"
	prog += "@.str.i1 = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n"
	prog += "@.str.i32 = private unnamed_addr constant [4 x i8] c\"%d\\0A\\00\", align 1\n"
	prog += "@.str.i64 = private unnamed_addr constant [5 x i8] c\"%ld\\0A\\00\", align 1\n"
	prog += "@.str.double = private unnamed_addr constant [5 x i8] c\"%lg\\0A\\00\", align 1\n"
	prog += "@.str.ptr = private unnamed_addr constant [4 x i8] c\"%s\\0A\\00\", align 1\n"
	prog +="\n"
	prog += "@.scan.str = private unnamed_addr constant [6 x i8] c\"%255s\\00\", align 1\n"
	prog += "@.read.buf = internal global [256 x i8] zeroinitializer, align 1\n"
	prog += "\n"
	prog += "define dso_local i32 @main() {\n"
}

func (l *kochanowskiListener) ExitBody(ctx *parser.BodyContext) {
	for str := range strings {
		prog = str + " = private unnamed_addr constant [" + fmt.Sprint(len(strings[str])+1) + " x i8] c\"" + strings[str] + "\\00\", align 1\n" + prog
	}
	for arr := range env.context.arrays {
		prog += "call void @free(ptr " + env.context.arrays[arr]._value + ")\n"
	}
	for mat := range env.context.matrices {
		prog += "call void @free(ptr " + env.context.matrices[mat]._value + ")\n"
	}
	prog += "ret i32 0\n}\n"
}

// PRINT
func (l *kochanowskiListener) EnterPrint(ctx *parser.PrintContext) {
}

func (l *kochanowskiListener) ExitPrint(ctx *parser.PrintContext) {
	v, _ := env.context.varStack.pop()
	if v._type == "float" {
		var err error
		v, err = castType(v, "double")
		if err != nil {
			sa.addError(ctx.GetParser().GetError().GetMessage(), ctx.Expr().GetStart().GetLine(), ctx.Expr().GetStart().GetColumn())
		}
	}
	ret := nextVar()
	prog += ret + " = call i32 (ptr, ...) @printf(ptr noundef @.str." + v._type + ", " + v._type + " noundef " + v._value + ")\n"
}

// READ
func (l *kochanowskiListener) EnterRead(ctx *parser.ReadContext) {

}

func (l *kochanowskiListener) ExitRead(ctx *parser.ReadContext) {
	variable := env.context.variables[ctx.ID().GetText()]
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
	case "ptr":
		length := env.context.arrays[ctx.ID().GetText()]._rowLen
		prog += "call void @llvm.memcpy.p0i8.p0i8.i64(ptr align 1 " + variable._value + ", ptr align 1 " + ptr + ", i64 " + length._value + ", i1 false)\n"
	default:
		sa.addError("Nie można wczytać wartości do zmiennej typu " + variable._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}
}

// VAR_ASSIGN
func (l *kochanowskiListener) EnterVar_assign(ctx *parser.Var_assignContext) {
}

func (l *kochanowskiListener) ExitVar_assign(ctx *parser.Var_assignContext) {
	variable := env.context.variables[ctx.ID().GetText()]
	value, _ := env.context.varStack.peek()
	if variable._type == "ptr" {
		if value._type != "ptr" {
			sa.addError("Błąd typów: nie można przypisać wartości typu " + value._type + " do zmiennej typu " + variable._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		}
		prog += "call void @llvm.memcpy.p0i8.p0i8.i64(ptr align 1 " + variable._value + ", ptr align 1 " + value._value + ", i64 " + env.context.arrays[ctx.ID().GetText()]._rowLen._value + ", i1 false)\n"
		env.context.varStack.pop()//
		return
	}
	if variable._type != value._type { //TODO: Better type mismatch error handling
		switch {
		case value._type == "ptr":
			sa.addError("Błąd typów: nie można przypisać wartości typu " + value._type + " do zmiennej typu " + variable._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		case isIntType(variable._type) && isIntType(value._type):
			varRank, _ := intRank(variable._type)
			valRank, _ := intRank(value._type)
			if varRank >= valRank {
				err := castLastTo(variable._type)
				if err != nil {
					sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
				}
			} else {
				sa.addError("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
			}
		case isFloatType(variable._type) && isFloatType(value._type):
				err := castLastTo(variable._type)
				if err != nil {
					sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
				}
		case isIntType(variable._type) && isFloatType(value._type):
			sa.addError("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		case isFloatType(variable._type) && isIntType(value._type):
			err := castLastTo(variable._type)
			if err != nil {
				sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
			}
		}
	}
	v, _ := env.context.varStack.pop()
	prog += "store " + v._type + " " + v._value + ",ptr " + variable._value + "\n"
}

// VAR_CREATE
func (l *kochanowskiListener) EnterVar_create(ctx *parser.Var_createContext) {
	sa.checkVariableExists(ctx.ID().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	v := nextVar()
	prog += v + " = alloca "
	env.context.variables[ctx.ID().GetText()] = t_var{"%" + fmt.Sprint(varCount), ""}
}

func (l *kochanowskiListener) ExitVar_create(ctx *parser.Var_createContext) {
	variable, _ := env.context.variables[ctx.ID().GetText()]
	if ctx.Expr() != nil {
		second, _ := env.context.varStack.peekSecond()
		variable._type = second._type 
		env.context.variables[ctx.ID().GetText()] = variable
		value, _ := env.context.varStack.peek()
		if variable._type != value._type {
			switch {
			case isIntType(variable._type) && isIntType(value._type):
				varRank, _ := intRank(variable._type)
				valRank, _ := intRank(value._type)
				if varRank >= valRank {
					err := castLastTo(variable._type)
					if err != nil {
						sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
					}
				} else {
					sa.addError("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
				}
			case isFloatType(variable._type) && isFloatType(value._type):
				err := castLastTo(variable._type)
				if err != nil {
					sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
				}
			case isIntType(variable._type) && isFloatType(value._type):
				sa.addError("Błąd typów" + variable._value + ":" +  variable._type + " " + value._value + ":" + value._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
			case isFloatType(variable._type) && isIntType(value._type):
				err := castLastTo(variable._type)
				if err != nil {
					sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
				}
			}
		}
		v, _ := env.context.varStack.pop()
		prog += "store " + v._type + " " + v._value + ", ptr " + variable._value + "\n"
	} else {
	peek, _ := env.context.varStack.peek()
	variable._type = peek._type 
	env.context.variables[ctx.ID().GetText()] = variable
	}
	env.context.varStack.pop()//
}

// TYPE
func (l *kochanowskiListener) EnterType(ctx *parser.TypeContext) {
	switch ctx.GetText() {
	case "całkowitą":
		prog += "i32, align 4\n"
		env.context.varStack.push("i32", "i32")
	case "całkowitą olbrzymiej wagi":
		prog += "i64, align 8\n"
		env.context.varStack.push("i64", "i64")
	case "zmiennoprzecinkową":
		prog += "float, align 4\n"
		env.context.varStack.push("float", "float")
	case "zmiennoprzecinkową olbrzymiej precyzji":
		prog += "double, align 8\n"
		env.context.varStack.push("double", "double")
	}
}

func (l *kochanowskiListener) ExitType(ctx *parser.TypeContext) {
	//PASS
}

// ARRAY_CREATE
func (l *kochanowskiListener) EnterArray_create(ctx *parser.Array_createContext) {
}

func (l *kochanowskiListener) ExitArray_create(ctx *parser.Array_createContext) {
	var str t_var
	if ctx.String_value() != nil {
		str, _ = env.context.varStack.pop()
	}

	name := ctx.ID().GetText()
	if env.context.arrays[name]._value != "" {
		sa.addError("Tablica '" + name + "' już istnieje", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}

	elemType := ""
	elemSizeBytes := 0
	switch ctx.Array_type().GetText() {
	case "tablicę liczb całkowitych":
		elemType = "i32"
		elemSizeBytes = 4
	case "tablicę liczb zmiennoprzecinkowych":
		elemType = "float"
		elemSizeBytes = 4
	case "napis":
		elemType = "i8"
		elemSizeBytes = 1
	}
	
	sizeVar, _ := env.context.varStack.pop()
	if !isIntType(sizeVar._type) {
		sa.addError("Błąd typów: rozmiar tablicy musi być typu całkowitego, a otrzymał " + sizeVar._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	size64, err := castType(sizeVar, "i64")
	if err != nil {
		sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	bytes := nextVar()
	prog += bytes + " = mul i64 " + size64._value + ", " + fmt.Sprint(elemSizeBytes) + "\n"

	base := nextVar()
	prog += base + " = call ptr @malloc(i64 " + bytes + ")\n"
	rowLen := nextVar()
	prog += rowLen + " = add i64 " + size64._value + ", 0\n"
	env.context.arrays[name] = t_array{base, elemType, t_var{rowLen, "i64"}}
	env.context.variables[name] = t_var{base, "ptr"}
	if ctx.String_value() != nil {
		if str._type != "ptr" {
			sa.addError("Błąd typów: wartość początkowa tablicy musi być typu napis, a otrzymał " + str._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		}
		prog += "call void @llvm.memcpy.p0i8.p0i8.i64(ptr align 1 " + base + ", ptr align 1 " + str._value + ", i64 " + bytes + ", i1 false)\n"
	}
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
	arr := env.context.arrays[ctx.ID().GetText()]

	value, _ := env.context.varStack.pop()
	index, _ := env.context.varStack.pop()

	if !isIntType(index._type) {
		sa.addError("Błąd typów: indeks tablicy musi być typu całkowitego, a otrzymał " + index._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}
	if isIntType(arr._type) && isFloatType(value._type) {
		sa.addError("Błąd typów: nie można przypisać wartości typu " + value._type + " do tablicy typu " + arr._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	if value._type != arr._type {
		var err error
		value, err = castType(value, arr._type)
		if err != nil {
			sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		}
	}

	index64, err := castType(index, "i64")
	if err != nil {
		sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + arr._type + ", ptr " + arr._value + ", i64 " + index64._value + "\n"
	prog += "store " + arr._type + " " + value._value + ", ptr " + elemPtr + ", " + typeAlignMap[arr._type] + "\n"
}

// ARRAY_VALUE
func (l *kochanowskiListener) EnterArray_value(ctx *parser.Array_valueContext) {
}

func (l *kochanowskiListener) ExitArray_value(ctx *parser.Array_valueContext) {
	arr := env.context.arrays[ctx.ID().GetText()]

	index, _ := env.context.varStack.pop()
	
	if !isIntType(index._type) {
		sa.addError("Błąd typów: indeks tablicy musi być typu całkowitego, a otrzymał " + index._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	index64, err := castType(index, "i64")
	if err != nil {
		sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
	}

	elemPtr := nextVar()
	prog += elemPtr + " = getelementptr inbounds " + arr._type + ", ptr " + arr._value + ", i64 " + index64._value + "\n"
	value := nextVar()
	prog += value + " = load " + arr._type + ", ptr " + elemPtr + ", " + typeAlignMap[arr._type] + "\n"
	env.context.varStack.push(value, arr._type)//
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
    
    rows, _ := env.context.varStack.pop()
    cols, _ := env.context.varStack.pop()
    
    if !isIntType(rows._type) || !isIntType(cols._type) {
        sa.addError("Błąd typów: wymiary macierzy muszą być typu całkowitego", ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

    rows64, err := castType(rows, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }
    cols64, err := castType(cols, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

	colLen := nextVar()
	prog += colLen + " = add i64 " + cols64._value + ", 0\n"

	rowLen := nextVar()
	prog += rowLen + " = add i64 " + rows64._value + ", 0\n"

    // Calculate total elements: rows * cols
    totalElems := nextVar()
    prog += totalElems + " = mul i64 " + rows64._value + ", " + cols64._value + "\n"

    // Calculate total bytes needed
    bytes := nextVar()
    prog += bytes + " = mul i64 " + totalElems + ", " + fmt.Sprint(elemSizeBytes) + "\n"

    base := nextVar()
    prog += base + " = call ptr @malloc(i64 " + bytes + ")\n"
    env.context.matrices[name] = t_matrix{base, elemType, t_var{colLen, "i64"}, t_var{rowLen, "i64"}}
}

// MATRIX_VALUE
func (l *kochanowskiListener) EnterMatrix_value(ctx *parser.Matrix_valueContext) {
}

func (l *kochanowskiListener) ExitMatrix_value(ctx *parser.Matrix_valueContext) {
	mat := env.context.matrices[ctx.ID().GetText()]

    col, _ := env.context.varStack.pop()
    row, _ := env.context.varStack.pop()
    
    if !isIntType(row._type) || !isIntType(col._type) {
		sa.addError("Błąd typów: indeksy macierzy muszą być typu całkowitego", ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

    row64, err := castType(row, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }
    col64, err := castType(col, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

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
	env.context.varStack.push(value, mat._type)
}

// MATRIX_ASSIGN
func (l *kochanowskiListener) EnterMatrix_assign(ctx *parser.Matrix_assignContext) {
}

func (l *kochanowskiListener) ExitMatrix_assign(ctx *parser.Matrix_assignContext) {
	mat := env.context.matrices[ctx.ID().GetText()]

    value, _ := env.context.varStack.pop()
    col, _ := env.context.varStack.pop()
    row, _ := env.context.varStack.pop()

    if !isIntType(row._type) || !isIntType(col._type) {
        sa.addError("Błąd typów: indeksy macierzy muszą być typu całkowitego", ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

    if isIntType(mat._type) && isFloatType(value._type) {
		sa.addError("Błąd typów: nie można przypisać wartości typu " + value._type + " do macierzy typu " + mat._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

    if value._type != mat._type {
		var err error
        value, err = castType(value, mat._type)
        if err != nil {
            sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
        }
    }

    row64, err := castType(row, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }
    col64, err := castType(col, "i64")
    if err != nil {
        sa.addError(err.Error(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
    }

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
		env.context.logicStack.newElement(varCount)
	}
}

func (l *kochanowskiListener) ExitExpr_logic(ctx *parser.Expr_logicContext) {
    if ctx.Expr_logic() != nil {
        castLastTo("i1")
        v := nextVar()
        lastFrame, _ := env.context.logicStack.pop()
        
        prog += "br label %" + lastFrame.endLabel + "\n"
        prog += lastFrame.shortLabel + ":\n"
        prog += "br label %" + lastFrame.endLabel + "\n"
        prog += lastFrame.endLabel + ":\n"
        
        topFrame, err := env.context.logicStack.peek()
        hasParent := err == nil
        
		peek, _ := env.context.varStack.peek()

        if lastFrame.operator == "and" {
            if hasParent {
                prog += v + " = phi i1 [ 0, %" + lastFrame.shortLabel + " ], [ " + peek._value + ", %" + topFrame.endLabel + " ]\n"
            } else {
                prog += v + " = phi i1 [ 0, %" + lastFrame.shortLabel + " ], [ " + peek._value + ", %" + lastFrame.rhsLabel + " ]\n"
            }
        } else if lastFrame.operator == "or" {
            if hasParent {
                prog += v + " = phi i1 [ 1, %" + lastFrame.shortLabel + " ], [ " + peek._value + ", %" + topFrame.endLabel + " ]\n"
            } else {
                prog += v + " = phi i1 [ 1, %" + lastFrame.shortLabel + " ], [ " + peek._value + ", %" + lastFrame.rhsLabel + " ]\n"
            }
        }
        
        env.context.varStack.pop()
        env.context.varStack.pop()
        env.context.varStack.push("%"+fmt.Sprint(varCount), "i1")
    }
}

// LOGIC_OPERATOR
func (l *kochanowskiListener) EnterLogic_operator(ctx *parser.Logic_operatorContext) {
}

func (l *kochanowskiListener) ExitLogic_operator(ctx *parser.Logic_operatorContext) {
	castLastTo("i1")
	logicPeek, _ := env.context.logicStack.peek()
	varPeek, _ := env.context.varStack.peek()
	if ctx.LOGIC_AND() != nil {
		logicPeek.operator = "and"
		prog += "br i1 " + varPeek._value + ", label %" + logicPeek.rhsLabel + ", label %" + logicPeek.shortLabel + "\n"
	} else if ctx.LOGIC_OR() != nil {
		logicPeek.operator = "or"
		prog += "br i1 " + varPeek._value + ", label %" + logicPeek.shortLabel + ", label %" + logicPeek.rhsLabel + "\n"
	}
	prog += logicPeek.rhsLabel + ":\n"
}

// EXPR_COMPARE
func (l *kochanowskiListener) EnterExpr_compare(ctx *parser.Expr_compareContext) {
}

func (l *kochanowskiListener) ExitExpr_compare(ctx *parser.Expr_compareContext) {
	if ctx.Expr_compare() != nil {
		matchLastTwoTypes()
		v := nextVar()
		first, _ := env.context.varStack.pop()
		second, _ := env.context.varStack.pop()
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
		prog += " " + second._type + " " + second._value + ", " + first._value + "\n"
		env.context.varStack.push(v, "i1")
	}
}

// EXPR_MOD
func (l *kochanowskiListener) EnterExpr_mod(ctx *parser.Expr_modContext) {
}

func (l *kochanowskiListener) ExitExpr_mod(ctx *parser.Expr_modContext) {
	if ctx.Expr_mod() != nil {
		matchLastTwoTypes()
		varCount++
		first, _ := env.context.varStack.pop()
		second, _ := env.context.varStack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			prog += "frem"
		} else {
			prog += "srem"
		}
		prog += " " + first._type + " " + second._value + ", " + first._value + "\n"
		env.context.varStack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_BIT
func (l *kochanowskiListener) EnterExpr_bit(ctx *parser.Expr_bitContext) {
}

func (l *kochanowskiListener) ExitExpr_bit(ctx *parser.Expr_bitContext) {
	if ctx.Expr_bit() != nil {
		matchLastTwoTypes()
		peek, _ := env.context.varStack.peek()
		if !isIntType(peek._type) {
			sa.addError("Błąd typów: operator bitowy wymaga typu całkowitego, a otrzymał " + peek._type, ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
		}
		varCount++
		first, _ := env.context.varStack.pop()
		second, _ := env.context.varStack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if ctx.AND() != nil {
			prog += "and"
		} else if ctx.OR() != nil {
			prog += "or"
		} else if ctx.XOR() != nil {
			prog += "xor"
		}
		prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
		env.context.varStack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_ADD
func (l *kochanowskiListener) EnterExpr_add(ctx *parser.Expr_addContext) {
}

func (l *kochanowskiListener) ExitExpr_add(ctx *parser.Expr_addContext) {
	if ctx.Expr_add() != nil {
		matchLastTwoTypes()
		varCount++
		first, _ := env.context.varStack.pop()
		second, _ := env.context.varStack.pop()
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
		env.context.varStack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

//EXPR_MULT

func (l *kochanowskiListener) EnterExpr_mult(ctx *parser.Expr_multContext) {
}

func (l *kochanowskiListener) ExitExpr_mult(ctx *parser.Expr_multContext) {
	if ctx.Expr_mult() != nil {
		matchLastTwoTypes()
		varCount++
		first, _ := env.context.varStack.pop()
		second, _ := env.context.varStack.pop()
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
		env.context.varStack.push("%"+fmt.Sprint(varCount), first._type)
	}
}

// EXPR_POWER
func (l *kochanowskiListener) EnterExpr_power(ctx *parser.Expr_powerContext) {
}

func (l *kochanowskiListener) ExitExpr_power(ctx *parser.Expr_powerContext) {
	if ctx.Expr_power() != nil {
			matchLastTwoTypes()
			peek, _ := env.context.varStack.peek()
			if isIntType(peek._type) {
				castLastTo("float")
				castSecondTo("float")
			}
			varCount++
			first, _ := env.context.varStack.pop()
			second, _ := env.context.varStack.pop()
			if first._type == "double" {
				prog += "%" + fmt.Sprint(varCount) + " = call double @llvm.pow.f64( double " + second._value + ", double " + first._value + ")\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), "double")
			} else {
				prog += "%" + fmt.Sprint(varCount) + " = call float @llvm.pow.f32( float " + second._value + ", float " + first._value + ")\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), "float")
			} 
	}
}

// EXPR_UNARY
func (l *kochanowskiListener) EnterUnary(ctx *parser.UnaryContext) {
}

func (l *kochanowskiListener) ExitUnary(ctx *parser.UnaryContext) {
	if ctx.Expr() != nil {
		varCount++
		first, _ := env.context.varStack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if isFloatType(first._type) {
			if ctx.MINUS() != nil {
				prog += "fneg float" + first._value + "\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), "float")
			} else if ctx.NOT() != nil {
				prog += "fcmp oeq " + first._type + " " + first._value + ", 0.0\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), "i1")
			}
		} else {
			if ctx.MINUS() != nil {
				prog += "sub " + first._type + " 0, " + first._value + "\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), first._type)
			} else if ctx.NOT() != nil {
				prog += "icmp eq " + first._type + " " + first._value + ", 0\n"
				env.context.varStack.push("%"+fmt.Sprint(varCount), "i1")
			}
		}
	}
}

//STRING_VALUE
func (l *kochanowskiListener) EnterString_value(ctx *parser.String_valueContext) {
}

func (l *kochanowskiListener) ExitString_value(ctx *parser.String_valueContext) {
	str := ctx.STRING_LITERAL().GetText()
	str = str[1:len(str)-1]
	strBytes := len(str) + 1

	strName := "@.str.user." + fmt.Sprint(strCount)
	strCount++
	strings[strName] = str

	strVar := nextVar()
	prog += strVar + " = getelementptr inbounds [" + fmt.Sprint(strBytes) + " x i8], ptr " + strName + ", i32 0, i32 0\n"
	env.context.varStack.push(strVar, "ptr")
}

// VALUE
func (l *kochanowskiListener) EnterValue(ctx *parser.ValueContext) {
	if ctx.INTEGER() != nil {
		env.context.varStack.push(fmt.Sprint(ctx.INTEGER().GetText()), "i32")
	} else if ctx.DECIMAL() != nil {
		env.context.varStack.push(fmt.Sprint(ctx.DECIMAL().GetText()), "double")
	} else if ctx.ID() != nil {
			if env.context.variables[ctx.ID().GetText()]._type == "ptr" {
				env.context.varStack.push(env.context.variables[ctx.ID().GetText()]._value, "ptr")
				return
			}
		varCount++
		prog += "%" + fmt.Sprint(varCount) + " = load " + env.context.variables[ctx.ID().GetText()]._type + ", ptr " + env.context.variables[ctx.ID().GetText()]._value + ", " + typeAlignMap[env.context.variables[ctx.ID().GetText()]._type] + "\n"
		env.context.varStack.push("%"+fmt.Sprint(varCount), env.context.variables[ctx.ID().GetText()]._type)
	}
}

func (l *kochanowskiListener) ExitValue(ctx *parser.ValueContext) {
}

func (l *kochanowskiListener) progprint() string {
	return prog
}
