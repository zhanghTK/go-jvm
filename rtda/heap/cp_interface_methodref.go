package heap

import "GJvm/classfile"

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
func (i *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := i.ResolveClass()
	// todo
}
