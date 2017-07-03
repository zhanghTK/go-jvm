package control

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32   // 默认情况下执行跳转所需的字节码偏移量
	low           int32   // case取值范围
	high          int32   // case取值范围
	jumpOffsets   []int32 // 索引表，范围又low,high确定
}

func (t *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	t.defaultOffset = reader.ReadInt32()
	t.low = reader.ReadInt32()
	t.high = reader.ReadInt32()
	jumpOffsetsCount := t.high - t.low + 1
	t.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (t *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= t.low && index <= t.high {
		offset = int(t.jumpOffsets[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}

	base.Branch(frame, offset)
}
