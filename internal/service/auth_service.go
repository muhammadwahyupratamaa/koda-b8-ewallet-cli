package service

import (
	"errors"
	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"

	"github.com/jackc/pgx/v5"
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
	user, err := s.userRepo.GetUserByUsername(username)

	if err == nil && user != nil {
	return errors.New("username already exists")
	}

	if err != nil && err != pgx.ErrNoRows {
	return err
	}
	user = &model.User{
		FullName: fullName,
		UserName: username,
		Age: age,
		Address: address,
		Email: email,
		Password: password,
	}

	return s.userRepo.CreateUser(user)
}