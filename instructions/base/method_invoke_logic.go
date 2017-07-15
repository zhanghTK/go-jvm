package base

import (
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// 方法调用逻辑
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 创建新的栈帧并推入栈
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	// 参数个数
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
