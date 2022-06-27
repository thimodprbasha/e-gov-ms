package rule

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"service/pkg/env"
	Logger "service/pkg/logger"
	"service/pkg/model"
	"service/pkg/util"
	"time"
)

func InitializeDB() {
	connectionNode()
	CreateRoles()
}

func CreateRoles()  {
	createAdmin("adminiuser@admin.com" , "admin123" , env.ADMIN)
	createAdmin("supervisoriuser@admin.com" , "user123" , env.SUPERVISOR)
	createAdmin("govuser@admin.com" , "user123" , env.GOV_WORKER)
}

func connectionNode() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		env.DB_HOST,
	))
	if err != nil {
		Logger.Log.Error(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		Logger.Log.Error(err)
	} else {
		Logger.Log.Infof(fmt.Sprintf("DB %s Connection succesfull !", env.DBName))
	}
	//env.MClient = client
	env.MDB = client.Database(env.DBName)

}

func createAdmin(email , password , role string  ) {
	db := env.MDB

	user := model.User{}

	adminID := xid.New()
	currentTime := time.Now()
	user.ID = adminID.String()
	user.MID_ = primitive.NewObjectID().Hex()
	user.CreatedAt = &currentTime
	user.UpdatedAt = &currentTime

	hashedPassword, errHash := util.GetHashedPassword(password)
	if errHash != nil {
		Logger.Log.Error(errHash)
	}

	user.Email = email
	user.Password = hashedPassword
	user.Role = role

	query := bson.D{{"email", email}}

	if isExists, err := util.IsUserExits(&query, db); err != nil {
		Logger.Log.Error(err)
	} else if isExists {
		Logger.Log.Infof("Defualt admin already created")
	} else {
		_, err := db.Collection("User").InsertOne(context.Background(), user)
		if err != nil {
			Logger.Log.Error(err)
		}
		Logger.Log.Infof("Deafult admin created")
	}

}
