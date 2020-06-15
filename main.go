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
	voteServer     = "localhost:50507"
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

	r1, err := lsClient.GetLocations(ctx1, &request)
	if err != nil {
		log.Fatalf("Location service error: %v", err)
	}

	log.Printf("Locations: %v", r1.GetLocations())

	// Set up a connection to the story server.
	ssConn, err := grpc.Dial(storyServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Story server did not connect: %v", err)
	}
	defer ssConn.Close()

	ssClient := pb.NewStoryServiceClient(ssConn)

	var locationIds []string
	locationIds = append(locationIds, "loc-123")
	locationIds = append(locationIds, "loc-456")

	var storyRequest pb.StoryRequest
	storyRequest.LocationIds = locationIds

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	r2, err := ssClient.GetStories(ctx2, &storyRequest)
	if err != nil {
		log.Fatalf("Story service error: %v", err)
	}

	log.Printf("Stories: %v", r2.GetStories())

	// Set up a connection to the vote server.
	vsConn, err := grpc.Dial(voteServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Story server did not connect: %v", err)
	}
	defer ssConn.Close()

	vsClient := pb.NewVoteServiceClient(vsConn)

	contextID := "s-123"

	var voteRequest pb.VoteRequest
	voteRequest.ContextId = contextID

	ctx3, cancel3 := context.WithTimeout(context.Background(), time.Second)
	defer cancel3()

	rs, err := vsClient.GetVotes(ctx3, &voteRequest)
	if err != nil {
		log.Fatalf("Story service error: %v", err)
	}

	log.Printf("Stories: %v", rs.GetVotes())
}
