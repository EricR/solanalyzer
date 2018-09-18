package emulator

type Stack []*StackFrame

type StackFrame struct {
	LocalVariables    []*Variable
	LocalVariablesMap map[string]*Variable
}

func (s *Stack) CurrentFrame() *StackFrame {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Push() *StackFrame {
	var newFrame *StackFrame

	lastFrame := s.CurrentFrame()

	if lastFrame == nil {
		newFrame = &StackFrame{
			LocalVariables:    []*Variable{},
			LocalVariablesMap: map[string]*Variable{},
		}
	} else {
		newFrame = &StackFrame{
			LocalVariables:    lastFrame.LocalVariables,
			LocalVariablesMap: lastFrame.LocalVariablesMap,
		}
	}

	*s = append(*s, newFrame)

	return newFrame
}

func (s *Stack) Pop() *StackFrame {
	lastFrame := s.CurrentFrame()

	*s = append((*s)[:len(*s)-1], (*s)[len(*s):]...)

	return lastFrame
}
