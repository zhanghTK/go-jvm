package heap

// 数组特有的方法
// 根据数组类型返回具体数据
// 没有专门实现Booleans()方法，与Bytes()方法实现相同
func (o *Object) Bytes() []int8 {
	return o.data.([]int8)
}
func (o *Object) Shorts() []int16 {
	return o.data.([]int16)
}
func (o *Object) Ints() []int32 {
	return o.data.([]int32)
}
func (o *Object) Longs() []int64 {
	return o.data.([]int64)
}
func (o *Object) Chars() []uint16 {
	return o.data.([]uint16)
}
func (o *Object) Floats() []float32 {
	return o.data.([]float32)
}
func (o *Object) Doubles() []float64 {
	return o.data.([]float64)
}
func (o *Object) Refs() []*Object {
	return o.data.([]*Object)
}

// 获取数组长度
func (o *Object) ArrayLength() int32 {
	switch o.data.(type) {
	case []int8:
		return int32(len(o.data.([]int8)))
	case []int16:
		return int32(len(o.data.([]int16)))
	case []int32:
		return int32(len(o.data.([]int32)))
	case []int64:
		return int32(len(o.data.([]int64)))
	case []uint16:
		return int32(len(o.data.([]uint16)))
	case []float32:
		return int32(len(o.data.([]float32)))
	case []float64:
		return int32(len(o.data.([]float64)))
	case []*Object:
		return int32(len(o.data.([]*Object)))
	default:
		panic("Not array!")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dst.data.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
