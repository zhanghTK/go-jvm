package rtda

import "GJvm/rtda/heap"

// 栈帧
type Frame struct {
	lower        *Frame        // 后续的栈帧队列
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
func (f *Frame) Thread() *Thread {
	return f.thread
}
func (f *Frame) Method() *heap.Method {
	return f.method
}
func (f *Frame) NextPC() int {
	return f.nextPC
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) RevertNextPC() {
	f.nextPC = f.thread.pc
}
