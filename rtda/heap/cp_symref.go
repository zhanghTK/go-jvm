package heap

// 运行时常量池的符号引用基本结构
type SymRef struct {
	cp        *ConstantPool // 源类的运行时常量池指针
	className string        // 目标类的完全限定名
	class     *Class        // 目标类
}

func (s *SymRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

// jvms8 5.4.3.1
func (s *SymRef) resolveClassRef() {
	d := s.cp.class                      // 源类
	c := d.loader.LoadClass(s.className) // 使用源类的类加载器根据目标类的完全限定名加载
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
}
