package classfile

// 字段/方法表
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16          // 访问标识
	nameIndex       uint16          // 常量名称索引
	descriptorIndex uint16          // 描述符
	//attributes      []AttributeInfo // 属性表
}

// 读取字段/方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段/方法表
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		//attributes:      readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}
