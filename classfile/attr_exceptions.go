package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;  // 方法可能抛出受检异常的种类数
    u2 exception_index_table[number_of_exceptions];  // 具体的受检异常，CONSTANT_Class_info型常量索引
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (e *ExceptionsAttribute) readInfo(reader *ClassReader) {
	e.exceptionIndexTable = reader.readUint16s()
}

func (e *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}
