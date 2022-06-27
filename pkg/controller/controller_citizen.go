package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"service/pkg/api/citizen"
)

func createCitizen(c echo.Context) error {
	result, err := citizen.CreateCitizen(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func createComplain(c echo.Context) error {
	result, err := citizen.CreateComplain(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func CitizenController(g *echo.Group) {
	g.POST("api/user/citizen/register", createCitizen)
	g.POST("api/user/citizen/complain", createComplain)
}

