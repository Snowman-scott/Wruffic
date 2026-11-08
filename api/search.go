package api

import (
	"fmt"
)

func (c *Client) SearchReq(query string) (SearchResult, error) {
	var w Wrapper
	_, err := c.http.R().
		SetQueryParam("query", query).
		SetResult(&w).
		Get("/rest/search3")
	if err != nil {
		return SearchResult{}, fmt.Errorf("searching for boy kissers in your area failed\nError: %v", err)
	}
	if w.SubsonicResponse.Status != "ok" {
		return SearchResult{}, fmt.Errorf("naviadwome status says no\n Navi Status: %v", err)
	}

	return w.SubsonicResponse.Search, nil
}
