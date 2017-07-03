package control

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (l *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	l.defaultOffset = reader.ReadInt32()
	l.npairs = reader.ReadInt32()
	l.matchOffsets = reader.ReadInt32s(l.npairs * 2)
}

func (l *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < l.npairs*2; i += 2 {
		if l.matchOffsets[i] == key {
			offset := l.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(l.defaultOffset))
}
