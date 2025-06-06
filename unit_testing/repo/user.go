package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetUser(userID string) (string, error)
	AddUser(username, password string) error
}

type userRepoImp struct {
	db mongo.Database
}

func NewUserRepo(db mongo.Database) UserRepo {
	return &userRepoImp{db}
}

func (u *userRepoImp) GetUser(userID string) (string, error) {
	u.db.Collection("users").FindOne(
		context.TODO(),
		map[string]interface{}{"user_id": userID},
	)
	return "token", nil
}

func (u *userRepoImp) AddUser(username, password string) error {
	return nil
}
