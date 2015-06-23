package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-fsnotify/fsnotify"
)

func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("D:\\Junky\\Projects\\SmartLing\\SourceSync\\golang")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func ScanAllFiles(location string) (err error) {

	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		if fileInfo.IsDir() {
			fmt.Println(path)
		}
		return
	}

	err = filepath.Walk(location, scan)

	return
}

func main() {
	fmt.Println("Hello, Alex")
	ScanAllFiles("D:\\Junky\\Projects\\SmartLing\\SourceSync\\golang\\")
	ExampleNewWatcher()
}
