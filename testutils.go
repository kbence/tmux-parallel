package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type tempDir struct {
	Path string
}

func newTempDir() *tempDir {
	var err error
	dir := &tempDir{}

	if dir.Path, err = ioutil.TempDir("", "test_XXXXXX"); err != nil {
		panic(err)
	}

	return dir
}

func (d *tempDir) Release() {
	os.RemoveAll(d.Path)
}

func (d *tempDir) PathOf(filename string) string {
	return fmt.Sprintf("%s/%s", d.Path, filename)
}

func (d *tempDir) CreateFile(filename, content string) {
	if f, err := os.Create(d.PathOf(filename)); err == nil {
		defer f.Close()
		f.Write([]byte(content))
	} else {
		panic(err)
	}
}
