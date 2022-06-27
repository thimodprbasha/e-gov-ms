package gov_worker

import (
	"github.com/labstack/echo/v4"
	"service/pkg/rule/gov_worker"
)

func ApproveComplain(c echo.Context) error {
	type _Payload struct {
		GovWorkerID string `json:"gov_worker_id"`
		ComplainID  string `json:"complain_id"`
		Description string `json:"description"`
	}
	payload := _Payload{}
	err := c.Bind(&payload)
	if err != nil {
		return err
	}

	err = gov_worker.ApproveComplain(payload.GovWorkerID , payload.ComplainID , payload.Description)
	return err

}
