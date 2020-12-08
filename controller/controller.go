package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraping/action"
)

type Controller struct{}

func (Controller) Ranking(c *gin.Context) {
	ranking, err := action.Ranking()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ranking": ranking,
	})
}
