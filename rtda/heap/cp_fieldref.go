package heap

import "GJvm/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolveFieldRef()
	}
	return f.field
}

// jvms 5.4.3.2
func (f *FieldRef) resolveFieldRef() {
	d := f.cp.class
	c := f.ResolvedClass()
	field := lookupField(c, f.name, f.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	f.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	// 类内寻找字段
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 遍历接口寻找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 父类中寻找
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
