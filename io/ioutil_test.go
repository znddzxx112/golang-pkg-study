package iot

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReadAll(t *testing.T) {
	bs := []byte("heelllfooo")
	bsReader := bytes.NewReader(bs)

	rBs, ReadAll := ioutil.ReadAll(bsReader)
	if ReadAll != nil {
		t.Fatal(ReadAll)
	}
	t.Log(string(rBs))
}

func TestReadDir(t *testing.T) {
	fileinfos, ReadDirErr := ioutil.ReadDir(".")
	if ReadDirErr != nil {
		t.Fatal(ReadDirErr)
	}

	for _, fileinfo := range fileinfos {
		t.Log(fileinfo.Name())
		t.Log(fileinfo.IsDir())
		t.Log(fileinfo.Mode())
		fileContent, ReadFileErr := ioutil.ReadFile(fileinfo.Name())
		if ReadFileErr != nil {
			t.Fatal(ReadFileErr)
		}
		t.Log(string(fileContent))
	}
}
