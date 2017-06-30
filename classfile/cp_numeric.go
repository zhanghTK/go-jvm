package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;  // 3
    u4 bytes;  // 按照高位在前存储int值
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}
func (c *ConstantIntegerInfo) Value() int32 {
	return c.val
}

/*
CONSTANT_Float_info {
    u1 tag;  // 4
    u4 bytes;  // 按照高位在前存储float
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = math.Float32frombits(bytes)
}
func (c *ConstantFloatInfo) Value() float32 {
	return c.val
}

/*
CONSTANT_Long_info {
    u1 tag;  // 5
    u4 high_bytes;  // 按照高位在前存储long，高位
    u4 low_bytes;  // 按照高位在前存储long，低位
}
*/
type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = int64(bytes)
}
func (c *ConstantLongInfo) Value() int64 {
	return c.val
}

/*
CONSTANT_Double_info {
    u1 tag;  // 6
    u4 high_bytes;  // 按照高位在前存储double，高位
    u4 low_bytes;  // 按照高位在前存储double，低位
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = math.Float64frombits(bytes)
}
func (c *ConstantDoubleInfo) Value() float64 {
	return c.val
}
