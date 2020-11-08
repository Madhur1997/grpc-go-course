package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	fmt.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)

	doUnaryWithDeadline(c, 5)
	doUnaryWithDeadline(c, 1)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Madhur",
			LastName: "Batra",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Madhur",
			LastName: "Batra",
		},
	}

	greetManyTimesClient, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		res, err := greetManyTimesClient.Recv()
		if err == io.EOF {
			// end of the stream
			break
		}
		if err != nil {
			log.Fatalf("received error while receiving from GreetManyTimes: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", res.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting to do Client streaming RPC...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Received error while invoking LongGreet RPC: %v", err)
	}

	reqs := []*greetpb.LongGreetRequest{
			&greetpb.LongGreetRequest{
				Greeting: &greetpb.Greeting{
					FirstName: "Madhur",
					LastName: "Batra",
				},
			},
			&greetpb.LongGreetRequest{
				Greeting: &greetpb.Greeting{
					FirstName: "Stephane",
					LastName: "Maarek",
				},
			},
			&greetpb.LongGreetRequest{
				Greeting: &greetpb.Greeting{
					FirstName: "Tarun",
					LastName: "Batra",
				},
			},
		}

	for _, req := range reqs {
		log.Printf("Sending %v\n", req.GetGreeting().GetFirstName())
		if err := stream.Send(req); err != nil {
			log.Fatalf("Received error while sending LongGreetRequest: %v", err)
		}
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving from LongGreet RPC: %v", err)
	}
	log.Printf("Repsonse from LongGreet RPC: %s\n", res.GetResult())
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, duration time.Duration) {
	fmt.Println("Starting to do a Unary RPC with deadline...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Madhur",
			LastName: "Batra",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * duration)
	defer cancel()

	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timeout was hit, deadline exceeded")
			} else {
				log.Println("unexpected error: %v", err)
			}
		} else {
			log.Fatalf("error while calling Greet RPC: %v", err)
		}
		return
	}
	log.Printf("Response from Greet: %v", res.Result)
}
