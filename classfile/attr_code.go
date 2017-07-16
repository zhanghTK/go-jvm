package classfile

/*
Code_attribute {
    u2 attribute_name_index;  // 指向执行CONSTANT_Utf8_info型常量索引，常量固定值为"Code"
    u4 attribute_length;  // 属性值长度
    u2 max_stack;  // 操作数栈深度最大值
    u2 max_locals;  // 局部变量表所需的存储空间
    u4 code_length;  // 字节码长度
    u1 code[code_length];  // 字节码指令
    u2 exception_table_length;
    {   u2 start_pc;  // 起始行（方法体开始偏移）
        u2 end_pc;  // 结束行（方法体开始偏移）
        u2 handler_pc;  // 异常处理（方法体开始偏移）
        u2 catch_type;  // 异常类型，CONSTANT_Class_info型常量索引
    } exception_table[exception_table_length];  // 显示异常处理表
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	c.code = reader.readBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}

func (c *CodeAttribute) MaxStack() uint {
	return uint(c.maxStack)
}
func (c *CodeAttribute) MaxLocals() uint {
	return uint(c.maxLocals)
}
func (c *CodeAttribute) Code() []byte {
	return c.code
}
func (c *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return c.exceptionTable
}
func (c *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range c.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (e *ExceptionTableEntry) StartPc() uint16 {
	return e.startPc
}
func (e *ExceptionTableEntry) EndPc() uint16 {
	return e.endPc
}
func (e *ExceptionTableEntry) HandlerPc() uint16 {
	return e.handlerPc
}
func (e *ExceptionTableEntry) CatchType() uint16 {
	return e.catchType
}
