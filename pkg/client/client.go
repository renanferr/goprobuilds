package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"path"
	"strconv"
	"strings"

	goprobuilds "github.com/renanferr/goprobuilds/pkg"
	"github.com/renanferr/goprobuilds/pkg/http"
	"github.com/renanferr/goprobuilds/pkg/parser"
)

const (
	defaultBaseURL = "https://www.probuilds.net/ajax"
	gamesPath      = "/games"
	gameVersion    = "9.18.1"
	championsURL   = "http://ddragon.leagueoflegends.com/cdn/" + gameVersion + "/data/en_US/champion.json"
)

type ChampionData struct {
	ID   string `json:"key"`
	Name string `json:"id"`
}

type Client interface {
	GetGames(champion string, limit int, sort string) ([]*goprobuilds.Game, error)
}

type client struct {
	http        http.Client
	championIDs map[string]*ChampionData
}

func NewClient() (Client, error) {
	c := &client{
		http: http.DefaultClient,
	}
	championIDs, err := getChampionIDs()
	if err != nil {
		return nil, err
	}

	c.championIDs = championIDs
	return c, nil
}

func (c *client) getGames(championID string, limit int, sort string) ([]*goprobuilds.Game, error) {
	uri, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}
	path := path.Join(uri.Path, gamesPath)
	uri.Path = path

	q := url.Values{
		"championId": []string{championID},
		"limit":      []string{strconv.Itoa(limit)},
		"sort":       []string{sort},
	}
	uri.RawQuery = q.Encode()

	log.Print(uri.String())
	resp, err := c.http.Get(uri.String())
	if err != nil {
		return nil, err
	}
	return parser.ParseGames(resp)
}

func (c *client) GetGames(champion string, limit int, sort string) ([]*goprobuilds.Game, error) {
	championID, err := c.getChampionID(champion)
	if err != nil {
		return nil, err
	}

	return c.getGames(championID, limit, sort)
}

func getChampionIDs() (map[string]*ChampionData, error) {
	resp, err := http.DefaultClient.Get(championsURL)
	if err != nil {
		return nil, err
	}
	decodedResp := struct {
		Data map[string]*ChampionData
	}{}
	json.Unmarshal([]byte(resp), &decodedResp)
	return decodedResp.Data, nil
}

func (c *client) getChampionID(champion string) (string, error) {
	data, ok := c.championIDs[strings.Title(champion)]
	if !ok {
		return "", errors.New("champion ID not found")
	}
	return data.ID, nil
}
