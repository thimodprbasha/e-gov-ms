package gov_worker

import (
	"github.com/labstack/echo/v4"
	"service/pkg/model"
	"service/pkg/rule/gov_worker"
)

func Complains(c echo.Context) ([]*model.Complain , error) {
	target , err := gov_worker.GetComplains()
	return target, err

}
