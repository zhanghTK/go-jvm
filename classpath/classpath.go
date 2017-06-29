package classpath

import (
	"os"
	"path/filepath"
)

// CLASSPATH
// 包含启动类加载器，扩展类加载器，用户类加载器加载的类路径
type Classpath struct {
	bootClassPth  Entry // 启动类加载器加载的类路径
	extClassPath  Entry // 扩展类加载器加载的类路径
	userClassPath Entry // 用户类加载器加载的类路径
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

// 解析启动类加载器和用户类加载器对应类路径
func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	// JRE指定目录
	jreDir := getJreDir(jreOption)
	// {JRE}/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClassPth = newWildcardEntry(jreLibPath)
	// {JRE}/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClassPath = newWildcardEntry(jreExtPath)
}

// 解析用户类加载器对应类路径
func (c *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClassPath = newEntry(cpOption)
}

// 加载CLASSPATH下指定类
// 依次加载启动类加载器，扩展类加载器，用户类加载器下的类
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClassPth.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return c.userClassPath.readClass(className)
}

// 返回用户类加载器对应的类路径
func (c *Classpath) String() string {
	return c.userClassPath.String()
}

// 获取JRE目录
// 依此判断以下路径有效性：输入类路径、当前目录下jre目录、JAVA_HOME/jre目录下
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can't find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
