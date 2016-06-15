package main

import (
    "net/http"
    "io"
    // (LINUX ONLY) "github.com/facebookgo/grace/gracehttp"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"
    "text/template"
)


// Enable templating
type Template struct {
    templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


func test(c echo.Context) error {

	name := c.Param("name")
	var param = "default parameter"

	if len(name) > 0 { param = name }

    return c.Render(http.StatusOK, "test", param)
}


func main() {

	t := &Template{
	    templates: template.Must(template.ParseGlob("public/views/*.html")),
	}


    e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
	e.SetRenderer(t)


    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello index")
    })

    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello world")
    })

    // Route using template
    e.GET("/test", test)
    e.GET("/test/:name", test)

    
    // don't drop connections with stop restart
	/* LINUX 
	std := standard.New(":1323")
	std.SetHandler(e)
	gracehttp.Serve(std.Server) */

	e.Run(standard.New(":1323"))



}