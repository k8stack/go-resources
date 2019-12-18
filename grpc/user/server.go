package userpb

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server!")

}

func initServer(){
	s := grpc.NewServer()

}