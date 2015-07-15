package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func ClearTestsFolder() {
	os.RemoveAll("tests")
}

func CreateTestsFolder() {
	os.MkdirAll("tests/Folder-1", 0777)
	ioutil.WriteFile("tests/Folder-1/test1.txt", []byte("text"), 0666)
}

func TestFSEvents(t *testing.T) {
	ClearTestsFolder()
	CreateTestsFolder()
}
