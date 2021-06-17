package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

type HandlerReply struct {
	Code int
	Success bool
	Message interface{}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logChan := GetLogChannel()
		fmt.Println("MIDDLEWARE")
		logChan <- r.Header.Get("User-Agent")
		next.ServeHTTP(w, r)
	})
}

func WriteReply(code int, success bool, message interface{}, w http.ResponseWriter) {
	hr := HandlerReply{
		Code: code,
		Success: success,
		Message: message,
	}
	reply, _ := json.Marshal(hr)
	w.WriteHeader(code)
	w.Write(reply)
}

func GetFilms(w http.ResponseWriter, r *http.Request) {
	queryString, _ := url.ParseQuery(r.URL.RawQuery)
	page := GetQueryStringValue(queryString, "page", "1")
	searchName := GetQueryStringValue(queryString, "search", "")
	fmt.Println(page, searchName)
	if searchName == "" {
		WriteReply(int(http.StatusBadRequest), false, nil, w)
		return
	}
	result, ok := GetFilmList(searchName, page)
	if !ok {
		WriteReply(int(http.StatusInternalServerError), false, nil, w)
		return
	}
	WriteReply(int(http.StatusOK), true, result, w)
	return
}

func GetFilmDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		WriteReply(int(http.StatusBadRequest), false, nil, w)
		return
	}
	result, ok := GetFilmDetail(id)
	if !ok {
		WriteReply(int(http.StatusInternalServerError), false, nil, w)
		return
	}
	WriteReply(int(http.StatusOK), true, result, w)
	return
}

func GetQueryStringValue(qs map[string][]string, key string, defaultValue string) string {
	var queryStringValue []string
	queryStringValue, ok := qs[key]; if !ok  {
		queryStringValue = []string{defaultValue}
	}
	return queryStringValue[0]
}