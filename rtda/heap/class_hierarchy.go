package heap

// class继承结构
// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (cl *Class) isAssignableFrom(other *Class) bool {
	s, t := other, cl

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
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

// iface extends cl
func (cl *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(cl)
}
