package store

import (
	"GJvm/instructions/base"
	"GJvm/rtda"
)

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (f *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(f.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (f *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (f *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (f *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (f *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
