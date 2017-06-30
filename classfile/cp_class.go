package classfile

/*
CONSTANT_Class_info {
    u1 tag;  // 7
    u2 name_index;  // 指向全限定名常量项的索引
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
