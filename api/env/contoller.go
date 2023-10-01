package env

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowAllEnvs(c *gin.Context) {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		keyval := strings.SplitN(env, "=", 2)
		envs[keyval[0]] = keyval[1]
	}
	c.JSON(200, envs)
}
