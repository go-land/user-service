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

func NewUserServiceHandler() *UserServiceImpl {

	userHandler := UserServiceImpl{}

	userHandler.users = make(map[string]*user.User)

	userHandler.users["maksym"] = &user.User{
		FirstName: "Maksym",
		LastName:  "Stepanenko",
	}

	userHandler.users["olesia"] = &user.User{
		FirstName: "Olesia",
		LastName:  "Stepanenko",
	}

	return &userHandler
}
