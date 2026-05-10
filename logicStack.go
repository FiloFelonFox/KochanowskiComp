package main

import "fmt"

type LogicStack struct {
	el []logicFrame
	size int
}

func (s *LogicStack) newElement(ord int) {
	s.el = append(s.el, logicFrame{
			rhsLabel:   fmt.Sprintf("logic_rhs_%d", ord),
			endLabel:   fmt.Sprintf("logic_end_%d", ord),
			shortLabel: fmt.Sprintf("logic_short_%d", ord),
	})
	s.size++
}

func (s *LogicStack) push(operator string, rhsLabel string, endLabel string, shortLabel string) {
	s.el = append(s.el, logicFrame{operator: operator, rhsLabel: rhsLabel, endLabel: endLabel, shortLabel: shortLabel})
	s.size++
}

func (s *LogicStack) pop() (logicFrame, error) {
	if s.size == 0 {
		return logicFrame{}, fmt.Errorf("Compiler error: Could not perform pop() on LogicStack: Stack too small")
	}
	s.size--
	top := s.el[s.size]
	s.el = s.el[:s.size]
	return top, nil
}

func (s *LogicStack) peek() (logicFrame, error) {
	if s.size == 0 {
		return logicFrame{}, fmt.Errorf("Compiler error: Could not perform peek() on LogicStack: Stack too small")
	}
	return s.el[s.size-1], nil
}