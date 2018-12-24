package main

import (
	"github.com/go-land/job-service/proto"
	"github.com/go-land/user-service/dao"
	"github.com/go-land/user-service/handlers"
	"github.com/go-land/user-service/proto"
	"github.com/micro/go-micro"
	"gopkg.in/mgo.v2"
	"log"
)

const (
	mongoDB = "mongo"
)

func main() {

	// Create a new service. Optionally include some options here.
	server := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	server.Init()

	session, err := createMongoSession(mongoDB)
	defer session.Close()

	if err != nil {
		log.Fatalf("Can't properly connect to MongoDB %v\n", err)
	}

	// Register handler
	user.RegisterUserServiceHandler(server.Server(),
		handlers.NewUserServiceHandler(job.NewJobServiceClient("job",
			server.Client()), dao.NewUserDao(session)))

	// Run the server
	if err := server.Run(); err != nil {
		log.Fatalf("Can't properly start the server: %v\n", err)
	}
}

func createMongoSession(hostname string) (*mgo.Session, error) {

	session, err := mgo.Dial(hostname)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
