package models

type Movies struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
}

type Info struct {
	Search       []Movies
	TotalResults string `json:"TotalResults"`
	Response     string `json:"Response"`
}
