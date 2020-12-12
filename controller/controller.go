package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraping/action"
)

type Controller struct{}

func (Controller) Ranking(c *gin.Context) {
	league := c.Query("league")
	year := c.Query("year")
	if league == "" {
		league = "j1"
	}
	if year == "" {
		year = "2020"
	}

	ranking, err := action.Ranking(league, year)
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
	league := c.Param("league")
	team, err := action.TeamDetail(league, year, c.Param("teamName"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teamData": team,
	})
}

func (Controller) ScoreDifference(c *gin.Context) {
	year := c.Query("year")
	league := c.Query("league")
	team1Name := c.Query("team1")
	team2Name := c.Query("team2")

	team1, team2, difference, err := action.ScoreDifference(league, year, team1Name, team2Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team1_data":       team1,
		"team2_data":       team2,
		"score_difference": difference,
	})
}
