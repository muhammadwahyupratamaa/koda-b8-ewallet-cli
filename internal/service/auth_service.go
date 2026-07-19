package service

import (
	"errors"
	"fmt"
	"time"

	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/utils"

	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	sessionRepo *repository.SessionRepository
	walletRepo  *repository.WalletRepository
}

func NewAuthService(
	userRepo *repository.UserRepository,
	sessionRepo *repository.SessionRepository,
	walletRepo *repository.WalletRepository,
) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		walletRepo:  walletRepo,
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
		Age:      age,
		Address:  address,
		Email:    email,
		Password: hashedPassword,
	}
	
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	walletNumber := fmt.Sprintf("WAL%06d", user.ID)

	err = s.walletRepo.CreateWallet(user.ID, walletNumber)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(
	username string,
	password string,
) (*model.User, error) {

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	hashedPassword := utils.HashPassword(password)

	if user.Password != hashedPassword {
		return nil, errors.New("username or password is incorrect")
	}

	err = s.sessionRepo.DeleteSessionByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	session := &model.UserSession{
		UserID:       user.ID,
		SessionToken: utils.GenerateSessionToken(),
		ExpiredAt:    time.Now().Add(24 * time.Hour),
	}

	err = s.sessionRepo.CreateSession(session)
	if err != nil {
		return nil, err
	}

	return user, nil
}