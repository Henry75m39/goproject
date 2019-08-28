package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/cost_models"
	"GRM/src/tms-srv/wrapper/project_group"
	"GRM/src/tms-srv/wrapper/projects"
	"GRM/src/tms-srv/wrapper/ws_files"
	"GRM/src/tms-srv/wrapper/ws_tasks"
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

func ProjectsGroupCreateCaller(c *gin.Context) {
	var projectGroup entity.ProjectGroup
	err := c.ShouldBindJSON(&projectGroup)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := project_group.CreateProjectGroup(projectGroup)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}

}

func CostModelsCaller(c *gin.Context) {
	var costModels entity.CostModels
	err := c.ShouldBindJSON(&costModels)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := cost_models.CostModels(costModels)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}

}

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

func TasksCaller(c *gin.Context) {
	var tasks entity.Tasks
	err := c.ShouldBindJSON(&tasks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := ws_tasks.GetAllTasks(tasks)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}
}
