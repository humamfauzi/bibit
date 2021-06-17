package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

const (
	BASE_URL = "http://example.com"
)

func TestWriteReply(t *testing.T) {
	w := httptest.NewRecorder()
	WriteReply(int(http.StatusOK), true, "{\"message\":\"ok\"}", w)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var replyMessage struct{
		Message string
	}
	var reply HandlerReply
	json.Unmarshal(body, &reply)
	json.Unmarshal([]byte(reply.Message.(string)), &replyMessage)
	if replyMessage.Message != "ok" {
		t.Fatalf("Should be ok")
	}
}

func TestGetFilms(t *testing.T) {
	queryString := "?page=1&search=Batman"
	req := httptest.NewRequest(http.MethodGet, BASE_URL + "/films" + queryString, nil) 
	w := httptest.NewRecorder()
	GetFilms(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var reply HandlerReply
	json.Unmarshal(body, &reply)
	if reply.Code != http.StatusOK && reply.Code != http.StatusInternalServerError {
		t.Fatalf("get statys %v instead of %v", reply.Code, http.StatusOK)
	}
}

func TestGetFilmDetails(t *testing.T) {
	pathVariableMap := make(map[string]string)
	pathVariableMap["id"] = "tt0372784"
	req := httptest.NewRequest(http.MethodGet, BASE_URL + "/films/{id}", nil) 
	req = mux.SetURLVars(req, pathVariableMap)
	w := httptest.NewRecorder()
	GetFilmDetails(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var reply HandlerReply
	json.Unmarshal(body, &reply)
	if reply.Code != http.StatusOK && reply.Code != http.StatusInternalServerError {
		t.Fatalf("get statys %v instead of %v", reply.Code, http.StatusOK)
	}
}

func TestGetQueryStringValue(t *testing.T) {
	qsMap := make(map[string][]string)
	qsMap["key1"] = []string{"value1"}
	qsMap["key2"] = []string{"value2", "value3"}
	result := GetQueryStringValue(qsMap, "key1", "")
	if result != "value1" {
		t.Fatalf("Should be equal to value 1")
	}

	result = GetQueryStringValue(qsMap, "key2", "")
	if result != "value2" {
		t.Fatalf("Should be equal to value 2")
	}

	result = GetQueryStringValue(qsMap, "not-exist", "replace")
	if result != "replace" {
		t.Fatalf("Should be equal to replace")
	}
}
