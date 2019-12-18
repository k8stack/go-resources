package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Request from client!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error connecting", err)
	}
	req, err := &calculatorpb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 2,
	}, nil
	s := calculatorpb.NewSumServiceClient(conn)
	resp, err := s.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Error", err)
	}
	fmt.Println(resp)
}
