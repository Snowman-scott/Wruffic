package api

type Song struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Track      int    `json:"track"`
	Duration   int    `json:"duration"`
	DiscNumber int    `json:"discNumber"`
}

type Album struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Songs  []Song `json:"song"`
}

type Artist struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	AC    int      `json:"albumCount"`
	Roles []string `json:"roles"`
}

type ArtistIndex struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artist"`
}

type ArtistsResult struct {
	Index []ArtistIndex `json:"index"`
}

type Response struct {
	Status  string        `json:"status"`
	Artists ArtistsResult `json:"artists"`
	Album   Album         `json:"album"`
}

type Wrapper struct {
	SubsonicResponse Response `json:"subsonic-response"`
}
