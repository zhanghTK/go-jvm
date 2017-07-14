package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// newarray创建基本类型数组，指令需要两个操作数
// 第一个操作数在指令字节码后面，操作数值表示对应的数组类型，对应关系如下
const (
	//Array Type  atype
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// Create new array
type NEW_ARRAY struct {
	atype uint8
}

func (n *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	n.atype = reader.ReadUint8()
}
func (n *NEW_ARRAY) Execute(frame *rtda.Frame) {
	// 第二个操作数，从操作数栈获取数组长度
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	// 使用当前类加载器数组类
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, n.atype)
	// 创建数组并入栈
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}
