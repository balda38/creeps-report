package types

type OpenDotaMatchShort struct {
	ID            int    `json:"match_id"`
	Duration      int64  `json:"duration"`
	StartTime     int64  `json:"start_time"`
	RadiandTeamId int    `json:"radiant_team_id"`
	RadiantName   string `json:"radiant_name"`
	DireTeamId    int    `json:"dire_team_id"`
	DireName      string `json:"dire_name"`
	LeagueId      int    `json:"leagueid"`
	LeagueName    string `json:"league_name"`
	SeriesId      int    `json:"series_id"`
	SeriesType    int    `json:"series_type"`
	RadiantScore  int    `json:"radiant_score"`
	DireScore     int    `json:"dire_score"`
	RadiantWin    bool   `json:"radiant_win"`
}
