package rtda

// 栈帧
type Frame struct {
	lower        *Frame        // 后续的栈帧队列
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread
	nextPC       int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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
func (f *Frame) NextPC() int {
	return f.nextPC
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}
