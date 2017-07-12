package heap

import "GJvm/classfile"

// 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (i *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if i.method == nil {
		i.resolveInterfaceMethodRef()
	}
	return i.method
}

// jvms8 5.4.3.4
// 接口方法解析
func (i *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := i.cp.class
	c := i.ResolvedClass()
	// 限制必须是接口
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 查找接口方法
	method := lookupInterfaceMethod(c, i.name, i.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	i.method = method
}

// 根据方法名和描述符查找方法
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
