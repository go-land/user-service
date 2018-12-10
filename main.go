package main

import (
	"fmt"
	"github.com/go-land/job-service/proto"
	"github.com/go-land/user-service/handlers"
	"github.com/go-land/user-service/proto"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	log.Println("Service started 123")

	// Create a new service. Optionally include some options here.
	server := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	server.Init()

	// Register handler
	user.RegisterUserServiceHandler(server.Server(),
		handlers.NewUserServiceHandler(job.NewJobServiceClient("job", server.Client())))

	// Run the server
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
