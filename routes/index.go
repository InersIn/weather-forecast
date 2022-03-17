package routes

import (
	"weather-forecast/controllers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	c.Redirect(302, "/")
}

func Init() *echo.Echo {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("hk}4|MHicel;3sS@_}QlPDDp~LgPwmSZwy9RAw]Dcgs;i:`t2sIdTD{^YvmH^oCh2eEgB_JkmLiULC`?Rj]8KDa:LJo~^3>{8W_8YUEEc<WooewZGKPGpNlbhh@F"))))

	e.GET("/api/start", controllers.Start)
	e.GET("/api/question", controllers.Question)
	e.POST("/api/answer", controllers.Answer)

	e.HTTPErrorHandler = customHTTPErrorHandler

	return e
}
