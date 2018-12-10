package main

import (
	"fmt"
	"github.com/go-land/user-service/handlers"
	"github.com/go-land/user-service/proto"
	"github.com/micro/go-micro"
)

func main() {

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	user.RegisterUserServiceHandler(srv.Server(), handlers.NewUserServiceHandler())

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
