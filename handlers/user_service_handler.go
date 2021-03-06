package handlers

import (
	"github.com/go-land/job-service/proto"
	"github.com/go-land/user-service/dao"
	"github.com/go-land/user-service/proto"
	"golang.org/x/net/context"
	"log"
)

type UserServiceImpl struct {
	users      map[string]*user.User
	jobService job.JobServiceClient
	userDao    dao.UserDao
}

func (service *UserServiceImpl) GetAll(ctx context.Context, request *user.GetAllRequest, resp *user.UserResponse) error {

	log.Println("GetAll called")

	service.userDao.GetAll()

	var usersData []*user.User

	for _, singleUser := range service.users {

		jobResponse, err := service.jobService.GetJob(ctx, &job.GetJobRequest{
			Name: singleUser.Alias,
		})

		if err != nil {
			log.Printf("Error calling job service: %v\n", err)
			singleUser.Job = "<unknown>"
		} else {
			singleUser.Job = jobResponse.Job
		}

		usersData = append(usersData, singleUser)
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

func NewUserServiceHandler(jobClient job.JobServiceClient, userDao dao.UserDao) *UserServiceImpl {

	userHandler := UserServiceImpl{}

	userHandler.jobService = jobClient

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

	userHandler.userDao = userDao

	return &userHandler
}
