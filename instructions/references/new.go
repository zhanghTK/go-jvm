package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// Create new object
type NEW struct{ base.Index16Instruction }

func (n *NEW) Execute(frame *rtda.Frame) {
	// 根据操作数从当前运行时常量池中找到一个类符号的引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	// 解析类
	class := classRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	// 创建对象
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
