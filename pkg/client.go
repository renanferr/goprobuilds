package goprobuilds

import (
	"net/url"
	"path"

	"github.com/renanferr/goprobuilds/pkg/http"
	"github.com/renanferr/goprobuilds/pkg/parser"
)

const GamesPath = "/games"

type Client interface {
	GetGames(string, int, string) ([]*Game, error)
}

type client struct {
	baseURL string
	http    http.Client
}

func NewClient(baseURL string, httpClient http.Client) Client {
	return &client{
		baseURL: baseURL,
		http:    httpClient,
	}
}

func (c *client) GetGames(championID string, limit int, sort string) ([]*Game, error) {
	gamesURL, err := url.Parse(path.Join(c.baseURL, GamesPath))
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Get(gamesURL.String())
	if err != nil {
		return nil, err
	}
	return parser.ParseGames(resp)
}
