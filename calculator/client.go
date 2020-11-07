package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	invokeAvg(c)
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
	log.Printf("Response from Add RPC: %v", res.GetRes())
}

func invokeAvg(c calculatorpb.CalculatorClient) {
	fmt.Println("Starting to invoke Avg RPC...")
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v", err)
	}
	
	reqNums := []*calculatorpb.ReqNum{
			&calculatorpb.ReqNum{ Num: 4, },
			&calculatorpb.ReqNum{ Num: 5, },
			&calculatorpb.ReqNum{ Num: 6, },
			&calculatorpb.ReqNum{ Num: 7, },
			&calculatorpb.ReqNum{ Num: 8, },
			&calculatorpb.ReqNum{ Num: 9, },
		}

	for _, reqNum := range reqNums {
		err := stream.Send(reqNum)
		if err != nil {
			log.Fatalf("error while sending req in the stream: %v", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while closing the stream: %v", err)
	}
	log.Printf("Calculated average is %v", res.GetRes())
}
