package heap

import (
	"strings"
	"GJvm/classfile"
)

// 运行时类信息
type Class struct {
	accessFlags       uint16        // 类的访问标志
	name              string        // 类名
	superClassName    string        // 超类名
	interfaceNames    []string      // 接口名
	constantPool      *ConstantPool // 运行时常量池指针
	fields            []*Field      // 字段表
	methods           []*Method     // 方法表
	loader            *ClassLoader  // 类加载器指针
	superClass        *Class        // 超类指针
	interfaces        []*Class      // 接口指针
	instanceSlotCount uint          // 实例变量数量
	staticSlotCount   uint          // 静态变量数量
	staticVars        Slots         // 静态变量
}

// class文件信息转换为class结构体信息
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (cl *Class) IsPublic() bool {
	return 0 != cl.accessFlags&ACC_PUBLIC
}
func (cl *Class) IsFinal() bool {
	return 0 != cl.accessFlags&ACC_FINAL
}
func (cl *Class) IsSuper() bool {
	return 0 != cl.accessFlags&ACC_SUPER
}
func (cl *Class) IsInterface() bool {
	return 0 != cl.accessFlags&ACC_INTERFACE
}
func (cl *Class) IsAbstract() bool {
	return 0 != cl.accessFlags&ACC_ABSTRACT
}
func (cl *Class) IsSynthetic() bool {
	return 0 != cl.accessFlags&ACC_SYNTHETIC
}
func (cl *Class) IsAnnotation() bool {
	return 0 != cl.accessFlags&ACC_ANNOTATION
}
func (cl *Class) IsEnum() bool {
	return 0 != cl.accessFlags&ACC_ENUM
}

// getters
func (cl *Class) ConstantPool() *ConstantPool {
	return cl.constantPool
}
func (cl *Class) StaticVars() Slots {
	return cl.staticVars
}

// jvms 5.4.4
// 类的访问条件：
// 1. public修饰
// 2. 同在一个包下
func (cl *Class) isAccessibleTo(other *Class) bool {
	return cl.IsPublic() || cl.getPackageName() == other.getPackageName()
}

func (cl *Class) getPackageName() string {
	if i := strings.LastIndex(cl.name, "/"); i >= 0 {
		return cl.name[:i]
	}
	return ""
}

func (cl *Class) GetMainMethod() *Method {
	return cl.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (cl *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range cl.methods {
		if isStaticMethod(method, name, descriptor) {
			return method
		}
	}
	return nil
}

func isStaticMethod(method *Method, name string, descriptor string) bool {
	return method.IsStatic() && method.name == name && method.descriptor == descriptor
}

func (cl *Class) NewObject() *Object {
	return newObject(cl)
}
