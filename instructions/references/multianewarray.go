package references

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
)

// 创建多维数组
type MULTI_ANEW_ARRAY struct {
	index      uint16 // 类符号引用
	dimensions uint8  // 数组维度
}

func (m *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	m.index = reader.ReadUint16()
	m.dimensions = reader.ReadUint8()
}
func (m *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	// 解析类，注意这里解析出来的是数组类，而不是数组元素的类
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(m.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	// 数组有多少维度从操作数栈中弹出多少个整数表示每个维度长度
	counts := popAndCheckCounts(stack, int(m.dimensions))
	// 创建多维数组
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

// N维数组弹出N个正整数
func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

// N维数组创建
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
