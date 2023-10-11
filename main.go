package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mosaics-web/cable"
	"mosaics-web/proto"
)

func AddUserId(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Set("UserId", userId)
}

func main() {
	proto.InitGrpcClient()

	r := gin.Default()
	r.Use(cable.AddWsConnectionMap)
	r.Use(proto.AddGrpcClient)
	r.Use(AddUserId)
	InitRouter(r)
	r.Run()
}
