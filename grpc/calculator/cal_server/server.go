package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Sum(ctx context.Context, in *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	result := in.FirstNumber + in.SecondNumber
	return &calculatorpb.SumResponse{
		SumResult: result,
	}, nil
}

func main() {
	fmt.Println("Starting server!")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("error!", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(listener); err !=nil {
		log.Fatalln("Error", err)
	}
}
