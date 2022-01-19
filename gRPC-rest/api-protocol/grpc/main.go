package main

import (
	"context"
	"fmt"
	"grpc-rest/api-protocol/rest/request"
	pb "grpc-rest/proto"
	"log"
	"reflect"

	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GetApi struct {
	pb.UnimplementedMoviesApiServer
	request request.RequestMovies
}

func (s *GetApi) GetAllMovies(ctx context.Context, req *pb.PageSearchRequest) (*pb.Info, error) {
	fmt.Println(reflect.TypeOf(req.Page))
	value := s.request.GetAllMovies(string(req.Page), string(req.Search))
	fmt.Println(value)
	return &pb.Info{
		Movies: &pb.Movies{
			Title:  "value.Search",
			Year:   "2001",
			ImdbID: "1001",
			Poster: "contoh",
			Type:   "contoh",
		},
		TotalResults: value.TotalResults,
		Response:     value.Response,
	}, nil
}

func (s *GetApi) GetMoviesById(ctx context.Context, req *pb.IdRequest) (*pb.Movies, error) {
	fmt.Println(req)
	return &pb.Movies{
		Title:  "contoh",
		Year:   "2001",
		ImdbID: "1001",
		Poster: "contoh",
		Type:   "contoh",
	}, nil
}

func main() {

	go func() {
		//mux
		mux := runtime.NewServeMux()

		//Register
		pb.RegisterMoviesApiHandlerServer(context.Background(), mux, &GetApi{})
		//pb.RegisterMoviesApiHandlerServer(context.Background(), mux, &Api{})
		fmt.Println("Server running on port : 8000")
		//http Server
		log.Fatalln(http.ListenAndServe("localhost:8000", mux))

	}()

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterMoviesApiServer(grpcServer, &GetApi{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

//Belum selesai
