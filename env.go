package main

import (
	"KochanowskiComp/parser"
	"fmt"
)

var typeAlignMap = map[string]string{
	"i1":    "align 1",
	"i32":   "align 4",
	"i64":   "align 8",
	"float": "align 4",
	"double": "align 8",
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
}

var prog string = ""

var variables = make(map[string]t_var)
var arrays = make(map[string]t_array)
var matrices = make(map[string]t_matrix)
var strings = make(map[string]string)

var strCount int = 0

var varCount int = 0

var stack VariableStack

type logicFrame struct {
	operator   string
	shortLabel string
	rhsLabel   string
	endLabel   string
}

var logicStack []logicFrame
var logicCount int = 0

type kochanowskiListener struct {
	*parser.BasekochanowskiListener
}


func nextVar() string {
	varCount++
	return "%" + fmt.Sprint(varCount)
}

func isIntType(_type string) bool {
	return _type == "i1" || _type == "i32" || _type == "i64"
}

func isFloatType(_type string) bool {
	return _type == "float" || _type == "double"
}

func intRank(_type string) int {
	switch _type {
	case "i1":
		return 1
	case "i32":
		return 32
	case "i64":
		return 64
	default:
		return 0
	}
}

func castType(v t_var, targetType string) (t_var, error) {
	if v._type == targetType {
		return v, nil
	}

	dst := nextVar()

	if isIntType(v._type) && isIntType(targetType) {
		if intRank(v._type) < intRank(targetType) {
			prog += dst + " = zext " + v._type + " " + v._value + " to " + targetType + "\n"
		} else {
			prog += dst + " = trunc " + v._type + " " + v._value + " to " + targetType + "\n"
		}
		return t_var{dst, targetType}, nil
	}

	if isIntType(v._type) && isFloatType(targetType) {
		prog += dst + " = sitofp " + v._type + " " + v._value + " to " + targetType + "\n"
		return t_var{dst, targetType}, nil
	}

	if isFloatType(v._type) && isIntType(targetType) {
		prog += dst + " = fptosi " + v._type + " " + v._value + " to " + targetType + "\n"
		return t_var{dst, targetType}, nil
	}

	if isFloatType(v._type) && isFloatType(targetType) {
		if v._type == "float" && targetType == "double" {
			prog += dst + " = fpext float " + v._value + " to double\n"
		} else if v._type == "double" && targetType == "float" {
			prog += dst + " = fptrunc double " + v._value + " to float\n"
		}
		return t_var{dst, targetType}, nil
	}

	return t_var{}, fmt.Errorf("Nie można przekonwertować typu %s na %s", v._type, targetType)
}

func castLastTo(targetType string) error {
	v, err := castType(stack.peek(), targetType)
	if err != nil {
		return err
	}
	stack.updateLast(v._value, v._type)
	return nil
}

func castSecondTo(targetType string) error {
	v, err := castType(stack.peekSecond(), targetType)
	if err != nil {
		return err
	}
	stack.updateSecond(v._value, v._type)
	return nil
}

func matchLastTwoTypes() {
	last := stack.peek()
	secondLast := stack.peekSecond()
	if last._type == secondLast._type {
		return
	}

	common := "i1"
	switch {
	case last._type  == "double" || secondLast._type == "double":
		common = "double"
	case last._type  == "float" || secondLast._type == "float":
		common = "float"
	case last._type  == "i64" || secondLast._type == "i64":
		common = "i64"
	case last._type  == "i32" || secondLast._type == "i32":
		common = "i32"
	}
	
	castLastTo(common)
	castSecondTo(common)
}

func lastTwoToFloat() {
	castLastTo("float")
	castSecondTo("float")
}

func convertLastToi1() {
	castLastTo("i1")
}

func castIntToI64(v t_var) (t_var, error) {
	if v._type == "i64" {
		return v, nil
	}
	if isIntType(v._type) {
		return castType(v, "i64")
	}
	return t_var{}, fmt.Errorf("%s", "Nie można przekonwertować typu " + v._type + " na i64")
}