package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 结尾包含通配符（*）的类路径
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	// 遍历通配符路径的上级目录
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过子目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 处理Jar/jar文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	})
	return compositeEntry
}
