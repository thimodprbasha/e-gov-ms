package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"service/pkg/api/supervisor"
)

func feedback(c echo.Context) error {
	err := supervisor.ComplainFeedback(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}


func SupervisorController(g *echo.Group) {
	g.POST("api/user/supervisor/feedback/:id", feedback)
}
