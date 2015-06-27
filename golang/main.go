package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dmotylev/goproperties"
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

func ScanAllFiles(location string) (res []string, err error) {
	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		if inpErr != nil {
			fmt.Println(inpErr)
		}
		if fileInfo.IsDir() {
			res = append(res, path)
		}
		return
	}

	err = filepath.Walk(location, scan)
	return
}

func GetSyncFolder() (folder string, err error) {
	current_folder, err := os.Getwd()
	if nil != err {
		return
	}

	p, _ := properties.Load(current_folder + "/sync.settings")
	folder = p.String("path", "")

	return
}

func main() {

	sync_folder, _ := GetSyncFolder()
	folders, _ := ScanAllFiles(sync_folder)

	for _, folder := range folders {
		fmt.Println(folder)
	}

	fmt.Println("Sync Folder:" + sync_folder)
}
