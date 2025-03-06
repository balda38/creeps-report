package types

type OpenDotaPlayer struct {
	Name        string `json:"name"`
	PersonaName string `json:"personaname"`
	Kills       int    `json:"kills"`
	Deaths      int    `json:"deaths"`
	Assists     int    `json:"assists"`
	HeroId      int    `json:"hero_id"`
	IsRadiant   bool   `json:"isRadiant"`
}
