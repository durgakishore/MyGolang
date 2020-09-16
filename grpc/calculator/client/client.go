package main

import (
	"Projects/GoSamples/grpc/calculator/calculatepb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v", err)
	}

	defer conn.Close()

	c := calculatepb.NewCalculateMaxServiceClient(conn)
	//findmaximum(c)
	squareRootOfNumber(c)
}

func squareRootOfNumber(c calculatepb.CalculateMaxServiceClient) {

	request := &calculatepb.SquareRootRequest{
		Num: -25,
	}
	res, err := c.SquareRoot(context.Background(), request)

	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Println(s.Message())
			fmt.Println(s.Code())
			if s.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("error while receving squareroot value from server %v", err)
			return
		}
	}

	fmt.Printf("Square of the number is %v\n", res.GetNumberRoot())
}

func findmaximum(c calculatepb.CalculateMaxServiceClient) {

	msg, _ := c.FindMaximum(context.Background())

	requests := []*calculatepb.FindMaxRequest{
		&calculatepb.FindMaxRequest{
			Num: 12,
		},
		&calculatepb.FindMaxRequest{
			Num: 14,
		},
		&calculatepb.FindMaxRequest{
			Num: 2,
		},
		&calculatepb.FindMaxRequest{
			Num: 22,
		},
		&calculatepb.FindMaxRequest{
			Num: 13,
		},
		&calculatepb.FindMaxRequest{
			Num: 33,
		},
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range requests {
			msg.Send(req)
		}
	}()

	go func() {
		for {
			res, err := msg.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			fmt.Println("number is ", res.GetMaximum())
		}
		close(wait)
	}()
	<-wait
}
