package main

import (
	"context"
	pb "grpc-rest/proto"
	"log"

	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Api struct {
	pb.UnimplementedMoviesApiServer
}

func (s *Api) GetAllMovies(ctx context.Context, req *pb.PageSearchRequest) (*pb.PageSearchResponse, error) {
	return &pb.PageSearchResponse{}, nil
}

func (s *Api) GetMoviesById(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	return &pb.IdResponse{}, nil
}

func main() {

	go func() {
		//mux
		mux := runtime.NewServeMux()

		//Register
		pb.RegisterMoviesApiHandlerServer(context.Background(), mux, nil)

		//http Server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))

	}()

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterMoviesApiServer(grpcServer, nil)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

//Belum selesai
