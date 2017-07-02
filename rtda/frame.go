package rtda

// 栈帧
type Frame struct {
	lower        *Frame        // 后续的栈帧队列
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
