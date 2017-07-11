package extended

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (i *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (i *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, i.Offset)
	}
}
