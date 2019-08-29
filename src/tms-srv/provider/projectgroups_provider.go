package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/project_group"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
