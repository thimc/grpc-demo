package main

import (
	"context"
	"flag"
	"log"

	"github.com/thimc/grpc-demo/proto"

	"google.golang.org/grpc"
)

var ctx = context.Background()

func main() {
	connectionString := flag.String("connectionString", "localhost:50051", "the remote address of the server")
	flag.Parse()

	conn, err := grpc.Dial(*connectionString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := proto.NewHelloServiceClient(conn)

	log.Println("RPC Call")
	Call(serviceClient)

	log.Println("Subscribing")
	Subscribe(serviceClient)
}

func Call(serviceClient proto.HelloServiceClient) {
	request := &proto.HelloRequest{
		Greeting: "hello",
	}
	resp, err := serviceClient.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Request: %s, Response: %s\n", request.GetGreeting(), resp.GetReply())
}

func Subscribe(serviceClient proto.HelloServiceClient) {
	client, err := serviceClient.SubscribeHello(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		received := proto.HelloResponse{}
		if err := client.RecvMsg(&received); err != nil {
			log.Fatal(err)
		}
		log.Printf("Server says %s\n", received.GetReply())

		request := proto.HelloRequest{
			Greeting: "world",
		}
		if err := client.SendMsg(&request); err != nil {
			log.Fatal(err)
		}
	}
}
