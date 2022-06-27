package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"service/pkg/controller"
	"service/pkg/env"
)

func ConfigEchoNode(e *echo.Echo) string {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/")
	controller.CitizenController(r)
	controller.GovWorkerController(r)
	controller.SupervisorController(r)
	controller.UserAuthenticationController(r)

	return env.REST_PORT
}
