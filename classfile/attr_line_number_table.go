package classfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;  // 字节码行号
        u2 line_number;  // Java源码行号
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	l.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (l *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(l.lineNumberTable) - 1; i >= 0; i-- {
		entry := l.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
