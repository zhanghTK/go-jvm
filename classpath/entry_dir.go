package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
// 实现了entry接口
type DirEntry struct {
	absDir string // 目录的绝对路径
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir: absDir}
}

// 加载该目录下指定的类文件
func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

// 返回该目录的绝对路径
func (d *DirEntry) String() string {
	return d.absDir
}
