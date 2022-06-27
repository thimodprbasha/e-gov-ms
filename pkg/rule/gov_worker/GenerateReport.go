package gov_worker

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"service/pkg/env"
	Logger "service/pkg/logger"
	"service/pkg/model"
)

func GenerateReport()(interface{} , error){
	db := env.MDB
	var complains []*model.Complain
	var resolvedComplains []*model.Complain
	var unResovedcomplains []*model.Complain
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
			complains = append(complains, &data)
		}
	}

	for _ , complain := range complains{
		if complain.ComplainFeedback.FeedbackApproved == true {
			resolvedComplains = append( resolvedComplains , complain)
		}else {
			unResovedcomplains = append( unResovedcomplains , complain)
		}

	}

	response := map[string]interface{}{
		"complains" : complains,
		"num_of_complains" : len(complains),
		"resolved_complains" : resolvedComplains,
		"num_of_resolved_complains" : len(resolvedComplains),
		"unresolved_complains" : unResovedcomplains,
		"num_of_unresolved_complains" : len(unResovedcomplains),
	}

	return response, nil

}
