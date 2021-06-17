package main

import (
	"testing"
)

func TestGetFilmList(t *testing.T) {
	result, _ := GetFilmList("", "1")
	if result.Response != "False" {
		t.Fatalf("Should be not ok")
	}
	result, _ = GetFilmList("Batman", "0")
	if result.Response != "False" {
		t.Fatalf("Should be not ok")
	}
	result, _ = GetFilmList("Batman", "1")
	if result.Response != "True" {
		t.Logf("%v", result)
		t.Fatalf("Should be ok")
	}
}

func TestGetFilmDetail(t * testing.T) {
	result, _ := GetFilmDetail("tt0372784")
	if result.Response != "True" {
		t.Fatalf("should be ok")
	}
	result, _ = GetFilmDetail("random")
	if result.Response != "False" {
		t.Fatalf("should be not ok")
	}
}