package main

import (
	"encoding/json"
	"net/url"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	OMDB_HOST = "http://www.omdbapi.com"
	API_KEY = "API_KEY"
)

type OmdbApiResult struct {
	Search []SearchResult `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response string `json:"Response"`
}

type SearchResult struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	ImdbId string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}

type OmdbFilmDetails struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string`json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Ratings []Rating `json:"Ratings"`
	Metascore string `json:"Metascore"`
	ImdbRating string `json:"ImdbRating"`
	ImdbVotes string `json:"ImdbVotes"`
	ImdbID string `json:"ImdbID"`
	Type string `json:"Type"`
	DVD string `json:"DVD"`
	BoxOffice string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website string `json:"Website"`
	Response string `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value string `json:"Value"`
}

func GetFilmList(searchName string, page string) (OmdbApiResult, bool) {
	base, _ := url.Parse(OMDB_HOST)
	params := url.Values{}
	apiKey := GetApiKey()

	params.Add("s", searchName)
	params.Add("page", page)
	params.Add("apikey", apiKey)

	base.RawQuery = params.Encode()
	var finalResult OmdbApiResult
	body, ok := RequestAPI(base.String())
	if !ok {
		return finalResult, false
	}
	err := json.Unmarshal(body, &finalResult)
	if err != nil {
		return finalResult, false
	}
	return finalResult, true
}

func GetFilmDetail(imdbId string) (OmdbFilmDetails, bool) {
	base, _ := url.Parse(OMDB_HOST)
	params := url.Values{}
	apiKey := GetApiKey()
	params.Add("i", imdbId)
	params.Add("plot", "full")
	params.Add("apikey", apiKey)

	base.RawQuery = params.Encode()
	var finalResult OmdbFilmDetails
	body, ok := RequestAPI(base.String())
	if !ok {
		return finalResult, false
	}

	err := json.Unmarshal(body, &finalResult)
	if err != nil {
		return finalResult, false
	}
	return finalResult, true
}

func RequestAPI(urlString string) ([]byte, bool) {
	resp, err := http.Get(urlString)
	if err != nil {
		return []byte{}, false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, false
	}
	return body, true
}

func GetApiKey() string {
	return os.Getenv(API_KEY)
}