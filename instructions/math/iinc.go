package math

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint  // 局部变量表索引
	Const int32 // 常量值
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
	i.Const = int32(reader.ReadInt8())
}

func (i *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	val += i.Const
	localVars.SetInt(i.Index, val)
}
