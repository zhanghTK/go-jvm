package classfile

import "fmt"

const MAGIC_NUMBER uint32 = 0xCAFEBABE // 魔数值

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	// magic uint32  魔数
	minorVersion uint16        // 次版本
	majorVersion uint16        // 主版本
	constantPool ConstantPool  // 常量池
	accessFlags  uint16        // 访问标识
	thisClass    uint16        // 类名
	superClass   uint16        // 超类名
	interfaces   []uint16      // 接口索引表
	fields       []*MemberInfo // 字段表
	methods      []*MemberInfo // 方法表
	//attributes   []AttributeInfo
}

// 将已加载的类解析为ClassFile
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 匿名函数捕获异常
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{data: classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	c.constantPool = readConstantPool(reader)
	c.accessFlags = reader.readUint16()
	c.thisClass = reader.readUint16()
	c.superClass = reader.readUint16()
	c.interfaces = reader.readUint16s()
	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
	//c.attributes = readAttributes(reader, c.constantPool)
}

func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != MAGIC_NUMBER {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUint16()
	c.majorVersion = reader.readUint16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}
func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

func (c *ClassFile) ClassName() string {
	return c.constantPool.getClassName(c.thisClass)
}

func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return ""
}

func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
