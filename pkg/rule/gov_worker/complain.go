package gov_worker

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	Logger "service/pkg/logger"
	"service/pkg/model"
)

func GetComplains() ([]*model.Complain , error){
	db := env.MDB
	var target []*model.Complain
	ctx := context.Background()

	cursorUser, err := db.Collection("Complain").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursorUser.Close(ctx)
	for cursorUser.Next(ctx) {
		var data model.Complain
		if err = cursorUser.Decode(&data); err != nil {
			Logger.Log.Error(err)
			return nil, err
		} else {
			target = append(target, &data)
		}
	}
	return target, nil
}

func GetComplain(complainID string) (*model.Complain ,  error) {
	db := env.MDB
	target := model.Complain{}
	ctx := context.Background()

	err := db.Collection("Complain").FindOne(ctx, bson.M{"complain_id": complainID}).Decode(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil

}
