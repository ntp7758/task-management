package service

import (
	"errors"
	"time"

	"github.com/ntp7758/task-management/internal/auth/model"
	"github.com/ntp7758/task-management/internal/auth/repository"
	"github.com/ntp7758/task-management/pkg/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Signup(username, password string) error
	Login(username, password string) (string, string, error)
}

type authService struct {
	AuthRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Signup(username, password string) error {

	_, err := s.AuthRepo.FindByUsername(username)
	if err == nil {
		return errors.New("username is already exists")
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	auth := model.Auth{
		ID:           primitive.NewObjectID(),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		Username:     username,
		Password:     hash,
		Role:         model.AuthRoleUser,
		RefreshToken: "",
	}

	_, err = s.AuthRepo.Insert(auth)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(username, password string) (string, string, error) {
	auth, err := s.AuthRepo.FindByUsername(username)
	if err != nil {
		return "", "", errors.New("invalid username")
	}

	err = security.CheckPasswordHash(password, auth.Password)
	if err != nil {
		return "", "", errors.New("unauthorized")
	}

	token, err := security.GenerateJWTToken(auth.ID.Hex(), auth.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := security.GenerateRefreshToken()
	return token, refreshToken, err
}
