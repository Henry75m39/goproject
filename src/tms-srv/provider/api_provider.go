package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/cost_models"
	"GRM/src/tms-srv/wrapper/project_group"
	"GRM/src/tms-srv/wrapper/projects"
	"GRM/src/tms-srv/wrapper/ws_files"
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

	if err := c.ShouldBindJSON(&cancelProject); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "project JSON input data is valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, _ = projects.CancelProject(cancelProject)
}

func ProjectsGroupCreateCaller(c *gin.Context) {
	var projectGroup entity.ProjectGroup

	if err := c.ShouldBindJSON(&projectGroup); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "project group JSON input data is valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, _ = project_group.CreateProjectGroup(projectGroup)
}

func CostModelsCaller(c *gin.Context) {
	var costModels entity.CostModels

	if err := c.ShouldBindJSON(&costModels); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Cost Models JSON input data is valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, _ = cost_models.CostModels(costModels)
}

func UploadFileCaller(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "A Bad File upload request! "})
	}
	_, _ = ws_files.Upload(&file, handler)
}
