package main

import (
	"context"
	"fmt"
	"go-grpc/user/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) User(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	uname := req.User.Username
	return &userpb.UserResponse{
		UserResponse: "Welcome " + uname + "!",
	}, nil
}

func main() {
	fmt.Println("Starting server!")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Error at server", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Error at server", err)
	}
}
