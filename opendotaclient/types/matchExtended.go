package types

type OpenDotaMatchExtended struct {
	ID            int                `json:"match_id"`
	Players       [10]OpenDotaPlayer `json:"players"`
	Duration      int64              `json:"duration"`
	StartTime     int64              `json:"start_time"`
	RadiandTeamId int                `json:"radiant_team_id"`
	RadiantName   string             `json:"radiant_name"`
	DireTeamId    int                `json:"dire_team_id"`
	DireName      string             `json:"dire_name"`
	League        OpenDotaLeague     `json:"league"`
	SeriesId      int                `json:"series_id"`
	SeriesType    int                `json:"series_type"`
	RadiantScore  int                `json:"radiant_score"`
	DireScore     int                `json:"dire_score"`
	RadiantWin    bool               `json:"radiant_win"`
}
