package gov_worker

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	"service/pkg/model"
)

func ShareComplain(supervisorID , complainID string) error {
	db := env.MDB
	target := model.User{}
	complain:= model.Complain{}
	ctx := context.Background()

	err := db.Collection("User").FindOne(ctx, bson.M{"id": supervisorID}).Decode(&target)
	if err != nil {
		return  err
	}
	target.Complain = append(target.Complain , &model.ComplainIDs{ComplainID: complainID})

	if _, err := db.Collection("User").UpdateOne(context.Background(), bson.M{"id": supervisorID}, bson.D{{"$set", target}}); err != nil {
		return err
	}

	err = db.Collection("Complain").FindOne(ctx, bson.M{"complain_id": complainID}).Decode(&complain)
	if err != nil {
		return  err
	}
	complain.ComplainFeedback.Pending = true

	if _, err := db.Collection("Complain").UpdateOne(context.Background(), bson.M{"complain_id": complainID}, bson.D{{"$set", complain}}); err != nil {
		return err
	}


	return nil
}
