package api

import (
	"fmt"
)

func (c *Client) GetArtists() ([]Artist, error) {
	var w Wrapper
	_, err := c.http.R().
		SetResult(&w).
		Get("/rest/getArtists")
	if err != nil {
		return nil, fmt.Errorf("i wruffed too loudly and sent your packet to the shadow realm 3:\nError: %v", err)
	}
	if w.SubsonicResponse.Status != "ok" {
		return nil, fmt.Errorf("navi no happi 3:\n Navi Status: %v", w.SubsonicResponse.Status)
	}

	var artists []Artist
	for _, group := range w.SubsonicResponse.Artists.Index {
		artists = append(artists, group.Artists...)
	}

	return artists, nil
}
