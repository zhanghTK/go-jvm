package lang

import (
	"GJvm/native"
	"GJvm/rtda"
	"GJvm/rtda/heap"
	"fmt"
)

const jlThrowable = "java/lang/Throwable"

// 虚拟机栈帧信息
type StackTraceElement struct {
	fileName   string // 类所在文件名
	className  string // 方法的类名
	methodName string // 方法名
	lineNumber int    // 帧正在执行代码位置
}

func (s *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)", s.className, s.methodName, s.fileName, s.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	// this:异常对象
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 跳过所有异常类的构造（继承层次中的所有异常类都跳过）
	// 2：跳过fillInStackTrace()和fillInStackTrace(int)方法
	skip := distanceToObject(tObj.Class()) + 2
	// 获取有效虚拟机栈帧
	frames := thread.GetFrames()[skip:]
	// 根据虚拟机栈信息创建StackTraceElement
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 跳过异常类的构造方法
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
