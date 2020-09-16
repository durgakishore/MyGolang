package main

import (
	"Projects/GoSamples/grpc/greet/greetpb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	fmt.Println("hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Client created %f", c)

	doUnary(c)
	doServerStreaming(c)

	doClientStreaming(c)
	doBiDirectionalStreaming(c)
	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 1*time.Second)
}

//Unary
func doUnary(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kishore",
			LastName:  "Yerubandi",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("ERROR WHILE CALLING  Greet RPC %v", err)
	}

	fmt.Printf("Response from Greet %v", res.Result)
}

//Server streaming

func doServerStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a server streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rishika",
			LastName:  "Yerubandi",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling  GreetManyTimes RPC %v", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading  stream from  GreetManyTimes RPC %v", err)
		}

		log.Printf("Response from Greet %v", msg.GetResult())
	}
}

//Client streaming
func doClientStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a client  streaming RPC")
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling  LongGreet RPC %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rishika",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mrudula",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lakshmi",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Karthik",
			},
		},
	}

	for _, req := range requests {
		fmt.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receving response from   LongGreet RPC %v", err)
	}

	log.Printf("Response from LongGreet %v", res)
}

//Bidirectional streaming
func doBiDirectionalStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Bidirectional streaming RPC")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling  GreetEveryone RPC %v", err)
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore1",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore2",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore3",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kishore4",
			},
		},
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending  request %v\n", req)
			err := stream.Send(req)
			time.Sleep(100 * time.Millisecond)
			if err != nil {
				log.Fatalf("error while sending request to server %v \n", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receving response from server %v", err)
			}

			fmt.Printf("Received response %v\n", res.GetResult())
		}
		close(wait)
	}()

	<-wait
}

//UnaryWithDeadline
func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {

	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kishore",
			LastName:  "Yerubandi",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout is hit.Deadline exceeded")
			} else {
				log.Fatalf("ERROR WHILE CALLING  GreetWithDealine RPC %v", err)
			}
		}
		return
	}
	fmt.Printf("Response from GreetWithDealine %v\n", res.Result)
}
