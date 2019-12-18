package main

import (
	"context"
	"fmt"
	"go-grpc/user/userpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Client!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error establishing connection!")
	}
	u := userpb.NewUserServiceClient(conn)
	req := &userpb.UserRequest{
		User: &userpb.User{
			Username: "zillani",
		},
	}

	res, err := u.User(context.Background(), req)
	if err != nil {
		log.Fatalln("Error!", err)
	}

	fmt.Println(res)
	defer conn.Close()
}
