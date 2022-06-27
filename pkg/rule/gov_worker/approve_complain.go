package gov_worker

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	"service/pkg/model"
)

func ApproveComplain(govWorkerID , complainID , description string) error {
	db := env.MDB
	target := model.Complain{}
	ctx := context.Background()

	err := db.Collection("Complain").FindOne(ctx, bson.M{"complain_id": complainID}).Decode(&target)
	if err != nil {
		return err
	}
	feedback := model.ComplainApproved{}
	feedback.Approved = true
	feedback.GovWorkerID = govWorkerID
	feedback.GovWorkerFeedback = description

	target.ComplainApproved = feedback

	if _, err := db.Collection("Complain").UpdateOne(context.Background(), bson.M{"complain_id": complainID}, bson.D{{"$set", target}}); err != nil {
		return err
	}

	return nil

}
