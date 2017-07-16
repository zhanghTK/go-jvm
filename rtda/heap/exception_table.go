package heap

import "GJvm/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (e ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range e {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		// 判断范围
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				// 捕获所有类型异常
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			// 判断异常类型
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
