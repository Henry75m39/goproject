package provider

import (
	"github.com/gin-gonic/gin"
)

func ServiceRouterGroup(router *gin.Engine) {

	//group 1: This group define for project router service
	v1 := router.Group("/projects")
	{
		v1.POST("/createfortask", ProjectsCreateForTaskCaller)

	}

	//group 2: This group define for project group router service
	v2 := router.Group("/projectgroup")
	{
		v2.POST("/create", ProjectsGroupCreateCaller)
		//v2.POST("/submit", ProjectsGroupCreateCaller)
		//v2.POST("/read", readEndpoint)
	}
}
