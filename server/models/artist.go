package models

type Artist struct {
	Titles struct {
		Normalized string `json:"normalized"`
	} `json:"titles"`
	Blurb string `json:"extract"`
	Image struct {
		ImageURL string `json:"source"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"originalimage"`
	Description string `json:"description"`
}
