package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Team struct {
	Rank     int    `json:id`
	TeamName string `json:"team_name"`
}

func getDoc() []Team {
	doc, err := goquery.NewDocument("https://www.jleague.jp/standings/j1/")
	if err != nil {
		fmt.Println(err)
	}

	var teams []Team
	doc.Find(" td.tdTeam > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span").Text()
		team := Team{
			Rank:     i + 1,
			TeamName: name,
		}
		teams = append(teams, team)
	})
	return teams
}

func main() {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"messsage": getDoc(),
			})
		})
		r.Run("localhost:8080")
}
