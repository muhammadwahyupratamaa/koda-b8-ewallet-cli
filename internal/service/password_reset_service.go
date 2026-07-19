package service

import (
	"errors"
	"math/rand"
	"time"

	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/utils"
)

type PasswordResetService struct {
	userRepo          *repository.UserRepository
	passwordResetRepo *repository.PasswordResetRepository
}

func NewPasswordResetService(
	userRepo *repository.UserRepository,
	passwordResetRepo *repository.PasswordResetRepository,
) *PasswordResetService {
	return &PasswordResetService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
	}
}

func (s *PasswordResetService) GenerateResetToken(username string) (string, error) {

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = s.passwordResetRepo.DeleteByUserID(int(user.ID))
	if err != nil {
	return "", err
	}

	token := randomToken()

	reset := &model.PasswordReset{
		UserID:     int(user.ID),
		Token: token,
		ExpiredAt:  time.Now().Add(15 * time.Minute),
	}

	err = s.passwordResetRepo.CreateResetToken(reset)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *PasswordResetService) ResetPassword(
	token string,
	newPassword string,
) error {

	reset, err := s.passwordResetRepo.GetResetToken(token)
	if err != nil {
		return errors.New("invalid reset token")
	}

	if time.Now().After(reset.ExpiredAt) {
		return errors.New("reset token has expired")
	}

	hashedPassword := utils.HashPassword(newPassword)

	err = s.userRepo.UpdatePassword(
		reset.UserID,
		hashedPassword,
	)
	if err != nil {
		return err
	}

	err = s.passwordResetRepo.DeleteResetToken(reset.ID)
	if err != nil {
		return err
	}

	return nil
}

func randomToken() string {

	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, 6)

	for i := range token {
		token[i] = chars[r.Intn(len(chars))]
	}

	return string(token)
}