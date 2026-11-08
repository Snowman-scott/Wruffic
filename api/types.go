package api

type Song struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	Album      string `json:"album"`
	ArtistID   string `json:"artistId"`
	AlbumID    string `json:"albumId"`
	Track      int    `json:"track"`
	Duration   int    `json:"duration"`
	DiscNumber int    `json:"discNumber"`
}

type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	ArtistID string `json:"artistId"`
	SongC    int    `json:"songCount"`
	Year     int    `json:"year"`
	Songs    []Song `json:"song"`
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

type SearchResult struct {
	Artists []Artist `json:"artist"`
	Albums  []Album  `json:"album"`
	Songs   []Song   `json:"song"`
}

type Response struct {
	Status     string        `json:"status"`
	Artists    ArtistsResult `json:"artists"`
	Album      Album         `json:"album"`
	SingleSong Song        `json:"song"`
	Search     SearchResult  `json:"searchResult3"`
}

type Wrapper struct {
	SubsonicResponse Response `json:"subsonic-response"`
}
