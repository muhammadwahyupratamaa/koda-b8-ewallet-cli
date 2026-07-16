package service

import (
	"errors"
	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/utils"

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
	hashedPassword := utils.HashPassword(password)
	user = &model.User{
		FullName: fullName,
		UserName: username,
		Age: age,
		Address: address,
		Email: email,
		Password: hashedPassword,
	}

	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(
	username string,
	password string,
) error {

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return errors.New("username or password is incorrect")
	}

	hashedPassword := utils.HashPassword(password)

	if user.Password != hashedPassword {
		return errors.New("username or password is incorrect")
	}

	return nil
}