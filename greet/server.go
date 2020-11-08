package main

import (
	 "context"
	 "fmt"
	 "io"
	 "log"
	 "net"
	 "strconv"
	 "time"

	 "github.com/grpc-go-course/greet/greetpb"

	 "google.golang.org/grpc"
	 "google.golang.org/grpc/codes"
	 "google.golang.org/grpc/status"
	 "google.golang.org/grpc/reflection"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{ Result: result, }
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " Number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{ Result: result, }
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}

func (*server) LongGreet (stream greetpb.GreetService_LongGreetServer) error {
	log.Println("Starting Long Greet service")
	var res string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Done receiving from LongGreetClient")
			break
		}
		if err != nil {
			log.Fatalf("received error while receiving from LongGreetClient: %v", err)
		}
		res = res + "Hello " + req.GetGreeting().GetFirstName() + " "
	}

	return stream.SendAndClose(&greetpb.LongGreetResponse{
		Result: res,
	})
}

func (*server) GreetWithDeadLine(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", req)

	for i := 1; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("Client cancelled the context")
			return nil, status.Error(codes.Canceled, "client canceled the request")
		}
		time.Sleep(time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{ Result: result, }
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// Enable reflection on Greet service
	reflection.Register(s)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
