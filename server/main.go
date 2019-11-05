package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"modules/user"
	"modules/utility"
	//"net"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"net/http"
)
import "github.com/facebookgo/grace/gracenet"
import "github.com/facebookgo/grace/gracehttp"

var cfg *config.Config

func main() {
	cfg = config.New("default")
	//load config
	cfg.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	cfg.AddDriver(json.Driver)

	err := cfg.LoadFiles("etc/config.json")
	if err != nil {
		panic(err)
	}

	utility.InitDb(cfg)       //{};
	utility.InitRedis(cfg) //{};
	//db := utility.Db{};
	//db.Config()
	//db.Init();
	val := cfg.Strings("db.default.master")
	fmt.Printf("\n master:\n %#v", val) // map[string]string{"key":"val2", "key2":"val20"}
	val2 := cfg.StringMap("db.default.slave.host")
	fmt.Printf("\n slave:\n %#v", val2) // map[string]string{"key":"val2", "key2":"val20"}

	log.Printf("config data: \n %#v\n", cfg.Data())
	listen_rpc := cfg.String("listen.rpc", "")
	if listen_rpc == "" {
		panic("rpc port is empty")
	}
	listen_http := cfg.String("listen.http", "")
	if listen_http == "" {
		panic("http port is empty")
	}
	proxy_rpc := cfg.String("proxy.rpc", "")
	if proxy_rpc == "" {
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
		user.RegisterHTTP(ctx, mux, proxy_rpc, opts)

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
