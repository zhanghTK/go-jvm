package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	count uint8
	zero  uint8
}

func (i *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	i.index = uint(reader.ReadUint16())
	i.count = reader.ReadUint8() // 参数个数
	i.zero = reader.ReadUint8()  // 0
}

func (i *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	// 解析获得方法
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(i.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	// 禁止接口中定义的静态&私有方法
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}
	// 未实现接口
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 查找方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
