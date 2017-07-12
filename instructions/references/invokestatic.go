package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// Invoke a class (static) method
// 静态方法调用
type INVOKE_STATIC struct{ base.Index16Instruction }

func (i *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	// 解析、获取方法
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(i.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 检测class是否初始化
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}
