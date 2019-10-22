package main

import (
"context"
"net/http"
	"google.golang.org/grpc"
	"log"
	"net"
	"models/user"
"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	port = ":50051"
)

/*
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
*/

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.Register(s)

	//http proxy
	go func(){
		ctx := context.Background()
	     ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		defer s.GracefulStop()
			<-ctx.Done()
	}()
	     mux := runtime.NewServeMux()
	     opts := []grpc.DialOption{grpc.WithInsecure()}
	err :=	     user.RegisterHttp(ctx, mux,  "127.0.0.1:50051", opts)
	     if err != nil {
		     log.Printf("error:%v",err);
	     }

     // Start HTTP server (and proxy calls to gRPC server endpoint)
     http.ListenAndServe(":8081", mux)
	}()


	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
