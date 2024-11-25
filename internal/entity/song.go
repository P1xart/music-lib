package entity

type Song struct {
	Group string `json:"group"`
	Name string `json:"name"`
	Link string `json:"link"`
	Text string `json:"text"`
	ReleaseDate string `json:"releaseDate"`
}

type SongResponse struct {
	Link string `json:"link"`
	Text string `json:"text"`
	ReleaseDate string `json:"releaseDate"`
}

type SongRequest struct {
	Group string `json:"group"`
	Name string `json:"name"`
}