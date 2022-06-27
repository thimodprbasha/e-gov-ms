package citizen

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"service/pkg/env"
	"service/pkg/model"
	"service/pkg/rule/citizen"
	"time"
)

func CreateCitizen(c echo.Context) (interface{}, error) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}

	user.Role = env.CITIZEN


	ObjID := xid.New()
	currentTime := time.Now()
	user.ID = ObjID.String()
	user.MID_ = primitive.NewObjectID().Hex()
	user.CreatedAt = &currentTime
	user.UpdatedAt = &currentTime


	returnValue, err := citizen.Register(&user)
	if err != nil {
		return nil, err
	}

	return returnValue, nil
}

