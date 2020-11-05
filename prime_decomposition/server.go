package main

import (
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/grpc-go-course/prime_decomposition/primedecompositionpb"
)

type server struct {
}

func (*server) GetPrimeFactors(req *primedecompositionpb.ReqNumber, stream primedecompositionpb.PrimeDecomposition_GetPrimeFactorsServer) error {
	origNum := req.GetNum()
	sqrt_num := uint64(math.Floor(math.Sqrt(float64(origNum))))
	primes := make(map[uint64]bool)
	var i uint64
	for i = 2; i <= sqrt_num; i++ {
		composite := false
		for prime, _ := range primes {
			if i % prime == 0 {
				composite = true
				break
			}
		}
		if !composite {
			primes[i] = true
		}
	}

	for num := uint64(2); num <= sqrt_num; num++ {
		if _, ok := primes[num]; !ok {
			continue
		}
		for {
			if uint64(origNum) % num == 0 {
				primeFactor := &primedecompositionpb.PrimeFactor{ Num: int32(num), }
				stream.Send(primeFactor)
				origNum = int32(uint64(origNum) / num)
			} else {
				break
			}
			time.Sleep(time.Second)
		}
	}

	if origNum != 1 {
		primeFactor := &primedecompositionpb.PrimeFactor{ Num: int32(origNum), }
		stream.Send(primeFactor)
	}
	return nil
}

func main() {
	log.Println("Starting Prime Decomposition service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}
	s := grpc.NewServer()
	primedecompositionpb.RegisterPrimeDecompositionServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: v", err)
	}
}
