package types

type OpenDotaTeam struct {
	ID            int    `json:"team_id"`
	Name          string `json:"name"`
	LastMatchTime int64  `json:"last_match_time"`
}
