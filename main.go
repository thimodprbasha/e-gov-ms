package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"service/pkg/middleware"
	"service/pkg/rule"
)

func main() {
	rule.InitializeDB()
	RunProxy()

}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

func run() {
	e := echo.New()
	port := middleware.ConfigEchoNode(e)
	e.Logger.Fatal(e.Start(":" + port))
}