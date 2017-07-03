package extended

import (
	"GJvm/rtda"
	"GJvm/instructions/base"
)

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (g *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	g.offset = int(reader.ReadInt32())
}
func (g *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.offset)
}
