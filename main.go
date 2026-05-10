package main

import (
	"fmt"
	"os"
	"KochanowskiComp/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	progPath := os.Args[1]
	outPath := os.Args[2]

	data, err := os.ReadFile(progPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	input := antlr.NewInputStream(string(data))
	
	lexer := parser.NewkochanowskiLexer(input)
	errorListener := &ErrorListener{}
    lexer.RemoveErrorListeners()
    lexer.AddErrorListener(errorListener)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	
	p := parser.NewkochanowskiParser(tokens)
	p.RemoveErrorListeners()
    p.AddErrorListener(errorListener)

	tree := p.Prog()
	if errorListener.HasErrors {
        fmt.Println("\n=== Wystąpiły błędy podczas parsowania ===")
        for _, errMsg := range errorListener.Errors {
            fmt.Println(errMsg)
        }
        return
    }

	listener := &kochanowskiListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if sa.hasErrors {
		fmt.Println("\n=== Wystąpiły błędy podczas analizy semantycznej ===")
		for _, errMsg := range sa.errors {
			fmt.Println(errMsg)
		}
		return
	}

	_, err = outFile.WriteString(listener.progprint())
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
}
