package constants

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// 将常量从常量池压入栈中
type LDC struct{ base.Index8Instruction }

func (l *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, l.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (l *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, l.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		// 先获取Go字符串转换成Java字符串再压栈
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
		// case *heap.ClassRef:
		// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (l *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(l.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
