package classfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (e *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	e.classIndex = reader.readUint16()
	e.methodIndex = reader.readUint16()
}

func (e *EnclosingMethodAttribute) ClassName() string {
	return e.cp.getClassName(e.classIndex)
}

func (e *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if e.methodIndex > 0 {
		return e.cp.getNameAndType(e.methodIndex)
	} else {
		return "", ""
	}
}
