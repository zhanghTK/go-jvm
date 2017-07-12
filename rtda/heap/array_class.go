package heap

func (cl *Class) IsArray() bool {
	return cl.name[0] == '['
}

func (cl *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(cl.name)
	return cl.loader.LoadClass(componentClassName)
}

func (cl *Class) NewArray(count uint) *Object {
	if !cl.IsArray() {
		panic("Not array class: " + cl.name)
	}
	switch cl.Name() {
	case "[Z":
		return &Object{cl, make([]int8, count)}
	case "[B":
		return &Object{cl, make([]int8, count)}
	case "[C":
		return &Object{cl, make([]uint16, count)}
	case "[S":
		return &Object{cl, make([]int16, count)}
	case "[I":
		return &Object{cl, make([]int32, count)}
	case "[J":
		return &Object{cl, make([]int64, count)}
	case "[F":
		return &Object{cl, make([]float32, count)}
	case "[D":
		return &Object{cl, make([]float64, count)}
	default:
		return &Object{cl, make([]*Object, count)}
	}
}
