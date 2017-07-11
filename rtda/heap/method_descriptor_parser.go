package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (m *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	m.raw = descriptor
	m.parsed = &MethodDescriptor{}
	m.startParams()
	m.parseParamTypes()
	m.endParams()
	m.parseReturnType()
	m.finish()
	return m.parsed
}

func (m *MethodDescriptorParser) startParams() {
	if m.readUint8() != '(' {
		m.causePanic()
	}
}
func (m *MethodDescriptorParser) endParams() {
	if m.readUint8() != ')' {
		m.causePanic()
	}
}
func (m *MethodDescriptorParser) finish() {
	if m.offset != len(m.raw) {
		m.causePanic()
	}
}

func (m *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + m.raw)
}

func (m *MethodDescriptorParser) readUint8() uint8 {
	b := m.raw[m.offset]
	m.offset++
	return b
}
func (m *MethodDescriptorParser) unreadUint8() {
	m.offset--
}

func (m *MethodDescriptorParser) parseParamTypes() {
	for {
		t := m.parseFieldType()
		if t != "" {
			m.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (m *MethodDescriptorParser) parseReturnType() {
	if m.readUint8() == 'V' {
		m.parsed.returnType = "V"
		return
	}

	m.unreadUint8()
	t := m.parseFieldType()
	if t != "" {
		m.parsed.returnType = t
		return
	}

	m.causePanic()
}

func (m *MethodDescriptorParser) parseFieldType() string {
	switch m.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return m.parseObjectType()
	case '[':
		return m.parseArrayType()
	default:
		m.unreadUint8()
		return ""
	}
}

func (m *MethodDescriptorParser) parseObjectType() string {
	unread := m.raw[m.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		m.causePanic()
		return ""
	} else {
		objStart := m.offset - 1
		objEnd := m.offset + semicolonIndex + 1
		m.offset = objEnd
		descriptor := m.raw[objStart:objEnd]
		return descriptor
	}
}

func (m *MethodDescriptorParser) parseArrayType() string {
	arrStart := m.offset - 1
	m.parseFieldType()
	arrEnd := m.offset
	descriptor := m.raw[arrStart:arrEnd]
	return descriptor
}
