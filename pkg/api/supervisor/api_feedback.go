package supervisor

import (
	"github.com/labstack/echo/v4"
	"service/pkg/model"
	"service/pkg/rule/supervisor"
)

func ComplainFeedback(c echo.Context) error {
	complainID := c.Param("id")
	complainFeedback := model.ComplainFeedback{}

	err := c.Bind(&complainFeedback)
	if err != nil {
		return err
	}

	complainFeedback.FeedbackApproved = true
	err = supervisor.Feedback(&complainFeedback , complainID)

	return err
}
