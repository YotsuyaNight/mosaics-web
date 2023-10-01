package process

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	newFilename, _ := uuid.NewRandom()
	savePath := fmt.Sprintf("uploads/%s%s", newFilename, path.Ext(file.Filename))
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	c.JSON(200, map[string]string{"uploaded": "true"})
}
