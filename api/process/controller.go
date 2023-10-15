package process

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"mosaics-web/cable"
	"mosaics-web/env"
	"mosaics-web/proto"
)

func ProcessUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	newFilename := uuid.NewString() + path.Ext(file.Filename)
	saveRelPath := path.Join(env.GetUploadsDir(), newFilename)
	savePath := path.Join(env.GetBaseDir(), saveRelPath)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	cable.WriteToConnection(c, c.GetString("UserId"), "upload_finished")

	go runMosaicate(c, saveRelPath)

	c.JSON(200, map[string]string{"uploaded": "true"})
}

func runMosaicate(c *gin.Context, input string) {
	// proto.GrpcRunFileProcess(c, input, "icons", output, 16, 16)
	output, err := proto.GrpcRunFileProcess(c, input)
	if err != nil {
		cable.WriteToConnection(c, c.GetString("UserId"), fmt.Sprintf("processing_error %s", err))
	} else {
		cable.WriteToConnection(c, c.GetString("UserId"), "processing_finished: "+output)
	}
}
