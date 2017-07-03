package comparisons

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (f *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (f *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
