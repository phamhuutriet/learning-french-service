package services

import (
	"learning-french-service/internal/repo"
	"learning-french-service/response"
)

// type UsersService struct {
// 	usersRepo *repo.UsersRepo
// }

// func NewUsersService() *UsersService {
// 	return &UsersService{
// 		usersRepo: repo.NewUsersRepo(),
// 	}
// }

// func (us *UsersService) GetUsers() []string {
// 	users := us.usersRepo.GetUsers()
// 	return users
// }

// INTERFACE VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	userExists := us.userRepo.GetUserByEmail(email)
	if userExists {
		return response.ErrCodeUserExists
	}

	return response.ErrCodeSuccess
}
