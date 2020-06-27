package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/nmcapule/yammo/protos/v1/compass"
)

var address = flag.String("address", "localhost:8080", "gRPC endpoint to connect to")

func main() {
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCompassClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.FindEntity(ctx, &pb.FindEntityRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", response)
}
