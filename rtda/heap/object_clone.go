package heap

func (o *Object) Clone() *Object {
	return &Object{
		class: o.class,
		data:  o.cloneData(),
	}
}

func (o *Object) cloneData() interface{} {
	switch o.data.(type) {
	case []int8:
		elements := o.data.([]int8)
		elements2 := make([]int8, len(elements))
		copy(elements2, elements)
		return elements2
	case []int16:
		elements := o.data.([]int16)
		elements2 := make([]int16, len(elements))
		copy(elements2, elements)
		return elements2
	case []uint16:
		elements := o.data.([]uint16)
		elements2 := make([]uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case []int32:
		elements := o.data.([]int32)
		elements2 := make([]int32, len(elements))
		copy(elements2, elements)
		return elements2
	case []int64:
		elements := o.data.([]int64)
		elements2 := make([]int64, len(elements))
		copy(elements2, elements)
		return elements2
	case []float32:
		elements := o.data.([]float32)
		elements2 := make([]float32, len(elements))
		copy(elements2, elements)
		return elements2
	case []float64:
		elements := o.data.([]float64)
		elements2 := make([]float64, len(elements))
		copy(elements2, elements)
		return elements2
	case []*Object:
		elements := o.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := o.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
