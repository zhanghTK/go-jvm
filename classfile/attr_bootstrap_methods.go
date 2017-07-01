package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;  // 对常量池CONSTANT_MethodHandle_info的索引
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];  // 常量池的索引
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (b *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	b.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range b.bootstrapMethods {
		b.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
