package process

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessUpload(c *gin.Context) {
	result := make(map[string]string)
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
  newFilename, _ := uuid.NewRandom()
	savePath := fmt.Sprintf("uploads/%s%s", newFilename, path.Ext(file.Filename))
  err := c.SaveUploadedFile(file, savePath)
	result["process"] = savePath
	c.JSON(200, result)
}
