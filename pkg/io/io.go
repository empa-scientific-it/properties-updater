package io

import (
	"io/ioutil"
	"os"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func TempFile() (*os.File, error) {	
	file, err := ioutil.TempFile("/tmp/", "test.*.properties")
	return file, err
}
