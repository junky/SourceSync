package main

import (
	"os"
	"os/signal"
	"path/filepath"

	"github.com/dmotylev/goproperties"
	"github.com/go-fsnotify/fsnotify"

	"./logger"
)

func ProcessEvents(watcher *fsnotify.Watcher) {
	for {
		select {
		case event := <-watcher.Events:
			logger.Println("event:", event)
			if event.Op&fsnotify.Create == fsnotify.Create {
				path := event.Name
				if info, err := os.Stat(path); err == nil && info.IsDir() {
					logger.Println("Adding new folder to watcher:", path)
					err = watcher.Add(path)
					if err != nil {
						logger.Fatal(err)
					}
				}
			}
		case err := <-watcher.Errors:
			logger.Println("error:", err)
		}
	}
}

func ScanAllFiles(location string) (res []string, err error) {
	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		if inpErr != nil {
			logger.Println(inpErr)
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
	logger.SetOutputFile("logs/sync.log")

	logger.Println("Starting SourceSync ...")

	sync_folder, _ := GetSyncFolder()
	logger.Println("Sync Folder:" + sync_folder)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal(err)
	}
	defer watcher.Close()
	defer logger.Close()

	go ProcessEvents(watcher)

	folders, _ := ScanAllFiles(sync_folder)
	for _, folder := range folders {
		logger.Println("Watching folder: " + folder)
		err = watcher.Add(folder)
		if err != nil {
			logger.Fatal(err)
		}
	}

	// kill event, ctrl+c
	onkill := make(chan os.Signal, 1)
	signal.Notify(onkill, os.Interrupt, os.Kill)
	<-onkill

	logger.Println("Stoping SourceSync ...")
}
