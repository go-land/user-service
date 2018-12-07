package handlers

import (
	"context"
	"github.com/go-land/user-service/proto"
)

type UserServiceImpl struct {
	Users map[string]*user.User
}

func (service *UserServiceImpl) GetAll(ctx context.Context, request *user.GetAllRequest) (*user.UserResponse, error) {

	var usersData []*user.User

	for _, value := range service.Users {
		usersData = append(usersData, value)
	}

	return &user.UserResponse{
		Users: usersData,
	}, nil
}
