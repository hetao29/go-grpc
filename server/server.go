package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"models/user"
	//"net"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"net/http"
)
import "github.com/facebookgo/grace/gracenet"
import "github.com/facebookgo/grace/gracehttp"

func main() {
	//load config
	config.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	config.AddDriver(json.Driver)

	err := config.LoadFiles("etc/config.json")
	if err != nil {
		panic(err)
	}
	log.Printf("config data: \n %#v\n", config.Data())
	listen_rpc := config.String("listen.rpc", "")
	if listen_rpc == "" {
		panic("rpc port is empty")
	}
	listen_http := config.String("listen.http", "")
	if listen_http == "" {
		panic("http port is empty")
	}
	proxy_rpc := config.String("proxy.rpc", "")
	if proxy_rpc== "" {
		panic("proxy_rpc is empty")
	}

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
		err := user.RegisterHttp(ctx, mux, proxy_rpc, opts)
		if err != nil {
			log.Printf("error:%v", err)
		}

		gracehttp.Serve(
			&http.Server{Addr: listen_http, Handler: mux},
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
	lis, err := net.Listen("tcp", listen_rpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
