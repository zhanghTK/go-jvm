package loads

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// 负责把本地doubt型变量的送到栈顶
type DLOAD struct{ base.Index8Instruction }

func (d *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(d.Index))
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (d *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (d *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (d *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (d *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
