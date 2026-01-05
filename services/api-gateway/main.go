package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func proxy(target string) echo.HandlerFunc {
	return func(c echo.Context) error {
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func main() {
	e := echo.New()

	e.Any("/product/*", proxy("http://product-service:8081"))

	e.Any("/order/*", proxy("http://order-service:8082"))

	e.Start(":8080")
}
