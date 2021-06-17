package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"time"
	"net"
	"google.golang.org/grpc"
	pb "./proto"
)
const (
	GPRC_PORT = "8010"
	HTTP_PORT = "8000"
)

func main() {
	CreateLogChannel()
	router := setMux()
	
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + HTTP_PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server working")
	setGrpc()
	log.Fatal(server.ListenAndServe())
}

func setMux() *mux.Router {
	router := mux.NewRouter()
	router.Use(LoggerMiddleware)
	router.HandleFunc("/films", GetFilms).Methods(http.MethodGet)
	router.HandleFunc("/films/{id}", GetFilmDetails).Methods(http.MethodGet)
	return router
}

func setGrpc() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", GPRC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFilmsServer(grpcServer, FilmsServer{})
	grpcServer.Serve(lis)
}