package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZIP或JAR包形式的类路径
type ZipEntry struct {
	absPath string // 包的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

// 加载该包文件下指定的类文件
func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer reader.Close()
	for _, f := range reader.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

// 返回该包文件的绝对路径
func (z *ZipEntry) String() string {
	return z.absPath
}
