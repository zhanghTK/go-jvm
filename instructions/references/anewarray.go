package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// anewarray创建引用类型数组，指令需要两个操作数
type ANEW_ARRAY struct{ base.Index16Instruction } // 来自字节码的第一个操作数，常量池索引（类符号引用）

func (a *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	// 获取数组元素的类
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	// 从操作数栈取出第二个操作数
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	// 加载数组类
	arrClass := componentClass.ArrayClass()
	// 数组类创建数组并入栈
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
