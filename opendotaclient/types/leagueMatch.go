package types

type OpenDotaLeagueMatch struct {
	ID            int   `json:"match_id"`
	Duration      int64 `json:"duration"`
	StartTime     int64 `json:"start_time"`
	RadiandTeamId int   `json:"radiant_team_id"`
	DireTeamId    int   `json:"dire_team_id"`
	LeagueId      int   `json:"leagueid"`
	SeriesId      int   `json:"series_id"`
	SeriesType    int   `json:"series_type"`
	RadiantScore  int   `json:"radiant_score"`
	DireScore     int   `json:"dire_score"`
	RadiantWin    bool  `json:"radiant_win"`
}
