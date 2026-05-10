package main

import "fmt"

//GLOBALS
var env Enviroment
var sa = new(SemanticAnalyzer)
var prog string = ""
var strings = make(map[string]string)
var varCount int = 0
var strCount int = 0


//
type Enviroment struct {
	context Context
	parent *Enviroment
}

type Context struct {
	varStack VariableStack
	logicStack LogicStack
	variables map[string]t_var
	arrays map[string]t_array
	matrices map[string]t_matrix
}

type t_var struct {
	_value string
	_type  string
}

type t_array struct {
	_value string
	_type  string
	_rowLen t_var
}

type t_matrix struct {
	_value string
	_type  string
	_colLen t_var
	_rowLen t_var
}

type logicFrame struct {
	operator   string
	shortLabel string
	rhsLabel   string
	endLabel   string
}

func makeEnviroment() Enviroment {
	e := Enviroment{parent: nil, context: Context{}}
	e.context.varStack = VariableStack{size: 0, vars: make([]t_var, 0)}
	e.context.logicStack = LogicStack{size: 0, el: make([]logicFrame, 0)}
	e.context.arrays = make(map[string]t_array)
	e.context.matrices = make(map[string]t_matrix)
	e.context.variables = make(map[string]t_var)
	return e
}

func initEnviroment() {
	e := makeEnviroment()
	env = e
}

func upscopeEnviroment() {
	e := makeEnviroment()
	e.parent = &env
	env = e
}

func downscopeEnviroment() {
	env = *env.parent
}

func (e* Enviroment) getVariable(name string) (t_var, error) {
	for ;e!=nil; {
		v, ok := e.context.variables[name]
		if ok{
			return v, nil
		}
		e = e.parent
	}
	return t_var{}, fmt.Errorf("Compiler error: Could not find variable " +name+ " in Environment")
}

func (e* Enviroment) getArray(name string) (t_array, error) {
	for ;e!=nil; {
		v, ok := e.context.arrays[name]
		if ok{
			return v, nil
		}
		e = e.parent
	}
	return t_array{}, fmt.Errorf("Compiler error: Could not find array " +name+ " in Environment")
}

func (e* Enviroment) getMatrix(name string) (t_matrix, error) {
	for ;e!=nil; {
		v, ok := e.context.matrices[name]
		if ok{
			return v, nil
		}
		e = e.parent
	}
	return t_matrix{}, fmt.Errorf("Compiler error: Could not find matrix " +name+ " in Environment")
}