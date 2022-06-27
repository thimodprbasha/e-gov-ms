package controller

import (
	"github.com/labstack/echo/v4"
	"service/pkg/api/user"
)

func UserAuthenticationController(g *echo.Group) {
	g.POST("api/user/authenticate", user.Login)
}
