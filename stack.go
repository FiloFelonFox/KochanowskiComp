package main

type VariableStack struct {
	vars []t_var
}

func (s *VariableStack) push(value string, typ string) {
	s.vars = append(s.vars, t_var{value, typ})
}

func (s *VariableStack) pop() t_var {
	if len(s.vars) == 0 {
		//TODO: should not happen, error out
		return t_var{"", ""}
	}
	top := s.vars[len(s.vars)-1]
	s.vars = s.vars[:len(s.vars)-1]
	return top
}

func (s *VariableStack) peek() t_var {
	if len(s.vars) == 0 {
		//TODO: should not happen, error out
		return t_var{"", ""}
	}
	return s.vars[len(s.vars)-1]
}

func (s *VariableStack) peekSecond() t_var {
	if len(s.vars) < 2 {
		//TODO: should not happen, error out
		return t_var{"", ""}
	}
	return s.vars[len(s.vars)-2]
}

func (s *VariableStack) updateLast(value string, typ string) {
	if len(s.vars) == 0 {
		//TODO: should not happen, error out
		return
	}
	s.vars[len(s.vars)-1] = t_var{value, typ}
}

func (s *VariableStack) updateSecond(value string, typ string) {
	if len(s.vars) < 2 {
		//TODO: should not happen, error out
		return
	}
	s.vars[len(s.vars)-2] = t_var{value, typ}
}