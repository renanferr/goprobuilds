package parser

import (
	"encoding/json"
	"fmt"

	"github.com/anaskhan96/soup"
	goprobuilds "github.com/renanferr/goprobuilds/pkg"
)

const (
	TimeClass     = "time"
	ChampClass    = "champ"
	OpponentClass = "opponent"
	PlayerClass   = "player gold"
	KDAClass      = "kda"
	GoldClass     = "_gold gold"
)

func GameFromRaw(raw *soup.Root) *goprobuilds.Game {
	a := raw.FindStrict("a")

	path := a.Attrs()["href"]

	root := a.FindStrict("div", "class", "block")
	Time := getTime(root)
	Champ := getChamp(root)
	Opponent := getOpponent(root)
	Gold := getGold(root)
	KDA := getKDA(root)
	Player := getPlayer(root)
	return &goprobuilds.Game{
		Path:     path,
		Time:     Time,
		Champ:    Champ,
		Player:   Player,
		Opponent: Opponent,
		KDA:      KDA,
		Gold:     Gold,
	}
}

func getPlayer(root soup.Root) string {
	return root.FindStrict("div", "class", PlayerClass).
		FindStrict("div", "class", "gold").
		Text()
}

func getTime(root soup.Root) string {
	return root.FindStrict("div", "class", TimeClass).
		Text()
}

func getChamp(root soup.Root) string {
	return root.FindStrict("div", "class", ChampClass).
		FindStrict("div").
		FindStrict("img").Attrs()["data-id"]
}

func getOpponent(root soup.Root) string {
	return root.FindStrict("div", "class", OpponentClass).
		FindStrict("img").Attrs()["data-id"]
}

func getKDA(root soup.Root) string {
	kdaRoot := root.FindStrict("div", "class", KDAClass)
	kills := kdaRoot.FindStrict("span", "class", "kill green").Text()
	deaths := kdaRoot.FindStrict("span", "class", "death red").Text()
	assists := kdaRoot.FindStrict("span", "class", "assists gold").Text()

	return fmt.Sprintf("%s/%s/%s", kills, deaths, assists)
}

func getGold(root soup.Root) string {
	return root.FindStrict("div", "class", GoldClass).
		Text()
}

func gamesFromRaw(raw []soup.Root) []*goprobuilds.Game {
	var games []*goprobuilds.Game
	for _, r := range raw {
		games = append(games, GameFromRaw(&r))
	}
	return games
}

func ParseGames(resp string) ([]*goprobuilds.Game, error) {
	rawGames := decodeGamesResponse(resp)
	games := gamesFromRaw(rawGames)
	return games, nil
}

func decodeGamesResponse(resp string) []soup.Root {
	decoded := []string{}
	json.Unmarshal([]byte(resp), &decoded)
	var rawGames []soup.Root
	for _, s := range decoded {
		rawGames = append(rawGames, soup.HTMLParse(s))
	}
	return rawGames
}
