package handlers

import (
    //"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

var rootPath = ".." 

func resetContext(e *echo.Echo, rec *httptest.ResponseRecorder, c echo.Context) (*httptest.ResponseRecorder, echo.Context) {

	rec = httptest.NewRecorder()
	c = e.NewContext(standard.NewRequest(new(http.Request), e.Logger()), standard.NewResponse(rec, e.Logger()))

	return rec,c
}

func TestHello(t *testing.T) {

	// Setup
	e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
	e.SetRenderer(GetTemplate(rootPath))

	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	h := &Handler{}

	c.SetPath("/hello/:name")
	c.SetParamNames("name")
	c.SetParamValues("booboo")
	if assert.NoError(t, h.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello mister : booboo!", rec.Body.String())
	}

	rec, c = resetContext(e,rec,c)
	c.SetPath("/hello/")
	if assert.NoError(t, h.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello mister : John Doe!", rec.Body.String())
	}

	rec, c = resetContext(e,rec,c)
	c.SetPath("/hello")
	if assert.NoError(t, h.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello mister : John Doe!", rec.Body.String())
	}

}