package emulator

import "github.com/ericr/solanalyzer/sources"

type Stack []*StackFrame

type StackFrame struct {
	Contract  *sources.Contract
	Function  *sources.Function
	Memory    []*Variable
	MemoryMap map[string]*Variable
}

func (s *Stack) CurrentFrame() *StackFrame {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(contract *sources.Contract, function *sources.Function) *StackFrame {
	currentFrame := s.CurrentFrame()
	frame := &StackFrame{
		Contract: contract,
		Function: function,
	}

	if currentFrame == nil {
		frame.Memory = []*Variable{}
		frame.MemoryMap = map[string]*Variable{}
	} else {
		frame.Memory = currentFrame.Memory
		frame.MemoryMap = currentFrame.MemoryMap
	}

	*s = append(*s, frame)

	return frame
}

func (s *Stack) Pop() *StackFrame {
	lastFrame := s.CurrentFrame()

	*s = append((*s)[:len(*s)-1], (*s)[len(*s):]...)

	return lastFrame
}
