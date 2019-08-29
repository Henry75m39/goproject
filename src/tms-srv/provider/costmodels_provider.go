package provider

import (
	"GRM/src/tms-srv/entity"
	"GRM/src/tms-srv/wrapper/cost_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
