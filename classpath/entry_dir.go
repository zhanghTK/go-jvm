package classpath

import (
	"path/filepath"
	"io/ioutil"
)

// 目录形式的类路径
type DirEntry struct {
	// 目录的绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err) // 抛出异常
	}
	return &DirEntry{absDir}
}

func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
