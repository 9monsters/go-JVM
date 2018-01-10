package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath, nil}
}

func (zipEntrySelf *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if zipEntrySelf.zipRC == nil {
		err := zipEntrySelf.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := zipEntrySelf.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, zipEntrySelf, err
}

// todo: close zip
func (zipEntrySelf *ZipEntry) openJar() error {
	r, err := zip.OpenReader(zipEntrySelf.absPath)
	if err == nil {
		zipEntrySelf.zipRC = r
	}
	return err
}

func (zipEntrySelf *ZipEntry) findClass(className string) *zip.File {
	for _, f := range zipEntrySelf.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (zipEntrySelf *ZipEntry) String() string {
	return zipEntrySelf.absPath
}
