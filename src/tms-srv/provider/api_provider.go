package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TMSAPIProvider() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//Create project group Rest API in Go
	router.POST("/projectgourps/create", func(c *gin.Context) {
		var pg entity.ProjectGroup

		if err := c.ShouldBind(&pg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "json data cannot parsing !!!"})
			return
		}
		x := len(pg)
		fmt.Println(x)
		y := pg
		fmt.Println(y)

		/*
			    persistToken := d.PersistToken
				name := d.Name
				description := d.Description
				projectTypeID := d.ProjectTypeID
				clientID := d.ClientID
				systemFiles := d.SystemFiles
				aisFiles := d.AisFiles
				locales := d.Locales
				attributes := d.Attributes

				projectGroup := entity.ProjectGroup{
					PersistToken: persistToken, Name: name, Description: description,
					ProjectTypeID: projectTypeID, ClientID: clientID, SystemFiles: systemFiles,
					AisFiles: aisFiles, Locales: locales, Attributes: attributes
				}


		*/
		//var resp []entity.ProjectGroup
		//resp = append(resp, pg)

		result, _ := wrapper.CreateProjectGroup(pg)

		fmt.Printf(result)

	})

	router.Run(":8080")
}
