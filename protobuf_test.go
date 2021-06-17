package main

import (
	"testing"
	"context"
	pb "./proto"
)

func TestFilmServer(t *testing.T) {
	fs := FilmsServer{}
	ctx := context.Background()

	getFilmRequest := &pb.ListFilmRequest{
		Search: "Batman",
		Page: "1",
	}
	resultGet, _ := fs.GetFilmsProto(ctx, getFilmRequest)
	if resultGet.Response != "True" {
		t.Fatalf("expect to have true")
	}

	filmId := &pb.FilmDetailRequest{
		Id: "tt0372784",
	}
	resultDetail, _ := fs.GetFilmDetails(ctx, filmId)
	if resultDetail.Response != "True" {
		t.Fatalf("expect to have true")
	}
}