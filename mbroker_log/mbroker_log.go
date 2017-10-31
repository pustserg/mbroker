package mbroker_log

import (
	"os"
	"log"
)

const logFileName  = "mbroker.log"

func InitLog() (*os.File, error) {
	logFile, err := openOrCreateLogFile()
	log.SetOutput(logFile)
	return logFile, err
}

func openOrCreateLogFile() (*os.File, error){
	_, checkError := os.Stat(logFileName)
	var file *os.File
	var err error
	if os.IsNotExist(checkError) {
		file, err = os.Create(logFileName)
	} else {
		file, err = os.OpenFile(logFileName, os.O_RDWR|os.O_APPEND, 0644)
	}
	return file, err
}
