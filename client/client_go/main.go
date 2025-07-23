package main

import (
	"context"
	"log"
	"os"
	"time"
	"fmt"
	"path/filepath"
	"google.golang.org/grpc"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"github.com/hetao29/go-grpc/proto/user/info"
)

const (
	defaultName = "world"
)

var cfg *config.Config
var (
    version string
    build   string
)
func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("version=", version)
	fmt.Println("build=", build)
	dir:= filepath.Dir(ex)
	//load config
	cfg = config.New("default")
	cfg.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	cfg.AddDriver(json.Driver)

	err = cfg.LoadFiles(dir+"/../etc/config.json")
	if err != nil {
		panic(err)
	}
	log.Printf("config data: \n %#v\n", cfg.Data())
	listRPC:= cfg.String("listen.rpc", "")
	if  listRPC== "" {
		panic("proxy_rpc is empty")
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(listRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := info.NewInfoClient(conn)

	// Contact the server and print out its response.
	name := "a"
	pwd := "123456"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &info.LoginRequest{Name: name, Password: pwd})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.GetInfo())
}
