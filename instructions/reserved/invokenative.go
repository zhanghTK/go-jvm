package reserved

import (
	"GJvm/instructions/base"
	"GJvm/native"
	_ "GJvm/native/java/lang"
	"GJvm/rtda"
)

// 使用预留指令作为本地方法调用指令
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (i *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError:" + methodInfo)
	}

	nativeMethod(frame)
}
