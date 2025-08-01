package services

import "learning-french-service/internal/repo"

type UsersService struct {
	usersRepo *repo.UsersRepo
}

func NewUsersService() *UsersService {
	return &UsersService{
		usersRepo: repo.NewUsersRepo(),
	}
}

func (us *UsersService) GetUsers() []string {
	users := us.usersRepo.GetUsers()
	return users
}
