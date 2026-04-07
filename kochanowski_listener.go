package main

import (
	"fmt"
	"KochanowskiComp/parser"
)

type _var struct {
	_value string
	_type string
}

var prog string = ""

var variables = make(map[string]_var)
var varCount int = 0

var stack []_var

type kochanowskiListener struct {
	*parser.BasekochanowskiListener
}

//BODY
func (l *kochanowskiListener) EnterBody(ctx *parser.BodyContext) {
	prog += "define dso_local i32 @main() {\n"
}

func (l *kochanowskiListener) ExitBody(ctx *parser.BodyContext) {
	prog += "ret i32 0\n}\n"
}

//PRINT
func (l *kochanowskiListener) EnterPrint(ctx *parser.PrintContext) {
}

func (l* kochanowskiListener) ExitPrint(ctx *parser.PrintContext) {
}

//VAR_ASSIGN
func (l *kochanowskiListener) EnterVar_assign(ctx *parser.Var_assignContext) {
}

func (l *kochanowskiListener) ExitVar_assign(ctx *parser.Var_assignContext) {
}

//VAR_CREATE
func (l *kochanowskiListener) EnterVar_create(ctx *parser.Var_createContext) {
	//TODO: check for existing variables
	varCount++;
	prog += "%" + fmt.Sprint(varCount) + " = alloca "
	variables[ctx.ID().GetText()] = _var{"%" + fmt.Sprint(varCount), ""}
}

func (l *kochanowskiListener) ExitVar_create(ctx *parser.Var_createContext) {
	variable := variables[ctx.ID().GetText()]
	variable._type = stack[len(stack)-1]._type
	variables[ctx.ID().GetText()] = variable
	prog += "store " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + variable._type + "* " + variable._value + "\n"
	stack = stack[:len(stack)-1]
}

//TYPE
func (l *kochanowskiListener) EnterType(ctx *parser.TypeContext) {
	switch ctx.GetText() {
	case "całkowitą":
		prog += "i32, align 4\n"
		stack = append(stack, _var{"i32", "i32"})
	}
}

func (l *kochanowskiListener) ExitType(ctx *parser.TypeContext) {
	//PASS
}

//EXPR
func (l *kochanowskiListener) EnterExpr(ctx *parser.ExprContext) {
}

func (l *kochanowskiListener) ExitExpr(ctx *parser.ExprContext) {
}

//EXPR_COMPARE
func (l *kochanowskiListener) EnterExpr_compare(ctx *parser.Expr_compareContext) {
}

func (l *kochanowskiListener) ExitExpr_compare(ctx *parser.Expr_compareContext) {
	if ctx.GREATER() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp sgt " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	} else if ctx.LESS() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp slt " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	} else if ctx.EQUAL() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp eq " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	} else if ctx.NOTEQUAL() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp ne " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	} else if ctx.GREATEREQUAL() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp sge " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	} else if ctx.LESSEQUAL() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = icmp sle " + stack[len(stack)-1]._type + " " + stack[len(stack)-1]._value + ", " + stack[len(stack)-2]._value + "\n"
		stack = stack[:len(stack)-2]
		stack = append(stack, _var{"%" + fmt.Sprint(varCount), "i1"})
	}
}

//VALUE
func (l *kochanowskiListener) EnterValue(ctx *parser.ValueContext) {
	if ctx.INTEGER() != nil {
		stack = append(stack, _var{fmt.Sprint(ctx.INTEGER().GetText()), "i32"})
	} else if ctx.DECIMAL() != nil {
		stack = append(stack, _var{fmt.Sprint(ctx.DECIMAL().GetText()), "float"})
	} else if ctx.ID() != nil {
		stack = append(stack, variables[ctx.ID().GetText()])
	}
}

func (l *kochanowskiListener) ExitValue(ctx *parser.ValueContext) {

}



func (l *kochanowskiListener) progprint() string{
	return prog
}