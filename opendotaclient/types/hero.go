package types

type Hero struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localized_name"`
}
