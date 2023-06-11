package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"time"

	"github.com/thimc/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	if in.GetGreeting() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Greeting is empty",
		)
	}
	log.Printf("Received %+v\n", in)
	resp := &proto.HelloResponse{
		Reply: "world",
	}
	return resp, nil
}

func (s *server) SubscribeHello(src proto.HelloService_SubscribeHelloServer) error {
	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				log.Printf("The client closed the stream")
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("Client says", rr.GetGreeting())
		}
	}()

	for {
		reply := &proto.HelloResponse{
			Reply: "Hello",
		}
		if err := src.Send(reply); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	listenAddr := flag.String("listenAddr", ":50051", "the port on which the server will listen on")
	flag.Parse()

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
