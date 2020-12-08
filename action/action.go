package action

import (
	"github.com/PuerkitoBio/goquery"
)

type Team struct {
	Rank     int    `json:"rank"`
	TeamName string `json:"team_name"`
}

func Ranking() ([]Team, error) {
	var teams []Team
	doc, err := goquery.NewDocument("https://www.jleague.jp/standings/j1/")
	if err != nil {
		return nil, err
	}
	doc.Find("td.tdTeam > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span").Text()
		team := Team{
			Rank:     i + 1,
			TeamName: name,
		}
		teams = append(teams, team)
	})
	return teams, err
}
