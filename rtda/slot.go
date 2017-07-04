package rtda

import "GJvm/rtda/heap"

// 局部变量表的变量槽
type Slot struct {
	num int32
	ref *heap.Object
}
