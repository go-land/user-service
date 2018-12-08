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
		log.Printf("firstName: %s, lastName: %s\n", singleUser.FirstName, singleUser.LastName)
	}

	singleUser, getByNameErr := client.GetByName(context.Background(), &user.GetByNameRequest{
		Name: "maksym",
	})

	if getByNameErr != nil {
		log.Fatalf("Can't properly call %s.GetByName(...) %v", serviceName, getByNameErr)
	}
	log.Printf("GetByName: %s\n", singleUser)
}
