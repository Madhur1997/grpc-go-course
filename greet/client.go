package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
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
