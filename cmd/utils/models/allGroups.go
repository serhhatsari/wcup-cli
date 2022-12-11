package models

type AllGroups struct {
	Groups []struct {
		Letter string `json:"letter"`
		Teams  []struct {
			Country          string `json:"country"`
			Name             string `json:"name"`
			GroupLetter      string `json:"group_letter"`
			GroupPoints      int    `json:"group_points"`
			Wins             int    `json:"wins"`
			Draws            int    `json:"draws"`
			Losses           int    `json:"losses"`
			GamesPlayed      int    `json:"games_played"`
			GoalsFor         int    `json:"goals_for"`
			GoalsAgainst     int    `json:"goals_against"`
			GoalDifferential int    `json:"goal_differential"`
		} `json:"teams"`
	} `json:"groups"`
}
