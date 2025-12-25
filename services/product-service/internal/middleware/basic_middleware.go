package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterBasicMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			return ctx.Path() == "/"
		},
		Format: "time: ${time_rfc3339_nano}\n" +
			"method: ${method}\n" +
			"uri: ${uri}\n" +
			"status: ${status}\n" +
			"user_agent: ${user_agent}\n" +
			"latency: ${latency}\n" +
			"bytes_out: ${bytes_out}\n\n",
	}))

}
