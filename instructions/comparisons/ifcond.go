package comparisons

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (i *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (i *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (i *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (i *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (i *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (i *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, i.Offset)
	}
}
