package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	
	"github.com/grpc-go-course/calculator/calculatorpb"
)

func main() {
	fmt.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorClient(cc)
	invokeAdd(c)
}

func invokeAdd(c calculatorpb.CalculatorClient) {
	fmt.Println("Starting to invoke Add RPC...")
	operands := &calculatorpb.Operands{
		FirstNumber: 3,
		SecondNumber: 10,
	}
	res, err := c.Add(context.Background(), operands)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Add RPC: %v", res.Sum)
}
