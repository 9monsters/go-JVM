package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (dirEntrySelf *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(dirEntrySelf.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, dirEntrySelf, err
}
func (dirEntrySelf *DirEntry) String() string {
	return dirEntrySelf.absDir
}
