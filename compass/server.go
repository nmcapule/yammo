package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/nmcapule/quadtree"
	"google.golang.org/grpc"

	pb "github.com/nmcapule/yammo/protos/v1/compass"
)

var (
	width  = flag.Float64("width", 100, "Number of vertical ticks")
	height = flag.Float64("height", 100, "Number of horizontal ticks")
	port   = flag.Int("port", 8080, "Port to serve")
)

type server struct {
	pb.UnimplementedCompassServer
}

func main() {
	qt := quadtree.NewQuadtree(quadtree.Bounds{
		X: 0, Y: 0, Width: *width, Height: *height,
	})
	fmt.Println(qt)

	s := grpc.NewServer()
	pb.RegisterCompassServer(s, &server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Fatalln(s.Serve(lis))
}
