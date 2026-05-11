// Code generated from kochanowski.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kochanowski

import "github.com/antlr4-go/antlr/v4"

// BasekochanowskiListener is a complete listener for a parse tree produced by kochanowskiParser.
type BasekochanowskiListener struct{}

var _ kochanowskiListener = &BasekochanowskiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekochanowskiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekochanowskiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekochanowskiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekochanowskiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasekochanowskiListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasekochanowskiListener) ExitProg(ctx *ProgContext) {}

// EnterBody is called when production body is entered.
func (s *BasekochanowskiListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasekochanowskiListener) ExitBody(ctx *BodyContext) {}

// EnterFunction_decl is called when production function_decl is entered.
func (s *BasekochanowskiListener) EnterFunction_decl(ctx *Function_declContext) {}

// ExitFunction_decl is called when production function_decl is exited.
func (s *BasekochanowskiListener) ExitFunction_decl(ctx *Function_declContext) {}

// EnterParam_list is called when production param_list is entered.
func (s *BasekochanowskiListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BasekochanowskiListener) ExitParam_list(ctx *Param_listContext) {}

// EnterParam is called when production param is entered.
func (s *BasekochanowskiListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasekochanowskiListener) ExitParam(ctx *ParamContext) {}

// EnterFunc_type is called when production func_type is entered.
func (s *BasekochanowskiListener) EnterFunc_type(ctx *Func_typeContext) {}

// ExitFunc_type is called when production func_type is exited.
func (s *BasekochanowskiListener) ExitFunc_type(ctx *Func_typeContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasekochanowskiListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasekochanowskiListener) ExitStatement(ctx *StatementContext) {}

// EnterIf is called when production if is entered.
func (s *BasekochanowskiListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BasekochanowskiListener) ExitIf(ctx *IfContext) {}

// EnterConditional_body is called when production conditional_body is entered.
func (s *BasekochanowskiListener) EnterConditional_body(ctx *Conditional_bodyContext) {}

// ExitConditional_body is called when production conditional_body is exited.
func (s *BasekochanowskiListener) ExitConditional_body(ctx *Conditional_bodyContext) {}

// EnterIf_expr is called when production if_expr is entered.
func (s *BasekochanowskiListener) EnterIf_expr(ctx *If_exprContext) {}

// ExitIf_expr is called when production if_expr is exited.
func (s *BasekochanowskiListener) ExitIf_expr(ctx *If_exprContext) {}

// EnterIf_body is called when production if_body is entered.
func (s *BasekochanowskiListener) EnterIf_body(ctx *If_bodyContext) {}

// ExitIf_body is called when production if_body is exited.
func (s *BasekochanowskiListener) ExitIf_body(ctx *If_bodyContext) {}

// EnterElse_body is called when production else_body is entered.
func (s *BasekochanowskiListener) EnterElse_body(ctx *Else_bodyContext) {}

// ExitElse_body is called when production else_body is exited.
func (s *BasekochanowskiListener) ExitElse_body(ctx *Else_bodyContext) {}

// EnterWhile is called when production while is entered.
func (s *BasekochanowskiListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BasekochanowskiListener) ExitWhile(ctx *WhileContext) {}

// EnterWhile_body is called when production while_body is entered.
func (s *BasekochanowskiListener) EnterWhile_body(ctx *While_bodyContext) {}

// ExitWhile_body is called when production while_body is exited.
func (s *BasekochanowskiListener) ExitWhile_body(ctx *While_bodyContext) {}

// EnterReturn is called when production return is entered.
func (s *BasekochanowskiListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BasekochanowskiListener) ExitReturn(ctx *ReturnContext) {}

// EnterBlock is called when production block is entered.
func (s *BasekochanowskiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasekochanowskiListener) ExitBlock(ctx *BlockContext) {}

// EnterVar_create is called when production var_create is entered.
func (s *BasekochanowskiListener) EnterVar_create(ctx *Var_createContext) {}

// ExitVar_create is called when production var_create is exited.
func (s *BasekochanowskiListener) ExitVar_create(ctx *Var_createContext) {}

// EnterType is called when production type is entered.
func (s *BasekochanowskiListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasekochanowskiListener) ExitType(ctx *TypeContext) {}

// EnterArray_create is called when production array_create is entered.
func (s *BasekochanowskiListener) EnterArray_create(ctx *Array_createContext) {}

// ExitArray_create is called when production array_create is exited.
func (s *BasekochanowskiListener) ExitArray_create(ctx *Array_createContext) {}

// EnterArray_type is called when production array_type is entered.
func (s *BasekochanowskiListener) EnterArray_type(ctx *Array_typeContext) {}

// ExitArray_type is called when production array_type is exited.
func (s *BasekochanowskiListener) ExitArray_type(ctx *Array_typeContext) {}

// EnterArray_assign is called when production array_assign is entered.
func (s *BasekochanowskiListener) EnterArray_assign(ctx *Array_assignContext) {}

// ExitArray_assign is called when production array_assign is exited.
func (s *BasekochanowskiListener) ExitArray_assign(ctx *Array_assignContext) {}

// EnterMatrix_create is called when production matrix_create is entered.
func (s *BasekochanowskiListener) EnterMatrix_create(ctx *Matrix_createContext) {}

// ExitMatrix_create is called when production matrix_create is exited.
func (s *BasekochanowskiListener) ExitMatrix_create(ctx *Matrix_createContext) {}

// EnterMatrix_type is called when production matrix_type is entered.
func (s *BasekochanowskiListener) EnterMatrix_type(ctx *Matrix_typeContext) {}

// ExitMatrix_type is called when production matrix_type is exited.
func (s *BasekochanowskiListener) ExitMatrix_type(ctx *Matrix_typeContext) {}

// EnterMatrix_assign is called when production matrix_assign is entered.
func (s *BasekochanowskiListener) EnterMatrix_assign(ctx *Matrix_assignContext) {}

// ExitMatrix_assign is called when production matrix_assign is exited.
func (s *BasekochanowskiListener) ExitMatrix_assign(ctx *Matrix_assignContext) {}

// EnterVar_assign is called when production var_assign is entered.
func (s *BasekochanowskiListener) EnterVar_assign(ctx *Var_assignContext) {}

// ExitVar_assign is called when production var_assign is exited.
func (s *BasekochanowskiListener) ExitVar_assign(ctx *Var_assignContext) {}

// EnterRead is called when production read is entered.
func (s *BasekochanowskiListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BasekochanowskiListener) ExitRead(ctx *ReadContext) {}

// EnterPrint is called when production print is entered.
func (s *BasekochanowskiListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BasekochanowskiListener) ExitPrint(ctx *PrintContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasekochanowskiListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasekochanowskiListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr_logic is called when production expr_logic is entered.
func (s *BasekochanowskiListener) EnterExpr_logic(ctx *Expr_logicContext) {}

// ExitExpr_logic is called when production expr_logic is exited.
func (s *BasekochanowskiListener) ExitExpr_logic(ctx *Expr_logicContext) {}

// EnterLogic_operator is called when production logic_operator is entered.
func (s *BasekochanowskiListener) EnterLogic_operator(ctx *Logic_operatorContext) {}

// ExitLogic_operator is called when production logic_operator is exited.
func (s *BasekochanowskiListener) ExitLogic_operator(ctx *Logic_operatorContext) {}

// EnterExpr_compare is called when production expr_compare is entered.
func (s *BasekochanowskiListener) EnterExpr_compare(ctx *Expr_compareContext) {}

// ExitExpr_compare is called when production expr_compare is exited.
func (s *BasekochanowskiListener) ExitExpr_compare(ctx *Expr_compareContext) {}

// EnterExpr_mod is called when production expr_mod is entered.
func (s *BasekochanowskiListener) EnterExpr_mod(ctx *Expr_modContext) {}

// ExitExpr_mod is called when production expr_mod is exited.
func (s *BasekochanowskiListener) ExitExpr_mod(ctx *Expr_modContext) {}

// EnterExpr_bit is called when production expr_bit is entered.
func (s *BasekochanowskiListener) EnterExpr_bit(ctx *Expr_bitContext) {}

// ExitExpr_bit is called when production expr_bit is exited.
func (s *BasekochanowskiListener) ExitExpr_bit(ctx *Expr_bitContext) {}

// EnterExpr_add is called when production expr_add is entered.
func (s *BasekochanowskiListener) EnterExpr_add(ctx *Expr_addContext) {}

// ExitExpr_add is called when production expr_add is exited.
func (s *BasekochanowskiListener) ExitExpr_add(ctx *Expr_addContext) {}

// EnterExpr_mult is called when production expr_mult is entered.
func (s *BasekochanowskiListener) EnterExpr_mult(ctx *Expr_multContext) {}

// ExitExpr_mult is called when production expr_mult is exited.
func (s *BasekochanowskiListener) ExitExpr_mult(ctx *Expr_multContext) {}

// EnterExpr_power is called when production expr_power is entered.
func (s *BasekochanowskiListener) EnterExpr_power(ctx *Expr_powerContext) {}

// ExitExpr_power is called when production expr_power is exited.
func (s *BasekochanowskiListener) ExitExpr_power(ctx *Expr_powerContext) {}

// EnterExpr_paren is called when production expr_paren is entered.
func (s *BasekochanowskiListener) EnterExpr_paren(ctx *Expr_parenContext) {}

// ExitExpr_paren is called when production expr_paren is exited.
func (s *BasekochanowskiListener) ExitExpr_paren(ctx *Expr_parenContext) {}

// EnterUnary is called when production unary is entered.
func (s *BasekochanowskiListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BasekochanowskiListener) ExitUnary(ctx *UnaryContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BasekochanowskiListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BasekochanowskiListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterCall_arguments is called when production call_arguments is entered.
func (s *BasekochanowskiListener) EnterCall_arguments(ctx *Call_argumentsContext) {}

// ExitCall_arguments is called when production call_arguments is exited.
func (s *BasekochanowskiListener) ExitCall_arguments(ctx *Call_argumentsContext) {}

// EnterString_value is called when production string_value is entered.
func (s *BasekochanowskiListener) EnterString_value(ctx *String_valueContext) {}

// ExitString_value is called when production string_value is exited.
func (s *BasekochanowskiListener) ExitString_value(ctx *String_valueContext) {}

// EnterArray_value is called when production array_value is entered.
func (s *BasekochanowskiListener) EnterArray_value(ctx *Array_valueContext) {}

// ExitArray_value is called when production array_value is exited.
func (s *BasekochanowskiListener) ExitArray_value(ctx *Array_valueContext) {}

// EnterMatrix_value is called when production matrix_value is entered.
func (s *BasekochanowskiListener) EnterMatrix_value(ctx *Matrix_valueContext) {}

// ExitMatrix_value is called when production matrix_value is exited.
func (s *BasekochanowskiListener) ExitMatrix_value(ctx *Matrix_valueContext) {}

// EnterValue is called when production value is entered.
func (s *BasekochanowskiListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasekochanowskiListener) ExitValue(ctx *ValueContext) {}
