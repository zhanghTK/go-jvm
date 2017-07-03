package control

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (g *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}
