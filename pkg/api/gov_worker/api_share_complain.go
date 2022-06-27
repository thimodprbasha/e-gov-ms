package gov_worker

import (
	"github.com/labstack/echo/v4"
	"service/pkg/rule/gov_worker"
)



func ShareComplain(c echo.Context) error {

	type _Payload struct {
		SupervisorID string `json:"supervisor_id"`
		ComplainID string `json:"complain_id"`
	}

	payload := _Payload{}
	err := c.Bind(&payload)
	if err != nil {
		return err
	}

	err = gov_worker.ShareComplain(payload.SupervisorID , payload.ComplainID)

	return err
}
