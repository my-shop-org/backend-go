package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/order/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message": "Order service is running"})
	})

	router.Run(":8082")
}
