// Code generated from kochanowski.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kochanowski

import "github.com/antlr4-go/antlr/v4"

// kochanowskiListener is a complete listener for a parse tree produced by kochanowskiParser.
type kochanowskiListener interface {
	antlr.ParseTreeListener

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVar_create is called when entering the var_create production.
	EnterVar_create(c *Var_createContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArray_create is called when entering the array_create production.
	EnterArray_create(c *Array_createContext)

	// EnterArray_type is called when entering the array_type production.
	EnterArray_type(c *Array_typeContext)

	// EnterArray_assign is called when entering the array_assign production.
	EnterArray_assign(c *Array_assignContext)

	// EnterMatrix_create is called when entering the matrix_create production.
	EnterMatrix_create(c *Matrix_createContext)

	// EnterMatrix_type is called when entering the matrix_type production.
	EnterMatrix_type(c *Matrix_typeContext)

	// EnterMatrix_assign is called when entering the matrix_assign production.
	EnterMatrix_assign(c *Matrix_assignContext)

	// EnterVar_assign is called when entering the var_assign production.
	EnterVar_assign(c *Var_assignContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr_logic is called when entering the expr_logic production.
	EnterExpr_logic(c *Expr_logicContext)

	// EnterLogic_operator is called when entering the logic_operator production.
	EnterLogic_operator(c *Logic_operatorContext)

	// EnterExpr_compare is called when entering the expr_compare production.
	EnterExpr_compare(c *Expr_compareContext)

	// EnterExpr_mod is called when entering the expr_mod production.
	EnterExpr_mod(c *Expr_modContext)

	// EnterExpr_bit is called when entering the expr_bit production.
	EnterExpr_bit(c *Expr_bitContext)

	// EnterExpr_add is called when entering the expr_add production.
	EnterExpr_add(c *Expr_addContext)

	// EnterExpr_mult is called when entering the expr_mult production.
	EnterExpr_mult(c *Expr_multContext)

	// EnterExpr_power is called when entering the expr_power production.
	EnterExpr_power(c *Expr_powerContext)

	// EnterExpr_paren is called when entering the expr_paren production.
	EnterExpr_paren(c *Expr_parenContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterArray_value is called when entering the array_value production.
	EnterArray_value(c *Array_valueContext)

	// EnterMatrix_value is called when entering the matrix_value production.
	EnterMatrix_value(c *Matrix_valueContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVar_create is called when exiting the var_create production.
	ExitVar_create(c *Var_createContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArray_create is called when exiting the array_create production.
	ExitArray_create(c *Array_createContext)

	// ExitArray_type is called when exiting the array_type production.
	ExitArray_type(c *Array_typeContext)

	// ExitArray_assign is called when exiting the array_assign production.
	ExitArray_assign(c *Array_assignContext)

	// ExitMatrix_create is called when exiting the matrix_create production.
	ExitMatrix_create(c *Matrix_createContext)

	// ExitMatrix_type is called when exiting the matrix_type production.
	ExitMatrix_type(c *Matrix_typeContext)

	// ExitMatrix_assign is called when exiting the matrix_assign production.
	ExitMatrix_assign(c *Matrix_assignContext)

	// ExitVar_assign is called when exiting the var_assign production.
	ExitVar_assign(c *Var_assignContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr_logic is called when exiting the expr_logic production.
	ExitExpr_logic(c *Expr_logicContext)

	// ExitLogic_operator is called when exiting the logic_operator production.
	ExitLogic_operator(c *Logic_operatorContext)

	// ExitExpr_compare is called when exiting the expr_compare production.
	ExitExpr_compare(c *Expr_compareContext)

	// ExitExpr_mod is called when exiting the expr_mod production.
	ExitExpr_mod(c *Expr_modContext)

	// ExitExpr_bit is called when exiting the expr_bit production.
	ExitExpr_bit(c *Expr_bitContext)

	// ExitExpr_add is called when exiting the expr_add production.
	ExitExpr_add(c *Expr_addContext)

	// ExitExpr_mult is called when exiting the expr_mult production.
	ExitExpr_mult(c *Expr_multContext)

	// ExitExpr_power is called when exiting the expr_power production.
	ExitExpr_power(c *Expr_powerContext)

	// ExitExpr_paren is called when exiting the expr_paren production.
	ExitExpr_paren(c *Expr_parenContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitArray_value is called when exiting the array_value production.
	ExitArray_value(c *Array_valueContext)

	// ExitMatrix_value is called when exiting the matrix_value production.
	ExitMatrix_value(c *Matrix_valueContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
