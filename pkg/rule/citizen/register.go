package citizen

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	Logger "service/pkg/logger"
	"service/pkg/model"
	"service/pkg/util"
)

func Register(user *model.User) (*model.User, error) {
	db := env.MDB

	query := bson.D{{"email", user.Email}}

	if isExists, err := util.IsUserExits(&query, db); err != nil {
		Logger.Log.Error(err)
		return nil, err

	} else if isExists {
		return nil, errors.New("email is already taken")
	}

	hashedPassword, errHash := util.GetHashedPassword(user.Password)
	if errHash != nil {
		return nil, errHash
	}
	user.Password = hashedPassword

	ctx := context.Background()

	_, errUser := db.Collection("User").InsertOne(ctx, user)
	if errUser != nil {
		Logger.Log.Error(errUser)
		return nil, errUser
	}
	return user, nil
}

