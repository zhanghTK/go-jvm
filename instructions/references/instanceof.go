package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (io *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// 从操作数栈取第二个操作数，对象引用
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	// 从字节码获取第二个操作数，类符号引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(io.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
