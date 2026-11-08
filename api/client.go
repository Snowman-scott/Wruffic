package api

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"net/url"
	"time"

	"resty.dev/v3"
)

type Client struct {
	http *resty.Client
	host string
	user string
	pass string
}

func NewClient(host, user, pass string) *Client {
	c := &Client{host: host, user: user, pass: pass}

	c.http = resty.New().
		SetBaseURL(host).
		SetTimeout(10 *time.Second).
		SetQueryParam("u", user).
		SetQueryParam("v", "1.16.1").
		SetQueryParam("c", "wruffic").
		SetQueryParam("f", "json")

	c.http.AddRequestMiddleware(func(_ *resty.Client, r *resty.Request) error {
		p, err := c.authParams()
		if err != nil {
			return err
		}
		for k := range p {
			r.SetQueryParam(k, p.Get(k))
		}
		return nil
	})

	return c
}

func (c *Client) authParams() (url.Values, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	salt := hex.EncodeToString(b)
	sum := md5.Sum([]byte(c.pass + salt))

	p := url.Values{}
	p.Set("s", salt)
	p.Set("t", hex.EncodeToString(sum[:]))
	return p, nil
}
