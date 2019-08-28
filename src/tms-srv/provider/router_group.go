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
}
