package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;  // 区分常量类型
    u2 class_index;  // 指向声明字段的类或者接口描述符 CONSTANT_Class_info 的索引项
    u2 name_and_type_index;  // 指向字段描述符 CONSTANT_NameAndType 的索引项
}
CONSTANT_Methodref_info {
    u1 tag;  // 区分常量类型
    u2 class_index;  // 指向声明方法的类描述符CONSTANT_Class_info的索引项
    u2 name_and_type_index;  // 指向字段描述符 CONSTANT_NameAndType 的索引项
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;  // 区分常量类型
    u2 class_index;  // 指向声明方法的接口描述符CONSTANT_Class_info的索引项
    u2 name_and_type_index;  // 指向字段描述符 CONSTANT_NameAndType 的索引项
}
*/
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantMemberrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}
func (c *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}
