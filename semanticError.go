package main

type SemanticError struct {
	Message string
	line int
	col int
}

type SemanticAnalyzer struct {
    errors []SemanticError
    hasErrors bool
}

func (sa *SemanticAnalyzer) checkVariableExists(name string, line int, col int) {
    if _, ok := env.context.variables[name]; ok {
        sa.addError("Zmienna '" + name + "' nie została zadeklarowana", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

// Check variable doesn't exist
func (sa *SemanticAnalyzer) checkVariableNotExists(name string, line int, col int) {
    if _, ok := env.context.variables[name]; !ok {
        sa.addError("Zmienna '" + name + "' już istnieje", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

func (sa *SemanticAnalyzer) checkArrayExists(name string, line int, col int) {
    if _, ok := env.context.arrays[name]; ok {
        sa.addError("Tablica '" + name + "' nie została zadeklarowana", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

func (sa *SemanticAnalyzer) checkArrayNotExists(name string, line int, col int) {
	if _, ok := env.context.arrays[name]; !ok {
		sa.addError("Tablica '" + name + "' nie została zadeklarowana", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

func (sa *SemanticAnalyzer) checkMatrixExists(name string, line int, col int) {
	if _, ok := env.context.matrices[name]; ok {
		sa.addError("Macierz '" + name + "' nie została zadeklarowana", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

func (sa *SemanticAnalyzer) checkMatrixNotExists(name string, line int, col int) {
	if _, ok := env.context.matrices[name]; !ok {
		sa.addError("Macierz '" + name + "' już istnieje", line, col)
		sa.hasErrors = true
	}
	sa.hasErrors = false || sa.hasErrors
}

func (sa *SemanticAnalyzer) addError(msg string, line int, col int) {
    sa.errors = append(sa.errors, SemanticError{Message: msg, line: line, col: col})
    sa.hasErrors = true
}