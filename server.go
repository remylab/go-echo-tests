package main

import (
    "net/http"

    // (LINUX ONLY) "github.com/facebookgo/grace/gracehttp"

    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"

    "github.com/remylab/go-echo-tests/handlers"
)

var rootPath = "."



func main() {

    e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
	e.SetRenderer(handlers.GetTemplate(rootPath))

    e.SetHTTPErrorHandler(handlers.ErrorHandler)


    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello index")
    })

    h := &handlers.Handler{}

    // Route using template
    e.GET("/hello", h.Hello)
    e.GET("/hello/:name", h.Hello)


    
    // don't drop connections with stop restart
	/* LINUX 
	std := standard.New(":1323")
	std.SetHandler(e)
	gracehttp.Serve(std.Server) */

	e.Run(standard.New(":1323"))



}