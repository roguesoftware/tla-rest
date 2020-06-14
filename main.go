package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const (
	locationServer = "localhost:50505"
	storyServer    = "localhost:50506"
)

func main() {
	// Set up a connection to the location server.
	lsConn, err := grpc.Dial(locationServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Location server did not connect: %v", err)
	}
	defer lsConn.Close()

	lsClient := pb.NewLocationServiceClient(lsConn)

	var request pb.LocationRequest
	request.Longitude = -112.4
	request.Latitude = 32.4
	request.Radius = 100000

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	r, err := lsClient.GetLocations(ctx1, &request)
	if err != nil {
		log.Fatalf("Location service error: %v", err)
	}

	log.Printf("Locations: %v", r.GetLocations())

	// Set up a connection to the story server.
	ssConn, err := grpc.Dial(storyServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Story server did not connect: %v", err)
	}
	defer ssConn.Close()

	cs := pb.NewStoryServiceClient(ssConn)

	// locationIds := [2]string{"loc-123", "loc-456"}
	var locationIds []string
	locationIds = append(locationIds, "loc-123")
	locationIds = append(locationIds, "loc-456")

	var storyRequest pb.StoryRequest
	storyRequest.LocationIds = locationIds

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	rs, err := cs.GetStories(ctx2, &storyRequest)
	if err != nil {
		log.Fatalf("Story service error: %v", err)
	}

	log.Printf("Stories: %v", rs.GetStories())
}
