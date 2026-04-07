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

// EnterBody is called when production body is entered.
func (s *BasekochanowskiListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasekochanowskiListener) ExitBody(ctx *BodyContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasekochanowskiListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasekochanowskiListener) ExitStatement(ctx *StatementContext) {}

// EnterVar_create is called when production var_create is entered.
func (s *BasekochanowskiListener) EnterVar_create(ctx *Var_createContext) {}

// ExitVar_create is called when production var_create is exited.
func (s *BasekochanowskiListener) ExitVar_create(ctx *Var_createContext) {}

// EnterType is called when production type is entered.
func (s *BasekochanowskiListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasekochanowskiListener) ExitType(ctx *TypeContext) {}

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

// EnterValue is called when production value is entered.
func (s *BasekochanowskiListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasekochanowskiListener) ExitValue(ctx *ValueContext) {}
