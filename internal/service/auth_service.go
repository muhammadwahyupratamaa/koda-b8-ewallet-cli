package service

import (
	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(
	fullName string,
	username string,
	age int,
	address string,
	email string,
	password string,
) error {

	user := &model.User{
		FullName: fullName,
		UserName: username,
		Age: age,
		Address: address,
		Email: email,
		Password: password,
	}

	return s.userRepo.CreateUser(user)
}