package comparisons

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (i *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (i *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, i.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}
