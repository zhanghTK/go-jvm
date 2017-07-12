package heap

// 方法描述符
type MethodDescriptor struct {
	parameterTypes []string // 参数类型
	returnType     string   // 返回值类型
}

func (m *MethodDescriptor) addParameterType(t string) {
	pLen := len(m.parameterTypes)
	if pLen == cap(m.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, m.parameterTypes)
		m.parameterTypes = s
	}

	m.parameterTypes = append(m.parameterTypes, t)
}
