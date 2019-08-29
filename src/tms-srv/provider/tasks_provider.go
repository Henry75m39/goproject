package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/ws_tasks"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func TasksClaimCaller(c *gin.Context) {
	var tasksClaim entity.TasksClaim
	err := c.ShouldBindJSON(&tasksClaim)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := ws_tasks.TasksClaim(tasksClaim)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}
}

func TasksCompleteCaller(c *gin.Context) {
	var tasksComplete entity.TasksComplete
	err := c.ShouldBindJSON(&tasksComplete)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		data, err := ws_tasks.TasksComplete(tasksComplete)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		_, _ = c.Writer.Write(data)
	}
}
