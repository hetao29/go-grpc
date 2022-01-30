package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"syscall"
	//https://pkg.go.dev/log#pkg-constants
	"log"
	"modules/user"
	"modules/utility"
	"path/filepath"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"net"
	"net/http"
	"os"
	"os/signal"
)

import _ "google.golang.org/grpc/encoding/gzip"

var cfg *config.Config
var (
	version         string
	build           string
	service         = "test"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Println("version=", version)
	log.Println("build=", build)
	dir := filepath.Dir(ex)
	cfg = config.New("default")

	//load config
	cfg.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	cfg.AddDriver(json.Driver)

	err = cfg.LoadFiles(dir + "/../etc/config.json")
	if err != nil {
		panic(err)
	}

	//
	log.SetFlags(log.LstdFlags | log.Llongfile)

	//init db
	utility.InitDb(cfg) //{};

	//init redis
	utility.InitRedis(cfg) //{};
	val := cfg.Strings("db.default.master")
	log.Printf("master: %#v", val) // map[string]string{"key":"val2", "key2":"val20"}
	val2 := cfg.StringMap("db.default.slave.host")
	log.Printf("slave: %#v", val2) // map[string]string{"key":"val2", "key2":"val20"}

	log.Printf("config data: %#v", cfg.Data())
	listenRPC := cfg.String("listen.rpc", "")
	if listenRPC == "" {
		panic("rpc port is empty")
	}
	log.Printf("listenRPC: %s", listenRPC)
	listenHTTP := cfg.String("listen.http", "")
	if listenHTTP == "" {
		panic("http port is empty")
	}
	log.Printf("listenHTTP: %s", listenHTTP)
	proxyRPC := cfg.String("proxy.rpc", "")
	if proxyRPC == "" {
		panic("proxy_rpc is empty")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	//http proxy
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	//register http proxy
	user.RegisterHTTP(ctx, mux, proxyRPC, opts)
	//server.ListenAndServe(listenHTTP, mux)
	httpServer := &http.Server{Addr: listenHTTP, Handler: mux}
	go func() {

		if err := httpServer.ListenAndServe(); err != nil {
			// handle err
			log.Printf("failed to listen: %v", err)
		}
	}()

	//grpc server
	s := grpc.NewServer()
	go func() {

		//register grpc
		user.Register(s)
		lis, err := net.Listen("tcp", listenRPC)
		if err != nil {
			log.Printf("failed to listen: %v", err)
		} else {
			log.Printf("ok to listen: %v", err)
		}
		if err := s.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
		} else {
			log.Printf("ok to serve: %v")
		}
	}()

	<-stop
	log.Printf("got stop signal and ready to shutting down ...\n")
	log.Printf("grpc server shutting down ...\n")

	s.GracefulStop()
	log.Printf("grpc server shutting down ok\n")

	log.Printf("http server shutting down ...\n")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Print(err)
	}
	log.Printf("http server shutting down ok\n")
	log.Printf("exit ok\n")

}
