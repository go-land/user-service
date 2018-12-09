package handlers

import (
	"context"
	"github.com/go-land/user-service/proto"
)

type UserServiceImpl struct {
	users map[string]*user.User
}

func (service *UserServiceImpl) GetAll(ctx context.Context, request *user.GetAllRequest) (*user.UserResponse, error) {

	var usersData []*user.User

	for _, value := range service.users {
		usersData = append(usersData, value)
	}

	return &user.UserResponse{
		Users: usersData,
	}, nil
}

func (service *UserServiceImpl) GetByName(ctx context.Context, request *user.GetByNameRequest) (*user.User, error) {

	userName := request.Name

	if singleUser, ok := service.users[userName]; ok {
		return singleUser, nil
	}

	return nil, nil
}

func (service *UserServiceImpl) AddUser(ctx context.Context, req *user.User) (*user.GenericResponse, error) {

	if service.users[req.Alias] != nil {
		return &user.GenericResponse{
			Message: "User with alias " + req.Alias + " already exists",
		}, nil
	}

	service.users[req.Alias] = req

	return &user.GenericResponse{
		Message: "Successfully added",
	}, nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, req *user.User) (*user.GenericResponse, error) {

	if service.users[req.Alias] == nil {
		return &user.GenericResponse{
			Message: "Can't update user with alias " + req.Alias + ". Not found.",
		}, nil
	}

	singleUser := service.users[req.Alias]
	singleUser.Alias = req.Alias
	singleUser.FirstName = req.FirstName
	singleUser.LastName = req.LastName

	service.users[singleUser.Alias] = singleUser

	return &user.GenericResponse{
		Message: "Successfully updated",
	}, nil
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, req *user.User) (*user.GenericResponse, error) {

	if service.users[req.Alias] == nil {
		return &user.GenericResponse{
			Message: "Can't delete user with alias " + req.Alias + ". Not found.",
		}, nil
	}

	return &user.GenericResponse{
		Message: "User with alis " + req.Alias + " deleted.",
	}, nil
}

func NewUserServiceHandler() *UserServiceImpl {

	userHandler := UserServiceImpl{}

	userHandler.users = make(map[string]*user.User)

	userHandler.users["maksym"] = &user.User{
		Alias:     "maksym",
		FirstName: "Maksym",
		LastName:  "Stepanenko",
	}

	userHandler.users["olesia"] = &user.User{
		Alias:     "olesia",
		FirstName: "Olesia",
		LastName:  "Stepanenko",
	}

	return &userHandler
}
