package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/projects"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProjectsCreateForTaskCaller(c *gin.Context) {
	var project entity.Projects

	if err := c.ShouldBindJSON(&project); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "project JSON input data is valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, _ = projects.CreateForTask(project)
}

func CancelProjectCaller(c *gin.Context) {
	var cancelProject entity.CancelProject
	err := c.ShouldBindJSON(&cancelProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := projects.CancelProject(cancelProject)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}

}
