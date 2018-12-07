package main

import (
	"context"
	"github.com/go-land/user-service/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	serviceName = "user-service"
	address     = "localhost:8080"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't not connect: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	resp, err := client.GetAll(context.Background(), &user.GetAllRequest{})

	if err != nil {
		log.Fatalf("Can't properly call %s %v", serviceName, err)
	}

	for _, singleUser := range resp.Users {
		log.Printf("UfirstName: %s, lastName: %s\n", singleUser.FirstName, singleUser.LastName)
	}
}
