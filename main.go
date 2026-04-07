package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"KochanowskiComp/parser"
)

func main() {
	input := antlr.NewInputStream("7 minus 2")
	
	lexer := parser.NewkochanowskiLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	
	p := parser.NewkochanowskiParser(tokens)
	tree := p.Expr()
	fmt.Println(tree.ToStringTree(nil, p))
}
