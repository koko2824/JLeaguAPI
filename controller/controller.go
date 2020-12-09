package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraping/action"
)

type Controller struct{}

func (Controller) Ranking(c *gin.Context) {
	year := c.Param("year")
	ranking, err := action.Ranking(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ranking": ranking,
	})
}

func (Controller) TeamDetail(c *gin.Context) {
	year := c.Param("year")
	team, err := action.TeamDetail(c.Param("teamName"), year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teamData": team,
	})
}
