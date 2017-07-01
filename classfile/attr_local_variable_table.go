package classfile

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;  // 局部变量生命周期开始的字节码偏移量
        u2 length;  // 范围覆盖长度
        u2 name_index;  // 常量池CONSTANT_Utf8_info型常量的索引，局变量名称
        u2 descriptor_index;  // 常量池CONSTANT_Utf8_info型常量的索引，局部变量描述符
        u2 index;  // 局部变量在栈帧局表变量表的位置
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	l.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range l.localVariableTable {
		l.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
