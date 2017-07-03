package math

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Boolean AND int
type IAND struct{ base.NoOperandsInstruction }

func (i *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperandsInstruction }

func (l *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
