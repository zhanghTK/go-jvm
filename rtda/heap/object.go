package heap

// 临时用来表示对象
type Object struct {
	class  *Class // 类信息指针
	fields Slots  // 实例变量表
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (o *Object) Class() *Class {
	return o.class
}
func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
