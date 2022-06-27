package citizen

import (
	"golang.org/x/net/context"
	"service/pkg/env"
	"service/pkg/model"
)

func Complain(complain *model.Complain) (*model.Complain , error){
	db := env.MDB
	ctx := context.Background()
	_, err := db.Collection("Complain").InsertOne(ctx, complain)

	if err != nil {
		return nil, err
	}
	return complain, err
}
