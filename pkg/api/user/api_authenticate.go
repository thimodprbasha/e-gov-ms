package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"service/pkg/env"
	"service/pkg/model"
	"time"
)

type Payload struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
}


func Login(c echo.Context) error {
	request := Payload{}
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	db := env.MDB
	target := model.User{}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	err = db.Collection("User").FindOne(ctx, bson.D{{"email", request.Email}}).Decode(&target)
	defer cancel()

	if err != nil {
		message := map[string]string{
			"message": "Invalid Credentials",
		}
		return c.JSON(http.StatusUnauthorized, message)
	}

	return c.JSON(http.StatusOK, target)
}

