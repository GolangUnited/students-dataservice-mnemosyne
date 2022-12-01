package main

import (
	"context"
	"flag"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewMnemosyneClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "userid", "1")
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if code := status.Code(err); code != codes.OK {
		log.Fatalf("could not greet: %v", err)
	}
	/*if err != nil {
		log.Fatalf("could not greet: %v", err)
	}*/
	log.Printf("Greeting: %s", r.GetMessage())
}
