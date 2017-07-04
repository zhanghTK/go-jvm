package heap

// 符号引用
type SymRef struct {
	cp        *ConstantPool // 运行时常量池指针
	className string        // 类的完全限定名
	class     *Class        // 运行时类结构指针
}

func (s *SymRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

// jvms8 5.4.3.1
func (s *SymRef) resolveClassRef() {
	d := s.cp.class
	c := d.loader.LoadClass(s.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	s.class = c
}
