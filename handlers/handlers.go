package handlers

import (
    "io"
    "net/http"
    "text/template"
    "github.com/labstack/echo"
)


type Handler struct {}

// Enable templating
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


func GetTemplate(s string) *Template {
    return  &Template{
        templates: template.Must( template.ParseGlob(s + "/public/views/*.html") ),
    }
}
// Handlers

// /hello
func  (h *Handler)Hello(c echo.Context) error {

	name := c.Param("name")
	var param = "John Doe"

	if len(name) > 0 { param = name }

    return c.Render(http.StatusOK, "hello", param)
}

// Errors
func ErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError
    //msg := http.StatusText(code)
    he, ok := err.(*echo.HTTPError)
    if ok {

        code = he.Code
        //msg = he.Message
        switch code {
        case http.StatusNotFound:
            varmap := map[string]interface{}{
                "var1": "hello",
                "var2": "you",
            }
            c.Render(code, "404", varmap)
        default:
            // TODO handle any other case
        }
    }
}