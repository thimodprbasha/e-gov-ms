package gov_worker

import (
	"github.com/labstack/echo/v4"
	"service/pkg/rule/gov_worker"
)

func GenerateReport(c echo.Context) (interface{}, error) {
	results, err := gov_worker.GenerateReport()
	return results, err

}
