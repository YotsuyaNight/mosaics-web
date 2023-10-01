package main

import (
	"net/http"

	"mosaics-web/cable"

	"github.com/gin-gonic/gin"
)

func AddUserId(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Set("UserId", userId)
}

func main() {
	r := gin.Default()
	r.Use(cable.AddWsConnectionMap)
	r.Use(AddUserId)
	InitRouter(r)
	r.Run()
}
