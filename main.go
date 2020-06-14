package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const (
	address = "localhost:50505"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewLocationServiceClient(conn)

	// Contact the server and print out its response.
	var request pb.LocationRequest
	request.Longitude = -112.4
	request.Latitude = 32.4
	request.Radius = 100000

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetLocations(ctx, &request)
	if err != nil {
		log.Fatalf("Location service error: %v", err)
	}

	log.Printf("Locations: %v", r.GetLocations())
}
