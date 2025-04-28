package repository

import (
	"context"
	"errors"

	"github.com/ntp7758/task-management/internal/user/model"
	"github.com/ntp7758/task-management/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "user"

type UserRepository interface {
	Insert(user model.User) (*mongo.InsertOneResult, error)
	FindByID(id string) (user *model.User, err error)
	FindByAuthId(authId string) (user *model.User, err error)
}

type userRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUserRepository(con db.MongoDBClient) (UserRepository, error) {
	c, err := con.DB()
	if err != nil {
		return nil, err
	}

	return &userRepository{collection: c.Collection(userCollection), ctx: context.TODO()}, nil
}

func (r *userRepository) Insert(user model.User) (*mongo.InsertOneResult, error) {
	_, err := r.FindByID(user.ID.Hex())
	if err == nil {
		return nil, errors.New("some error")
	}

	result, err := r.collection.InsertOne(r.ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) Update(user model.User) (*mongo.UpdateResult, error) {
	return r.collection.UpdateByID(r.ctx, user.ID, bson.M{"$set": user})
}

func (r *userRepository) FindByID(id string) (user *model.User, err error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(r.ctx, bson.M{"_id": oID})

	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByAuthId(authId string) (user *model.User, err error) {
	result := r.collection.FindOne(r.ctx, bson.M{"authId": authId})

	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}
