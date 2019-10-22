package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"models/user"
	//"net"
	"net/http"
)
import "github.com/facebookgo/grace/gracenet"
import "github.com/facebookgo/grace/gracehttp"

const (
	port = ":50051"
)

func main() {
	//load config

	//init db

	//init redis

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		//register http proxy
		err := user.RegisterHttp(ctx, mux, "127.0.0.1:50051", opts)
		if err != nil {
			log.Printf("error:%v", err)
		}

		gracehttp.Serve(
			&http.Server{Addr: ":8082", Handler: mux},
		)
	}()

	//grpc server
	s := grpc.NewServer()

	//graceful stop
	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	//register grpc
	user.Register(s)
	net := gracenet.Net{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
