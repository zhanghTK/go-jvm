package constants

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// PUSH指令：把一个整形数字（长度比较小）送到到栈顶

// 0x10
// bipush
// 将单字节的常量值(-128~127)推送至栈顶
type BIPUSH struct {
	val int8
}

func (b *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	b.val = reader.ReadInt8()
}
func (b *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(b.val)
	frame.OperandStack().PushInt(i)
}

// 0x11
// sipush
// 将一个短整型常量值(-32768~32767)推送至栈顶
type SIPUSH struct {
	val int16
}

func (s *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	s.val = reader.ReadInt16()
}
func (s *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}
