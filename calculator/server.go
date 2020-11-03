package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) Add(ctx context.Context, operands *calculatorpb.Operands) (*calculatorpb.Result, error) {
	log.Printf("Add function was invoked with %v", operands)
	first_number := operands.GetFirstNumber()
	second_number := operands.GetSecondNumber()

	res := &calculatorpb.Result{ Sum: first_number + second_number, }
	return res, nil
}

func main() {
	fmt.Println("Hello from Calculator service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
