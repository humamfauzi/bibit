package main

import (
	pb "./proto"
	"context"
	"errors"
)

type FilmsServer struct {
	pb.UnimplementedFilmsServer
}

func (fs FilmsServer) GetFilmsProto(ctx context.Context, req *pb.ListFilmRequest) (*pb.ListFilmReply, error) {
	searchName := req.Search
	page := req.Page
	result, ok := GetFilmList(searchName, page)
	if !ok {
		return &pb.ListFilmReply{}, errors.New("fail to fetch films")
	}
	return searchResultToProto(result), nil
}

func (fs FilmsServer) GetFilmDetails(ctx context.Context, req *pb.FilmDetailRequest) (*pb.FilmDetailReply, error) {
	id := req.Id
	result, ok := GetFilmDetail(id)
	if !ok {
		return &pb.FilmDetailReply{}, errors.New("fail to fetch film detail")
	}
	return detailResultToProto(result), nil
}

func searchResultToProto(result OmdbApiResult) *pb.ListFilmReply {
	searchList := make([]*pb.SearchResult, len(result.Search))
	for i:=0; i < len(searchList); i++ {
		searchList[i] = &pb.SearchResult{
			Title: result.Search[i].Title,
			Year: result.Search[i].Year,
			ImdbId: result.Search[i].ImdbId,
			Type: result.Search[i].Type,
			Poster: result.Search[i].Poster,
		}
	}
	return &pb.ListFilmReply{
		Search: searchList,
		TotalResults: result.TotalResults,
		Response: result.Response,
	}
}

func detailResultToProto(result OmdbFilmDetails) *pb.FilmDetailReply {
	ratings := make([]*pb.Rating, len(result.Ratings))
	for i:=0; i < len(ratings); i++ {
		ratings[i] = &pb.Rating{
			Source: result.Ratings[i].Source,
			Value: result.Ratings[i].Value,
		}
	}
	return &pb.FilmDetailReply{
		Title: result.Title,
		Year: result.Year,
		Rated: result.Rated,
		Released: result.Released,
		Runtime: result.Runtime,
		Genre: result.Genre,
		Director: result.Director,
		Writer: result.Writer,
		Actors: result.Actors,
		Plot: result.Plot,
		Language: result.Language,
		Country: result.Country,
		Awards: result.Awards,
		Poster: result.Poster,
		Ratings: ratings,
		Metascore: result.Metascore,
		ImdbRating: result.ImdbRating,
		ImdbVotes: result.ImdbVotes,
		ImdbID: result.ImdbID,
		Type: result.Type,
		DVD: result.DVD,
		BoxOffice: result.BoxOffice,
		Production: result.Production,
		Website: result.Website,
		Response: result.Response,
	}
}