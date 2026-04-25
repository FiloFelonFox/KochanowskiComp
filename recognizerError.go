package main

import (
    "fmt"
    "github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
    *antlr.DefaultErrorListener
    Errors []string
    HasErrors bool
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
    el.HasErrors = true
    errorMsg := fmt.Sprintf("Błąd w linii %d, kolumna %d: %s", line, column, msg)
    el.Errors = append(el.Errors, errorMsg)
}