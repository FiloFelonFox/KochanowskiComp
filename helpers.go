package main

import (
	"fmt"
	"strconv"
)

var typeAlignMap = map[string]string{
	"i1":    "align 1",
	"i32":   "align 4",
	"i64":   "align 8",
	"float": "align 4",
	"double": "align 8",
}

var typeToOrd = map[string]int{
	"i1":    	1,
	"i32":   	2,
	"i64":   	3,
	"float": 	4,
	"double":	5,
}

var ordToType = map[int]string{
	1 : "i1",
	2 : "i32",
	3 : "i64",
	4 : "float",
	5 : "double",
}

func nextVar() string {
	varCount++
	return "%" + fmt.Sprint(varCount)
}

func isIntType(_type string) bool {
	return _type[0] == 'i'
}

func isFloatType(_type string) bool {
	return _type == "float" || _type == "double"
}

func intRank(_type string) (int, error) {
	str, err := strconv.Atoi(string([]rune(_type)[1:]))
	if err != nil {
		return 0, fmt.Errorf("Compiler error: Could not convert type value " +_type + " to rank")
	}
	return str, err
}

//TODO: Rewrite
func castType(v t_var, targetType string) (t_var, error) {
	if v._type == targetType {
		return v, nil
	}

	dst := nextVar()

	if isIntType(v._type) && isIntType(targetType) {
		rank1, err := intRank(v._type) 
		if err != nil {
			return t_var{}, err
		}
		rank2, err := intRank(targetType)
		if err != nil {
			return t_var{}, err
		}
		if rank1 < rank2 {
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
	val, err := env.context.varStack.peek()
	if err != nil {
		return err
	}

	v, err := castType(val, targetType)
	if err != nil {
		return err
	}

	return env.context.varStack.updateLast(v._value, v._type)
}

func castSecondTo(targetType string) error {
	val, err := env.context.varStack.peekSecond()
	if err != nil {
		return err
	}

	v, err := castType(val, targetType)
	if err != nil {
		return err
	}
	return env.context.varStack.updateSecond(v._value, v._type)
}

func matchLastTwoTypes() error {
	last, err := env.context.varStack.peek()
	if err != nil {
		return err
	}
	secondLast, err := env.context.varStack.peekSecond()
	if err != nil {
		return err
	}
	if last._type == secondLast._type {
		return nil
	}
	commonOrd  := max(typeToOrd[last._type], typeToOrd[secondLast._type])
	common := ordToType[commonOrd]
	
	err = castLastTo(common)
	if err != nil {
		return err
	}
	return castSecondTo(common)
}

func lastTwoTo(_type string) error {
	err := castLastTo(_type)
	if err != nil {
		return err
	}
	return castSecondTo(_type)
}