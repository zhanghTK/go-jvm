package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;  // 12
    u2 name_index;  // 指向该字段或方法名称常量项的索引
    u2 descriptor_index;  // 指向该字段或方法描述符常量项的索引
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}
