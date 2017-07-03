package loads

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// 负责把本地int型变量的送到栈顶
type ILOAD struct{ base.Index8Instruction }

func (i *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(i.Index))
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (i *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (i *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (i *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (i *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
