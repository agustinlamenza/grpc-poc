package main

import (
	"grpc-poc/api"
	"time"

	"context"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer cc.Close()

	c := api.NewCalculatorServiceClient(cc)

	arg := &api.SumRequest{X: 1, Y: 2}
	reply, err := c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	arg = &api.SumRequest{X: 10, Y: 5}
	reply, err = c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	arg = &api.SumRequest{X: 10, Y: 3}
	reply, err = c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	a := &api.FibonacciRequest{Number: int64(99999999999)}

	stream, err := c.Fibonacci(context.Background(), a)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Printf("THE END: %v", err)
			break
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Fibo: %v -> %v", a.GetNumber(), res.GetNumber())
	}

	numbers := []int64{1, 3, 5, 7, 9, 123, 456, 789, 321, 654, 978, 123, 3456, 567, 46578, 3456, 24356, 2345, 2345, 2345, 2345}
	strm, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, n := range numbers {
		time.Sleep(200 * time.Millisecond)
		strm.Send(&api.AvrRequest{Number: n})
		log.Printf("Number: %v", n)
	}

	res, err := strm.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("El average is: %v", res.GetAvr())
}
