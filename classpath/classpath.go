package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPth  Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClassPth = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClassPath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClassPath = newEntry(cpOption)
}

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

func (c *Classpath) String() string {
	return c.userClassPath.String()
}

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

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
