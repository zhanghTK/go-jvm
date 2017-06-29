package classfile

/*
CONSTANT_Class_info {
    u1 tag;  // 区分常量类型
    u2 name_index;  // 索引值，指向常量池中一个CONSTANT_Utf8_info的常量
}
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}
func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
