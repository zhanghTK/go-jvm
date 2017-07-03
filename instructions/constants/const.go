package constants

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// 常量指令系列：把简单的数值类型送到栈顶

// 0x00
// nop
// 什么都不做
type NOP struct{ base.NoOperandsInstruction }

func (n *NOP) Execute(frame *rtda.Frame) {
}

// 0x01
// aconst_null
// 将null推送至栈顶
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (a *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 0x02
// iconst_m1
// 将int型(-1)推送至栈顶
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (i *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 0x03
// iconst_0
// 将int型(0)推送至栈顶
type ICONST_0 struct{ base.NoOperandsInstruction }

func (i *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// 0x04
// iconst_1
// 将int型(1)推送至栈顶
type ICONST_1 struct{ base.NoOperandsInstruction }

func (i *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

// 0x05
// iconst_2
// 将int型(2)推送至栈顶
type ICONST_2 struct{ base.NoOperandsInstruction }

func (i *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

// 0x06
// iconst_3
// 将int型(3)推送至栈顶
type ICONST_3 struct{ base.NoOperandsInstruction }

func (i *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

// 0x07
// iconst_4
// 将int型(4)推送至栈顶
type ICONST_4 struct{ base.NoOperandsInstruction }

func (i *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

// 0x08
// iconst_5
// 将int型(5)推送至栈顶
type ICONST_5 struct{ base.NoOperandsInstruction }

func (i *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// 0x09
// lconst_0
// 将long型(0)推送至栈顶
type LCONST_0 struct{ base.NoOperandsInstruction }

func (l *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

// 0x0a
// lconst_1
// 将long型(1)推送至栈顶
type LCONST_1 struct{ base.NoOperandsInstruction }

func (l *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

// 0x0b
// fconst_0
// 将float型(0)推送至栈顶
type FCONST_0 struct{ base.NoOperandsInstruction }

func (f *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// 0x0c
// fconst_1
// 将float型(1)推送至栈顶
type FCONST_1 struct{ base.NoOperandsInstruction }

func (f *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// 0x0d
// fconst_2
// 将float型(2)推送至栈顶
type FCONST_2 struct{ base.NoOperandsInstruction }

func (f *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 0x0e
// dconst_0
// 将double型(0)推送至栈顶
type DCONST_0 struct{ base.NoOperandsInstruction }

func (d *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 0x0f
// dconst_1
// 将double型(1)推送至栈顶
type DCONST_1 struct{ base.NoOperandsInstruction }

func (d *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}
