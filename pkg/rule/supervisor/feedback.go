package supervisor

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	"service/pkg/model"
)

func Feedback( complainFeedback *model.ComplainFeedback , complainID string) error  {
	db := env.MDB
	target := model.Complain{}
	ctx := context.Background()

	err := db.Collection("Complain").FindOne(ctx, bson.M{"complain_id": complainID}).Decode(&target)
	if err != nil {
		return  err
	}
	target.ComplainFeedback = *complainFeedback
	target.ComplainFeedback.Pending = false

	target.ComplainFeedback = *complainFeedback
	if _, err := db.Collection("Complain").UpdateOne(context.Background(), bson.M{"complain_id": complainID}, bson.D{{"$set", target}}); err != nil {
		return err
	}

	return nil
}
