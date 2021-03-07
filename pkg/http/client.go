package http

import "github.com/anaskhan96/soup"

type Client interface {
	Get(url string) (string, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) Get(url string) (string, error) {
	return soup.Get(url)
}
