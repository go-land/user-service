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

	addUser(client)

	updateUser(client)

	getAll(client)

	getByName(client)

	deleteUser(client)

}

func addUser(client user.UserServiceClient) {

	log.Println("AddUser <-- call")

	resp, err := client.AddUser(context.Background(), &user.User{
		Alias:     "zorro",
		FirstName: "Zorro",
		LastName:  "El Capitano",
	})

	if err != nil {
		log.Fatalf("Can't properly call %s.AddUser %v", serviceName, err)
	}

	log.Printf("Response message: %s\n", resp.Message)

	log.Println("AddUser <-- ok")
}

func updateUser(client user.UserServiceClient) {
	log.Println("UpdateUser <-- call")

	resp, err := client.UpdateUser(context.Background(), &user.User{
		Alias:     "zorro",
		FirstName: "Zorro-123",
		LastName:  "El Capitano-123",
	})

	if err != nil {
		log.Fatalf("Can't properly call %s.UpdateUser %v", serviceName, err)
	}

	log.Printf("Updated: %s\n", resp.Message)

	log.Println("UpdateUser <-- ok")
}

func getAll(client user.UserServiceClient) {

	log.Println("GetAll <-- call")

	resp, err := client.GetAll(context.Background(), &user.GetAllRequest{})

	if err != nil {
		log.Fatalf("Can't properly call %s %v", serviceName, err)
	}

	for _, singleUser := range resp.Users {
		log.Printf("User %s\n", singleUser)
	}

	log.Println("GetAll <-- ok")
}

func getByName(client user.UserServiceClient) {

	log.Println("GetByName <-- call")

	singleUser, getByNameErr := client.GetByName(context.Background(), &user.GetByNameRequest{
		Name: "maksym",
	})

	if getByNameErr != nil {
		log.Fatalf("Can't properly call %s.GetByName(...) %v", serviceName, getByNameErr)
	}

	log.Printf("user: %s \n", singleUser)

	log.Println("GetByName <-- ok")
}

func deleteUser(client user.UserServiceClient) {
	log.Println("DeleteUser <-- call")

	resp, err := client.DeleteUser(context.Background(), &user.User{Alias: "zorro"})

	if err != nil {
		log.Fatalf("Can't properly call %s.GetByName(...) %v", serviceName, err)
	}

	log.Printf("Deleted: " + resp.Message)

	log.Println("DeleteUser <-- ok")
}
