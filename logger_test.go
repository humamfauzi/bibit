package main

import (
	"testing"
	"os"
	"io/ioutil"
	"strings"
)

func TestCreateOrOpenLogFile(t *testing.T) {
	CreateOrOpenLogFile()
	_, err := os.Stat(LOG_FILE_NAME)
	if err != nil {
		t.Fatalf("shoul not error")
	}
}

func TestCreateLogChannel(t *testing.T) {
	CreateLogChannel()
}

func TestAppendToLogFile(t *testing.T) {
	ok := AppendToLogFile("LOG 1\n")
	if !ok {
		t.Fatalf("should be ok")
	}
	body, _ := ioutil.ReadFile(LOG_FILE_NAME)
	if !strings.Contains(string(body), "LOG 1") {
		t.Fatalf("%s should contain LOG 1", string(body))
	}
}

func TestGetLogChannel(t *testing.T) {
	lc := GetLogChannel()
	lc <- "LOG 2\n"
	body, _ := ioutil.ReadFile(LOG_FILE_NAME)
	if !strings.Contains(string(body), "LOG 2") {
		t.Fatalf("%s should contain LOG 2", string(body))
	}
}