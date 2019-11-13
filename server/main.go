package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/google/logger"
	//"log"
	"modules/log"
	"path/filepath"
	"modules/user"
	"modules/utility"
	//"net"
	"fmt"
	"os"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"net/http"
)
import "github.com/facebookgo/grace/gracenet"
import "github.com/facebookgo/grace/gracehttp"

var cfg *config.Config

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir:= filepath.Dir(ex)
	cfg = config.New("default")
	//load config
	cfg.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	cfg.AddDriver(json.Driver)

	err = cfg.LoadFiles(dir+"/../etc/config.json")
	if err != nil {
		panic(err)
	}

	log.Set(cfg.String("log.file"), cfg.Bool("log.verbose",true));
	utility.InitDb(cfg)    //{};
	utility.InitRedis(cfg) //{};
	//db := utility.Db{};
	//db.Config()
	//db.Init();
	val := cfg.Strings("db.default.master")
	fmt.Printf("\n master:\n %#v", val) // map[string]string{"key":"val2", "key2":"val20"}
	val2 := cfg.StringMap("db.default.slave.host")
	fmt.Printf("\n slave:\n %#v", val2) // map[string]string{"key":"val2", "key2":"val20"}

	logger.Info("config data: \n %#v\n", cfg.Data())
	listenRPC := cfg.String("listen.rpc", "")
	if listenRPC == "" {
		panic("rpc port is empty")
	}
	listenHTTP := cfg.String("listen.http", "")
	if listenHTTP == "" {
		panic("http port is empty")
	}
	proxyRPC := cfg.String("proxy.rpc", "")
	if proxyRPC == "" {
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
		user.RegisterHTTP(ctx, mux, proxyRPC, opts)

		gracehttp.Serve(
			&http.Server{Addr: listenHTTP, Handler: mux},
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
	lis, err := net.Listen("tcp", listenRPC)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

}
