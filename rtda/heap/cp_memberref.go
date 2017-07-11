package heap

import "GJvm/classfile"

// 类成员（字段/方法）符号引用
type MemberRef struct {
	SymRef
	name       string // 名称
	descriptor string // 描述符
}

func (m *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	m.className = refInfo.ClassName()
	m.name, m.descriptor = refInfo.NameAndDescriptor()
}

func (m *MemberRef) Name() string {
	return m.name
}
func (m *MemberRef) Descriptor() string {
	return m.descriptor
}
