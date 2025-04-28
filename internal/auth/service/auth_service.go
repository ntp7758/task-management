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
	Signup(username, password string) (string, error)
	Login(username, password string) (string, error)
	CreateToken(authId string) (string, string, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Signup(username, password string) (string, error) {

	_, err := s.authRepo.FindByUsername(username)
	if err == nil {
		return "", errors.New("username is already exists")
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		return "", err
	}

	auth := model.Auth{
		ID:           primitive.NewObjectID(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Username:     username,
		Password:     hash,
		Role:         model.AuthRoleUser,
		RefreshToken: "",
	}

	_, err = s.authRepo.Insert(auth)
	if err != nil {
		return "", err
	}

	return auth.ID.Hex(), nil
}

func (s *authService) Login(username, password string) (string, error) {
	auth, err := s.authRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username")
	}

	err = security.CheckPasswordHash(password, auth.Password)
	if err != nil {
		return "", errors.New("unauthorized")
	}

	return auth.ID.Hex(), err
}

func (s *authService) CreateToken(authId string) (string, string, error) {

	token, err := security.GenerateJWTToken(authId, "")
	if err != nil {
		return "", "", err
	}

	refreshToken, err := security.GenerateRefreshToken()
	return token, refreshToken, err
}
