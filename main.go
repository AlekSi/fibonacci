package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"github.com/AlekSi/fibonacci/api"
)

var (
	addrF  = flag.String("addr", "127.0.0.1:8000", "API address")
	debugF = flag.Bool("debug", false, "Enable debug mode")
)

func main() {
	flag.Parse()

	e := echo.New()
	if *debugF {
		e.SetDebug(true)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(api.ErrorHandler)

	h := api.NewHandler()
	v1 := e.Group("/api/v1")
	v1.GET("/fibonacci/:n", h.GetFibonacci)

	for _, r := range e.Routes() {
		e.Logger().Printf("%s %s -> %s", r.Method, r.Path, r.Handler)
	}
	e.Logger().Printf("Starting service at http://%s ...", *addrF)
	e.Run(standard.New(*addrF))
}
