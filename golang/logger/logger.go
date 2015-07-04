package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type logger struct {
	file   *os.File
	logger *log.Logger
}

var log_var = logger{file: nil, logger: nil}

func checkAndCreateDirectories(filename string) {
	directories := filepath.Dir(filename)
	os.MkdirAll(directories, 0777)
}

func SetOutputFile(filename string) {
	checkAndCreateDirectories(filename)

	var f_err error = nil
	log_var.file, f_err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if f_err != nil {
		log.Fatal("error opening file: ", f_err)
	}

	multi := io.MultiWriter(log_var.file, os.Stdout)

	log_var.logger = log.New(multi, "", log.LstdFlags)
}

func Close() {
	log_var.file.Close()
	log_var.file = nil
	log_var.logger = nil
}

func Fatal(v ...interface{}) {
	if log_var.logger != nil {
		log_var.logger.Fatal(v...)
	} else {
		log.Fatal(v...)
	}
}

func Println(v ...interface{}) {
	if log_var.logger != nil {
		log_var.logger.Println(v...)
	} else {
		log.Println(v...)
	}
}
