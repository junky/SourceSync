package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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

func RunClient(output chan string) *exec.Cmd {
	cmd := exec.Command("./sync_client")
	stdout, _ := cmd.StdoutPipe()

	buff := bufio.NewScanner(stdout)
	go func() {
		for buff.Scan() {
			output <- buff.Text()
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return cmd
}

func KillClient(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
	}
	cmd.Wait()
}

func ValidateOutput(output chan string, t *testing.T) {
	var s string

	s = <-output
	fmt.Println(s)
	if !strings.HasSuffix(s, "Starting SourceSync ...") {
		t.Fatal("Starting SourceSync ... not found")
	}

	s = <-output
	fmt.Println(s)
	if !strings.HasSuffix(s, "/tests") || !strings.Contains(s, "Sync Folder:") {
		t.Fatal("Starting SourceSync ... not found")
	}

}

func TestFSEvents(t *testing.T) {
	ClearTestsFolder()
	CreateTestsFolder()
	BuildClient()

	output := make(chan string)

	go ValidateOutput(output, t)

	cmd := RunClient(output)

	time.Sleep(100 * time.Millisecond)
	KillClient(cmd)
}
