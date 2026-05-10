package main

import "fmt"

type VariableStack struct {
	vars []t_var
	size int
}

func (s *VariableStack) push(value string, typ string) {
	s.vars = append(s.vars, t_var{value, typ})
	s.size++
}

func (s *VariableStack) pop() (t_var, error) {
	if s.size == 0 {
		return t_var{}, fmt.Errorf("Compiler error: Could not perform pop() on VariableStack: Stack too small")
	}
	s.size--
	top := s.vars[s.size]
	s.vars = s.vars[:s.size]
	return top, nil
}

func (s *VariableStack) peek() (t_var, error) {
	if s.size == 0 {
		return t_var{}, fmt.Errorf("Compiler error: Could not perform peek() on VariableStack: Stack too small")
	}
	return s.vars[s.size-1], nil
}

func (s *VariableStack) peekSecond() (t_var, error) {
	if s.size < 2 {
		return t_var{}, fmt.Errorf("Compiler error: Could not perform peekSecond() on VariableStack: Stack too small")
	}
	return s.vars[s.size-2], nil
}

func (s *VariableStack) updateLast(value string, typ string) error {
	if s.size == 0 {
		return fmt.Errorf("Compiler error: Could not perform updateLast() on VariableStack: Stack too small")
	}
	s.vars[s.size-1] = t_var{value, typ}
	return nil
}

func (s *VariableStack) updateSecond(value string, typ string) error {
	if s.size < 2 {
		return fmt.Errorf("Compiler error: Could not perform updateSecond() on VariableStack: Stack too small")
	}
	s.vars[s.size-2] = t_var{value, typ}
	return nil
}