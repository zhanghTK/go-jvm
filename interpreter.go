package main

import (
	"GJvm/instructions"
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
	"fmt"
)

// 解释器
func interpret(method *heap.Method, logInst bool) {
	// 创建线程
	thread := rtda.NewThread()
	// 创建栈帧
	frame := thread.NewFrame(method)
	// 插入栈帧
	thread.PushFrame(frame)
	// 异常处理
	defer catchErr(thread)
	// 循环处理虚拟机栈内容
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		// 操作码
		opcode := reader.ReadUint8()
		// 根据操作码获取指令
		inst := instructions.NewInstruction(opcode)
		// 读取操作数
		inst.FetchOperands(reader)
		// 更新PC位置
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// 执行指令
		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
