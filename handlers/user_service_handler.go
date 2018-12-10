package handlers

import (
	"github.com/go-land/user-service/proto"
	"golang.org/x/net/context"
	"log"
)

type UserServiceImpl struct {
	users map[string]*user.User
}

func (service *UserServiceImpl) GetAll(ctx context.Context, request *user.GetAllRequest, resp *user.UserResponse) error {

	log.Println("GetAll called")

	var usersData []*user.User

	for _, value := range service.users {
		usersData = append(usersData, value)
	}

	resp.Users = usersData

	return nil
}

func (service *UserServiceImpl) GetByName(ctx context.Context, request *user.GetByNameRequest, resp *user.User) error {

	log.Println("GetByName called")

	userName := request.Name

	if singleUser, ok := service.users[userName]; ok {
		resp = singleUser
	}

	return nil
}

func (service *UserServiceImpl) AddUser(ctx context.Context, req *user.User, resp *user.GenericResponse) error {

	log.Println("AddUser called")

	if service.users[req.Alias] != nil {

		resp = &user.GenericResponse{
			Message: "User with alias " + req.Alias + " already exists",
		}

		return nil
	}

	service.users[req.Alias] = req

	resp = &user.GenericResponse{
		Message: "Successfully added",
	}

	return nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, req *user.User, resp *user.GenericResponse) error {

	log.Println("UpdateUser called")

	if service.users[req.Alias] == nil {

		resp = &user.GenericResponse{
			Message: "Can't update user with alias " + req.Alias + ". Not found.",
		}

		return nil
	}

	singleUser := service.users[req.Alias]
	singleUser.Alias = req.Alias
	singleUser.FirstName = req.FirstName
	singleUser.LastName = req.LastName

	service.users[singleUser.Alias] = singleUser

	resp = &user.GenericResponse{
		Message: "Successfully updated",
	}

	return nil
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, req *user.User, resp *user.GenericResponse) error {

	log.Println("DeleteUser called")

	if service.users[req.Alias] == nil {
		resp = &user.GenericResponse{
			Message: "Can't delete user with alias " + req.Alias + ". Not found.",
		}
		return nil
	}

	resp = &user.GenericResponse{
		Message: "User with alis " + req.Alias + " deleted.",
	}
	return nil
}

func NewUserServiceHandler() *UserServiceImpl {

	userHandler := UserServiceImpl{}

	userHandler.users = make(map[string]*user.User)

	userHandler.users["maksym"] = &user.User{
		Alias:     "maksym",
		FirstName: "Maksym",
		LastName:  "Stepanenko",
		Job:       "<undefined>",
	}

	userHandler.users["olesia"] = &user.User{
		Alias:     "olesia",
		FirstName: "Olesia",
		LastName:  "Stepanenko",
		Job:       "<undefined>",
	}

	return &userHandler
}
