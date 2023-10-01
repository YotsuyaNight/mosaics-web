package helloworld

import (
	"github.com/gin-gonic/gin"
)

func ShowHelloWorld(c *gin.Context) {
	result := make(map[string]string)
	result["hello"] = "world"
	c.JSON(200, result)
}
