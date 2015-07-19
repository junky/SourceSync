package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func ClearTestsFolder() {
	os.RemoveAll("tests")
}

func CreateTestsFolder() {
	os.MkdirAll("tests/Folder-1", 0777)
	ioutil.WriteFile("tests/Folder-1/test1.txt", []byte("text"), 0666)
}

func BuildClient() {
	cmd := exec.Command("go", "build", "sync_client.go")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestFSEvents(t *testing.T) {
	ClearTestsFolder()
	CreateTestsFolder()
	BuildClient()

	cmd := exec.Command("./sync_client")
	stdout, _ := cmd.StdoutPipe()

	buff := bufio.NewScanner(stdout)
	go func() {
		for buff.Scan() {
			fmt.Printf("Stdout: %s\n", buff.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Kill")
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
	}

	cmd.Wait()

}
