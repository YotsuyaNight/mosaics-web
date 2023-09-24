package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		envs := make(map[string]string)
		for _, env := range os.Environ() {
			keyval := strings.SplitN(env, "=", 2)
			envs[keyval[0]] = keyval[1]
		}
		c.JSON(200, envs)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
