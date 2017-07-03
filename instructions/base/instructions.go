package base

import "GJvm/rtda"

// 指令接口
type Instruction interface {
	// 从字节码中读取操作数
	FetchOperands(reader *BytecodeReader)

	// 执行指令逻辑
	Execute(frame *rtda.Frame)
}

// 没有操作数的指令
type NoOperandsInstruction struct {
}

func (n *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	b.Offset = int(reader.ReadInt16())
}

// 访问局部变量表
type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

// 访问常量值指令
type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())
}
