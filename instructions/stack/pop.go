package stack

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (p *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
