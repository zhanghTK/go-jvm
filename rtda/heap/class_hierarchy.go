package heap

// class继承结构
// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (cl *Class) isAssignableFrom(other *Class) bool {
	s, t := other, cl

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// cl extends cl
func (cl *Class) IsSubClassOf(other *Class) bool {
	for c := cl.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// cl implements iface
func (cl *Class) IsImplements(iface *Class) bool {
	for c := cl; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// cl extends iface
func (cl *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range cl.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends cl
func (cl *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(cl)
}
