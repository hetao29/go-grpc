package main

import (
	"context"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"

	"proto/user/info"
	"proto/user/profile"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := info.NewInfoClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &info.LoginRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

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
