package env

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var MDB *mongo.Database

const DBUser = "dbuser"
const DBPwd = "12345"
const DBName = "e_gov"

var DB_HOST = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.rrmk7.mongodb.net/%s?retryWrites=true&w=majority", DBUser, DBPwd, DBName)

const REST_PORT = "8080"

const (
	CITIZEN          = "ROLE_CITIZEN"
	GOV_WORKER         = "ROLE_GOV_WORKER"
	SUPERVISOR = "ROLE_SUPERVISOR"
	ADMIN       = "ROLE_ADMIN"
)
