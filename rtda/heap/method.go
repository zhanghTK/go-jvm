package heap

import "GJvm/classfile"

type Method struct {
	ClassMember
	maxStack        uint           // 操作数栈大小
	maxLocals       uint           // 局部变量表大小
	code            []byte         // 字节码
	argSlotCount    uint           // 参数个数
	exceptionTable  ExceptionTable // 异常处理表
	lineNumberTable *classfile.LineNumberTableAttribute
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
		m.lineNumberTable = codeAttr.LineNumberTableAttribute()
		m.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), m.class.constantPool)
	}
}

func (m *Method) calcArgSlotCount(paramTypes []string) {
	// 计算参数个数
	for _, paramType := range paramTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	// 非静态方法隐含this
	if !m.IsStatic() {
		m.argSlotCount++ // `this` reference
	}
}

// 用于给静态方法注入字节码
func (m *Method) injectCodeAttribute(returnType string) {
	m.maxStack = 4 // todo
	m.maxLocals = m.argSlotCount
	switch returnType[0] {
	case 'V':
		m.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		m.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		m.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		m.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		m.code = []byte{0xfe, 0xad} // lreturn
	default:
		m.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags&ACC_SYNCHRONIZED
}
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags&ACC_BRIDGE
}
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags&ACC_VARARGS
}
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags&ACC_NATIVE
}
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags&ACC_ABSTRACT
}
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags&ACC_STRICT
}

// getters
func (m *Method) MaxStack() uint {
	return m.maxStack
}
func (m *Method) MaxLocals() uint {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}
func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}

func (m *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := m.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (m *Method) GetLineNumber(pc int) int {
	if m.IsNative() {
		return -2
	}
	if m.lineNumberTable == nil {
		return -1
	}
	return m.lineNumberTable.GetLineNumber(pc)
}
