package heap

import "GJvm/classfile"

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
func (m *MethodRef) resolveMethodRef() {
}
