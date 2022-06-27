package gov_worker

import (
	"github.com/labstack/echo/v4"
	"service/pkg/model"
	"service/pkg/rule/gov_worker"
)

func Complain(c echo.Context) (*model.Complain , error) {
	complainID := c.Param("id")
	target , err := gov_worker.GetComplain(complainID)
	return target, err
}
