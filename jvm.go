package main

import (
	"GJvm/classpath"
	"GJvm/instructions/base"
	"GJvm/rtda"
	"GJvm/rtda/heap"
	"fmt"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.isVerboseClass)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (j *JVM) start() {
	j.initVM()
	j.execMain()
}

func (j *JVM) initVM() {
	vmClass := j.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(j.mainThread, vmClass)
	interpret(j.mainThread, j.cmd.isVerboseInst)
}

func (j *JVM) execMain() {
	className := strings.Replace(j.cmd.class, ".", "/", -1)
	mainClass := j.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", j.cmd.class)
		return
	}

	argsArr := j.createArgsArray()
	frame := j.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	j.mainThread.PushFrame(frame)
	interpret(j.mainThread, j.cmd.isVerboseInst)
}

func (j *JVM) createArgsArray() *heap.Object {
	stringClass := j.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(j.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range j.cmd.args {
		jArgs[i] = heap.JString(j.classLoader, arg)
	}
	return argsArr
}
