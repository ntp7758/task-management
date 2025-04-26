package repository

import (
	"context"
	"errors"

	"github.com/ntp7758/task-management/internal/auth/model"
	"github.com/ntp7758/task-management/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const authCollection = "auth"

type AuthRepository interface {
	Insert(auth model.Auth) (*mongo.InsertOneResult, error)
	FindByID(id string) (auth *model.Auth, err error)
	FindByUsername(username string) (auth *model.Auth, err error)
}

type authRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewAuthRepository(con db.MongoDBClient) (AuthRepository, error) {
	c, err := con.DB()
	if err != nil {
		return nil, err
	}

	return &authRepository{collection: c.Collection(authCollection), ctx: context.TODO()}, nil
}

func (r *authRepository) Insert(auth model.Auth) (*mongo.InsertOneResult, error) {
	_, err := r.FindByID(auth.ID.Hex())
	if err == nil {
		return nil, errors.New("some error")
	}

	result, err := r.collection.InsertOne(r.ctx, auth)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *authRepository) Update(auth model.Auth) (*mongo.UpdateResult, error) {
	return r.collection.UpdateByID(r.ctx, auth.ID, bson.M{"$set": auth})
}

func (r *authRepository) UpdateRefreshToken(id, refreshToken string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"refreshToken": refreshToken}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *authRepository) FindByID(id string) (auth *model.Auth, err error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(r.ctx, bson.M{"_id": oID})

	err = result.Decode(&auth)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *authRepository) FindByUsername(username string) (auth *model.Auth, err error) {
	result := r.collection.FindOne(r.ctx, bson.M{"username": username})

	err = result.Decode(&auth)
	if err != nil {
		return nil, err
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *authRepository) FindByRefreshToken(refreshToken string) (*model.Auth, error) {
	var auth model.Auth
	err := r.collection.FindOne(context.TODO(), bson.M{"refreshToken": refreshToken}).Decode(&auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *authRepository) Delete(id string) (*mongo.DeleteResult, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := r.collection.DeleteOne(r.ctx, bson.M{"_id": oID})
	if err != nil {
		return nil, err
	}

	return result, nil
}
