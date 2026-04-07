package main

import (
	"fmt"
	"KochanowskiComp/parser"
)

var typeAlignMap = map[string]string{
	"i32": "align 4",
	"float": "align 4",
	"i1": "align 1",
}

type t_var struct {
	_value string
	_type string
}

var prog string = ""

var variables = make(map[string]t_var)
var varCount int = 0

var stack VariableStack

type kochanowskiListener struct {
	*parser.BasekochanowskiListener
}

func matchLastTwoTypes() {
	if stack.peek()._type != stack.peekSecond()._type {
		varCount++;
		if stack.peek()._type == "float"{
			if stack.peekSecond()._type == "i32" {
				prog += "%" + fmt.Sprint(varCount) + " = sitofp i32 " + stack.peekSecond()._value + " to float\n"					
				stack.updateSecond("%" + fmt.Sprint(varCount), "float")
			} else if stack.peekSecond()._type == "i1" {
				prog += "%" + fmt.Sprint(varCount) + " = sitofp i1 " + stack.peekSecond()._value + " to float\n"
				stack.updateSecond("%" + fmt.Sprint(varCount), "float")
			}
		} else if stack.peek()._type == "i32" {				
			if stack.peekSecond()._type == "float" {
				prog += "%" + fmt.Sprint(varCount) + " = sitofp i32 " + stack.peek()._value + " to float\n"
				stack.updateLast("%" + fmt.Sprint(varCount), "float")
			} else if stack.peekSecond()._type == "i1" {
				prog += "%" + fmt.Sprint(varCount) + " = sext i1 " + stack.peekSecond()._value + " to i32\n"
				stack.updateSecond("%" + fmt.Sprint(varCount), "i32")
			}
		} else if stack.peek()._type == "i1" {
			if stack.peekSecond()._type == "float" {
				prog += "%" + fmt.Sprint(varCount) + " = sitofp i1 " + stack.peek()._value + " to float\n"
				stack.updateLast("%" + fmt.Sprint(varCount), "float")
			} else if stack.peekSecond()._type == "i32" {
				prog += "%" + fmt.Sprint(varCount) + " = sext i1 " + stack.peek()._value + " to i32\n"
				stack.updateLast("%" + fmt.Sprint(varCount), "i32")
			}
		}
	}
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
	variables[ctx.ID().GetText()] = t_var{"%" + fmt.Sprint(varCount), ""}
}

func (l *kochanowskiListener) ExitVar_create(ctx *parser.Var_createContext) {
	if ctx.Expr() != nil {
		variable := variables[ctx.ID().GetText()]
		variable._type = stack.peek()._type
		variables[ctx.ID().GetText()] = variable
		prog += "store " + stack.peek()._type + " " + stack.peek()._value + ", " + variable._type + "* " + variable._value + "\n"
		stack.pop()
	}
}

//TYPE
func (l *kochanowskiListener) EnterType(ctx *parser.TypeContext) {
	switch ctx.GetText() {
	case "całkowitą":
		prog += "i32, align 4\n"
		stack.push("i32", "i32")
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
	if ctx.Expr_compare() != nil {
		matchLastTwoTypes()
		varCount++;
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if first._type == "float" {
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
		stack.push("%" + fmt.Sprint(varCount), "i1")
	}
}

//EXPR_MOD
func (l *kochanowskiListener) EnterExpr_mod(ctx *parser.Expr_modContext) {
}

func (l *kochanowskiListener) ExitExpr_mod(ctx *parser.Expr_modContext) {
	if ctx.Expr_mod() != nil {
			matchLastTwoTypes()
			varCount++;
			first := stack.pop()
			second := stack.pop()
			prog += "%" + fmt.Sprint(varCount) + " = "
			if ctx.MODULO() != nil {
				prog += "srem"
			}
			prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
			stack.push("%" + fmt.Sprint(varCount), first._type)
	}
}

//EXPR_BIT
func (l *kochanowskiListener) EnterExpr_bit(ctx *parser.Expr_bitContext) {
}

func (l *kochanowskiListener) ExitExpr_bit(ctx *parser.Expr_bitContext) {
	if ctx.Expr_bit() != nil {
		matchLastTwoTypes()
		varCount++;
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
		stack.push("%" + fmt.Sprint(varCount), first._type)
	}
}

//EXPR_ADD
func (l *kochanowskiListener) EnterExpr_add(ctx *parser.Expr_addContext) {
}

func (l *kochanowskiListener) ExitExpr_add(ctx *parser.Expr_addContext) {
	if ctx.Expr_add() != nil {
		matchLastTwoTypes()
		varCount++;
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if first._type == "float" {
			prog += "f"
		}
		if ctx.PLUS() != nil {
			prog += "add";
		} else if ctx.MINUS() != nil {
			prog += "sub";
		}
		prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
		stack.push("%" + fmt.Sprint(varCount), first._type)
	}
}

//EXPR_MULT

func (l *kochanowskiListener) EnterExpr_mult(ctx *parser.Expr_multContext) {
}

func (l *kochanowskiListener) ExitExpr_mult(ctx *parser.Expr_multContext) {
	if ctx.Expr_mult() != nil {
		matchLastTwoTypes()
		varCount++;
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if first._type == "float" {
			prog += "f"
		}
		if ctx.TIMES() != nil {
			prog += "mul";
		} else if ctx.DIVIDE() != nil {
			prog += "div";
		}
		prog += " " + first._type + " " + first._value + ", " + second._value + "\n"
		stack.push("%" + fmt.Sprint(varCount), first._type)
	}
}

//EXPR_POWER
func (l *kochanowskiListener) EnterExpr_power(ctx *parser.Expr_powerContext) {
}

func (l *kochanowskiListener) ExitExpr_power(ctx *parser.Expr_powerContext) {
	if ctx.Expr_power() != nil {
		matchLastTwoTypes()
		varCount++;
		first := stack.pop()
		second := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = call float @llvm.pow.f32(" + first._type + " " + first._value + ", " + second._type + " " + second._value + ")\n"
		stack.push("%" + fmt.Sprint(varCount), first._type)
	}
}

//EXPR_UNARY
func (l *kochanowskiListener) EnterUnary(ctx *parser.UnaryContext) {
}

func (l *kochanowskiListener) ExitUnary(ctx *parser.UnaryContext) {
	if ctx.Expr() != nil {
		varCount++;
		first := stack.pop()
		prog += "%" + fmt.Sprint(varCount) + " = "
		if first._type == "float" {
			if ctx.MINUS() != nil {
				prog += "fneg float" + first._value + "\n"
				stack.push("%" + fmt.Sprint(varCount), "float")
			} else if ctx.NOT() != nil {
				prog += "fcmp oeq " + first._type + " " + first._value + ", 0.0\n"
				stack.push("%" + fmt.Sprint(varCount), "i1")
			}
		} else {
			if ctx.MINUS() != nil {
				prog += "sub " + first._type + " 0, " + first._value + "\n"
				stack.push("%" + fmt.Sprint(varCount), first._type)
			} else if ctx.NOT() != nil {
				prog += "icmp eq " + first._type + " " + first._value + ", 0\n"
				stack.push("%" + fmt.Sprint(varCount), "i1")
			}
		}
	}
}

//VALUE
func (l *kochanowskiListener) EnterValue(ctx *parser.ValueContext) {
	if ctx.INTEGER() != nil {
		stack.push(fmt.Sprint(ctx.INTEGER().GetText()), "i32")
	} else if ctx.DECIMAL() != nil {
		varCount++;
		prog += "%" + fmt.Sprint(varCount) + " = fptrunc double " + fmt.Sprint(ctx.DECIMAL().GetText()) + " to float\n"
		stack.push("%" + fmt.Sprint(varCount), "float")
	} else if ctx.ID() != nil {
		varCount++;
		prog +="%" + fmt.Sprint(varCount) + " = load " + variables[ctx.ID().GetText()]._type + ", ptr " + variables[ctx.ID().GetText()]._value + ", " + typeAlignMap[variables[ctx.ID().GetText()]._type] + "\n"
		stack.push("%" + fmt.Sprint(varCount), variables[ctx.ID().GetText()]._type)
	}
}

func (l *kochanowskiListener) ExitValue(ctx *parser.ValueContext) {
}



func (l *kochanowskiListener) progprint() string{
	return prog
}