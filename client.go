package strava

import (
	"net/http"
)

type Client struct {
	ID     int
	Secret string
	HttpDo *http.Client
}

func NewClientWithSecret(id int, s string, c *http.Client) Client {
	return Client{
		ID:     id,
		Secret: s,
		HttpDo: c,
	}
}

func NewClient(id int, c *http.Client) Client {
	return Client{
		ID:     id,
		HttpDo: c,
	}
}
