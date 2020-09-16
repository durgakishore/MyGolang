package main

import (
	"Projects/GoSamples/grpc/calculator/calculatepb"
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) FindMaximum(stream calculatepb.CalculateMaxService_FindMaximumServer) error {

	var prevNumber int32 = 0
	for {
		req, _ := stream.Recv()
		number := req.GetNum()
		if number > prevNumber {
			stream.Send(&calculatepb.FindMaxResponse{
				Maximum: number,
			})
			prevNumber = number
		}
	}

}

func (*server) SquareRoot(ctx context.Context, req *calculatepb.SquareRootRequest) (*calculatepb.SquareRootResponse, error) {

	number := (float64)(req.GetNum())
	fmt.Printf("Received number is %v\n", number)

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received numver is negative number %v", number),
		)
	}

	return &calculatepb.SquareRootResponse{
		NumberRoot: math.Sqrt(number),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("error while listening to the localhost %v", err)
	}

	s := grpc.NewServer()

	calculatepb.RegisterCalculateMaxServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to connect the server %v", err)
	}
}
