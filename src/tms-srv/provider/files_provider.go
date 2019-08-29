package provider

import (
	"GRM/src/tms-srv/wrapper/ws_files"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFileCaller(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "A Bad File upload request! "})
	}
	data, err := ws_files.Upload(&file, handler)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	_, _ = c.Writer.Write(data)
}
