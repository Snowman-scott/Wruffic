package api

import (
	"fmt"
)

func (c *Client) GetID() ([]Song, error) {
	var w Wrapper
	_, err := c.http.R().
		SetQueryParam("id", "2S30zbYPVC0wSs7Qr5OkHL").
		SetResult(&w).
		Get("/rest/getAlbum")
	if err != nil {
		return nil, fmt.Errorf("wruff I bwoke :D\nError: %v", err)
	}
	if w.SubsonicResponse.Status != "ok" {
		return nil, fmt.Errorf("navi status not okay\nError: %v", err)
	}
	return w.SubsonicResponse.Album.Songs, nil
}

func (c *Client) StreamURL(songID string) (string, error) {
	p, err := c.authParams()
	if err != nil {
		return "", err
	}
	p.Set("id", songID)
	p.Set("u", c.user)
	p.Set("v", "1.16.1")
	p.Set("c", "wruffic")
	return c.host + "/rest/stream?" + p.Encode(), nil
}
