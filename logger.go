package main

import (
	"os"
)

const (
	LOG_FILE_NAME = "run.log"
)

var logChan chan string
var logFile *os.File

func WriteLog(logChannel chan string) {
	for {
		logString := <- logChannel
		AppendToLogFile(logString)
	}
}

func CreateLogChannel() {
	CreateOrOpenLogFile()
	logChan = make(chan string)
	go WriteLog(logChan)
}

func CreateOrOpenLogFile() bool {
	_, err := os.Stat(LOG_FILE_NAME)
	if os.IsNotExist(err) {	
		logFile, _ = os.Create(LOG_FILE_NAME)
		return true
	} else if err == nil {
		logFile, _ = os.OpenFile(LOG_FILE_NAME, os.O_APPEND|os.O_WRONLY, 0644)
		return true
	} else {
		return false
	}
}

func AppendToLogFile(logString string) bool {
	if _, err := logFile.WriteString(logString); err != nil {
		return false
	}
	return true
}

func DetachLogFile() {
	logFile.Close()
}

func GetLogChannel() chan string {
	return logChan
}