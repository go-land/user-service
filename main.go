package main

import (
	"github.com/go-land/user-service/handlers"
	"github.com/go-land/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	serviceName = "user-service"
	port        = ":8080"
)

func main() {

	// Set-up our gRPC server.
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()

	// Register our service with the gRPC server.
	user.RegisterUserServiceServer(server, createUserServiceHandler())

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Printf("%s started successfully at port %s\n", serviceName, port)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func createUserServiceHandler() *handlers.UserServiceImpl {

	userHandler := handlers.UserServiceImpl{}

	userHandler.Users = make(map[string]*user.User)

	userHandler.Users["maksym"] = &user.User{
		FirstName: "Maksym",
		LastName:  "Stepanenko",
	}

	userHandler.Users["olesia"] = &user.User{
		FirstName: "Olesia",
		LastName:  "Stepanenko",
	}

	return &userHandler
}
