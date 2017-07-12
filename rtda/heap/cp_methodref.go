package heap

import "GJvm/classfile"

// 普通方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (m *MethodRef) ResolvedMethod() *Method {
	if m.method == nil {
		m.resolveMethodRef()
	}
	return m.method
}

// jvms8 5.4.3.3
// 普通方法解析
func (m *MethodRef) resolveMethodRef() {
	d := m.cp.class
	c := m.ResolvedClass()
	// 接口抛出异常
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 查找方法
	method := lookupMethod(c, m.name, m.descriptor)
	// 找不到
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 不可访问
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	m.method = method
}

// 根据方法名和描述符查找方法
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
