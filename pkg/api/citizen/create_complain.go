package citizen

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"service/pkg/model"
	"service/pkg/rule/citizen"
	"time"
)

func CreateComplain(c echo.Context) (*model.Complain , error){
	complain := model.Complain{}
	err := c.Bind(&complain)
	if err != nil {
		return nil, err
	}

	ObjID := xid.New()
	currentTime := time.Now()
	complain.MID_ = primitive.NewObjectID().Hex()
	complain.ComplainID = ObjID.String()
	complain.CreatedAt = &currentTime
	complain.UpdatedAt = &currentTime

	target , err := citizen.Complain(&complain)

	return target, err

}
