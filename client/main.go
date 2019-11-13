package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"proto/user/info"
	"proto/user/profile"
)

const (
	defaultName = "world"
)

func main() {
	dir, _ := os.Getwd()
	//load config
	config.WithOptions(config.ParseEnv)

	// add Decoder and Encoder
	config.AddDriver(json.Driver)

	err := cfg.LoadFiles(dir+"/../etc/config.json")
	if err != nil {
		panic(err)
	}
	log.Printf("config data: \n %#v\n", config.Data())
	proxyRPC := config.String("proxy.rpc", "")
	if proxyRPC == "" {
		panic("proxy_rpc is empty")
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(proxyRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := info.NewInfoClient(conn)

	// Contact the server and print out its response.
	name := "admin"
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

	d := profile.NewProfileClient(conn)
	s, err := d.Update(ctx, &profile.UpdateRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", s.GetMessage())

	t, err := d.Get(ctx, &profile.GetRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", t.GetMessage())
}
