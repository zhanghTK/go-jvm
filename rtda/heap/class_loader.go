package heap

import (
	"GJvm/classfile"
	"GJvm/classpath"
	"fmt"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp        *classpath.Classpath // 类路径文件加载器指针
	isVerbose bool
	classMap  map[string]*Class // 已加载类数据（方法区）
}

func NewClassLoader(cp *classpath.Classpath, isVerbose bool) *ClassLoader {
	return &ClassLoader{
		cp:        cp,
		isVerbose: isVerbose,
		classMap:  make(map[string]*Class),
	}
}

func (c *ClassLoader) LoadClass(name string) *Class {
	if class, ok := c.classMap[name]; ok {
		return class
	}
	if name[0] == '[' {
		return c.loadArrayClass(name)
	}
	return c.loadNonArrayClass(name)
}

func (c *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //todo
		name:        name,
		loader:      c,
		initStarted: true,
		superClass:  c.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			c.LoadClass("java/lang/Cloneable"),
			c.LoadClass("java/lang/Serializable"),
		},
	}
	c.classMap[name] = class
	return class
}

// 加载普通
// 无法加载数组
func (c *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := c.readClass(name)
	class := c.defineClass(data)
	link(class)
	if c.isVerbose {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (c *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := c.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (c *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = c
	resolveSuperClass(class)
	resolveInterfaces(class)
	c.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
// 类变量分配空间并赋初值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 统计实例字段个数，为实例字段编号
// 从继承链最顶端开始编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 统计类字段个数，为实例字段编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 类变量分配空间并初始化
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 类变量赋初始值
// 描述符的补充：https://hacpai.com/article/1365927493304?m=0
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
