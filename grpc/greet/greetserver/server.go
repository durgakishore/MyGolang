package main

import (
	"Projects/GoSamples/grpc/greet/greetpb"
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"log"
	"net"
)

type server struct{}

//Unary
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("\nGreet Function is invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

//Server Streaming
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("\nGreetManyTimes Function is invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()

	for i := 1; i <= 10; i++ {
		result := "Hello " + firstName + " Response No: " + strconv.Itoa(i)
		res := &greetpb.GreetmanyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalln("GreetManyTimes send failed")
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

//Client Streaming

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	fmt.Printf("\nLongGreet Function is invoked as stream")

	result := ""

	for {
		res, err := stream.Recv()

		if err == io.EOF {

			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}
		result += "Hello " + res.GetGreeting().GetFirstName() + "!"
	}

}

//Bidirectional Streaming

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {

	fmt.Printf("\nGreetEveryone Function is invoked as stream")

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream in GreetEveryone %v", err)
		}

		result := "hello " + res.GetGreeting().GetFirstName()

		senderr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if senderr != nil {
			log.Fatalf("error while sending result to client %v", senderr)
		}
	}
}

//Unary with deadline
func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {

	fmt.Printf("\nGreetWithDeadline Function is invoked with %v", req)

	for i := 0; i < 3; i++ {

		if ctx.Err() == context.Canceled {
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil
}
func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listed the prot %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Printf("Server is started and waiting for client connections \n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
