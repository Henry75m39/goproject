package provider

import (
	"GRM/src/tms-srv/entity"
	projectgroup "GRM/src/tms-srv/wrapper/project-group"
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

func ProjectsGroupCreateCaller(c *gin.Context) {
	var projectGroup entity.ProjectGroup
	if err := c.ShouldBindJSON(&projectGroup); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "project group JSON input data is valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, _ = projectgroup.CreateProjectGroup(projectGroup)
}
