package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// 获取数组长度
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

// 从操作数栈获得数组引用，返回数组长度
func (a *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
