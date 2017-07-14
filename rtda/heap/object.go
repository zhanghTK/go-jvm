package heap

// 临时用来表示对象
type Object struct {
	class *Class      // 类信息指针
	// 以interface{}形式容纳各种类型的元素
	data  interface{} // 实例变量表
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (o *Object) Class() *Class {
	return o.class
}
func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

// reflection
func (o *Object) GetRefVar(name, descriptor string) *Object {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (o *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
