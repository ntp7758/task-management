package service

import (
	"errors"
	"time"

	"github.com/ntp7758/task-management/internal/user/model"
	"github.com/ntp7758/task-management/internal/user/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Create(authId string) error
	GetByUserId(userId string) (*model.User, error)
	GetByAuthId(authId string) (*model.User, error)
}

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Create(authId string) error {

	user := model.User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		AuthId:    authId,
		Role:      model.UserRoleEmployee,
		Task:      []string{},
	}

	_, err := s.UserRepo.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByUserId(userId string) (*model.User, error) {
	user, err := s.UserRepo.FindByID(userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, nil
}

func (s *userService) GetByAuthId(authId string) (*model.User, error) {
	user, err := s.UserRepo.FindByAuthId(authId)
	if err != nil {
		return nil, errors.New("invalid auth id")
	}

	return user, nil
}
