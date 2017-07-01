package classfile

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;  // 常量池CONSTANT_Class_info型常量的索引
        u2 outer_class_info_index;  // 常量池CONSTANT_Class_info型常量的索引
        u2 inner_name_index;  // 常量池CONSTANT_Utf8_info型常量的索引，代表内部类名称
        u2 inner_class_access_flags;  // 内部类访问标识
    } classes[number_of_classes];
}
*/
type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (ica *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()
	ica.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range ica.classes {
		ica.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
