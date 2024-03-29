package provider

import (
	"github.com/gin-gonic/gin"
)

func ServiceRouterGroup(router *gin.Engine) {

	//group 1: This group define for project router service
	g1 := router.Group("/projects")
	{
		g1.POST("/createfortask", ProjectsCreateForTaskCaller)
		g1.POST("/cancel", CancelProjectCaller)

	}

	//group 2: This group define for project group router service
	g2 := router.Group("/projectgroup")
	{
		g2.POST("/create", ProjectsGroupCreateCaller)

	}

	//group 3: This group define for cost models router service
	g3 := router.Group("/costmodels")
	{
		g3.GET("", CostModelsCaller)
	}

	//group 4: This group define for files router service
	g4 := router.Group("/files")
	{
		g4.POST("", UploadFileCaller)
	}

	//group 5: This group define for task router service
	g5 := router.Group("/tasks")
	{
		g5.GET("", TasksCaller)
		g5.POST("/claim", TasksClaimCaller)
		g5.POST("/complete", TasksCompleteCaller)
	}
}
