package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"service/pkg/api/gov_worker"
)

func approve(c echo.Context) error {
	 err := gov_worker.ApproveComplain(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func getComplains(c echo.Context) error {
	result , err := gov_worker.Complains(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func getComplain(c echo.Context) error {
	result , err := gov_worker.Complain(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func shareComplains(c echo.Context) error {
	err := gov_worker.ShareComplain(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func GovWorkerController(g *echo.Group) {
	g.POST("api/user/gov/approve", approve)
	g.POST("api/user/gov/share_complain", shareComplains)

	g.GET("api/user/gov/complains", getComplains)
	g.GET("api/user/gov/complain/:id", getComplain)


}



