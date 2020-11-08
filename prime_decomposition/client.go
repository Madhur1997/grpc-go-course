package main

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"

	"grpc-go-course/prime_decomposition/primedecompositionpb"

	"google.golang.org/grpc"
)

func main() {
	log.Println("prime-decomposition client")

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("error while converting string to integer: %v", err)
	}
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	c := primedecompositionpb.NewPrimeDecompositionClient(cc)

	defer cc.Close()
	reqNum := &primedecompositionpb.ReqNumber{ Num: int32(num), }
	stream, err := c.GetPrimeFactors(context.Background(), reqNum)
	if err != nil {
		log.Fatalf("error while calling GetPrimeFactors RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// end of stream
			break
		} else if err != nil {
			log.Fatalf("received error while receiving from GetPrimeFactors: %v", err)
		}
		log.Printf("Prime factor: %d", res.GetNum())
	}
}
