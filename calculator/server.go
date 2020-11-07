package main

import (
	"context"
	"fmt"
	"io"
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

	res := &calculatorpb.Result{ Res: first_number + second_number, }
	return res, nil
}

func (s *server) Avg(stream calculatorpb.Calculator_AvgServer) error {
	log.Printf("Avg function was invoked")

	sum := int32(0)
	count := int32(0)
	for {
		reqNum, err := stream.Recv()
		if err == io.EOF {
			avg := sum/count
			return stream.SendAndClose(&calculatorpb.Result{ Res: avg, })
		}
		if err != nil {
			log.Fatalf("received error while trying to receive from the stream: %v", err)
		}
		addResult, err := s.Add(context.Background(), &calculatorpb.Operands{FirstNumber: sum, SecondNumber: reqNum.GetNum(), })
		sum = addResult.GetRes()
		count++
	}
	return nil
}

func (s *server) Max(stream calculatorpb.Calculator_MaxServer) error {
	log.Printf("Max function was invoked")

	prevMax := int32(0)
	for {
		reqNum, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("received error while trying to receive from the stream: %v", err)
		}

		if reqNum.GetNum() > prevMax {
			prevMax = reqNum.GetNum()
			err := stream.Send(&calculatorpb.Result{Res: prevMax, })
			if err != nil {
				log.Fatalf("error while sending to the stream: v", err)
			}
		}
	}
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
