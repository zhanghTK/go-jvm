package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// 类路径的抽象接口
type Entry interface {
	// 寻找、加载class文件
	// className：类文件的相对路径，路径之间使用/分割
	// []byte：类加载进内存的二进制形式
	// Entry：类所在的绝对路径
	readClass(className string) ([]byte, Entry, error)

	// 返回变量字符串表示，类似Java的toString方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
