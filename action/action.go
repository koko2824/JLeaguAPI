package action

import (
	"github.com/PuerkitoBio/goquery"
)

type TeamData struct {
	TeamName       string `json:"team_name"`
	Rank           string `json:"rank"`
	Points         string `json:"points"`
	MatchPlayed    string `json:"match_played"`
	Wim            string `json:"wim"`
	Draw           string `json:"draw"`
	Lose           string `json:"lose"`
	GoalScored     string `json:"goal_scored"`
	GoalConceded   string `json:"goal_conceded"`
	GoalDifference string `json:"goal_difference"`
	GoalAve        string `json:"goal_ave"`
	ConcededAve    string `json:"conceded_ave"`
}

func Ranking() ([]TeamData, error) {
	var teams []TeamData
	doc, err:= goquery.NewDocument("https://www.football-lab.jp/summary/team_ranking/j1/?year=2020")
	if err != nil {
		return nil, err
	}

	doc.Find("table#standing > tbody > tr").Each(func(i int, s *goquery.Selection) {
		var team TeamData
		s.Find("td").Each(func(i int, v *goquery.Selection) {
			switch i {
			case 0 : team.Rank = v.Text()
			case 1 : return //チームエンブレム
			case 2 : team.TeamName = v.Find("span.dsktp").Text()
			case 3 : team.Points = v.Text()
			case 4 : team.MatchPlayed = v.Text()
			case 5 : team.Wim = v.Text()
			case 6 : team.Draw = v.Text()
			case 7 : team.Lose = v.Text()
			case 8 : team.GoalScored = v.Text()
			case 9 : team.GoalConceded = v.Text()
			case 10: team.GoalDifference = v.Text()
			case 11: team.GoalAve = v.Text()
			case 12: team.ConcededAve = v.Text()
			}
		})
		teams = append(teams, team)
	})
	return teams, err
}