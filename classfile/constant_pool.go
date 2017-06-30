package classfile

import "fmt"

// 存储常量池内容（14中表类型结构）
// 常量池的读取操作

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	// 常量池入口存放容量计数值
	cpCount := int(reader.readUint16())

	cp := make([]ConstantInfo, cpCount)
	// 常量池索引从 1 开始
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// ConstantLongInfo,ConstantDoubleInfo占两个位置
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// 按索引查找常量
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))
}

// 按索引字段查找名称和描述
func (c ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 按索引查找类名
func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

// 按索引查找UTF-8字符串
func (c ConstantPool) getUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
