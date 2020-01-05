package main

import (
	"context"
	"grpc-poc/api"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}
	defer cc.Close()

	c := api.NewCalculatorServiceClient(cc)

	arg := &api.SumRequest{X: 1, Y: 2}
	reply, err := c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	arg = &api.SumRequest{X: 10, Y: 5}
	reply, err = c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	arg = &api.SumRequest{X: 10, Y: 3}
	reply, err = c.Sum(context.Background(), arg)
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}
	log.Printf("%v + %v = %v", arg.GetX(), arg.GetY(), reply.GetResult())

	a := &api.FibonacciRequest{Number: int64(999999999999999999)}

	stream, err := c.Fibonacci(context.Background(), a)
	if err != nil {
		log.Fatalf("Error! : %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Printf("THE END : %v", err)
			break
		}
		if err != nil {
			log.Fatalf("Error : %v", err)
		}

		log.Printf("Fibo: %v -> %v", a.GetNumber(), res.GetNumber())
	}
}
