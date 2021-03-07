package parser

import (
	"encoding/json"

	"github.com/anaskhan96/soup"
	goprobuilds "github.com/renanferr/goprobuilds/pkg"
)

func GameFromRaw(raw *soup.Root) *goprobuilds.Game {
	root := raw.Find("a")

	return &goprobuilds.Game{
		Path:     getPath(root),
		Champ:    getChamp(root),
		Opponent: getOpponent(root),
	}
}

func getGames(raw []soup.Root) []*goprobuilds.Game {
	var matchups []*goprobuilds.Game
	for _, r := range raw {
		matchups = append(matchups, GameFromRaw(&r))
	}
	return matchups
}

func getPath(root soup.Root) string {
	return root.Attrs()["href"]
}

func getChamp(root soup.Root) string {
	return root.Find("div", "class", "block").
		Find("div", "class", "champ").
		Find("div").
		Find("img").Attrs()["data-id"]
}

func getOpponent(root soup.Root) string {
	return root.Find("div", "class", "block").
		Find("div", "class", "opponent").
		Find("img").Attrs()["data-id"]
}

func ParseGames(resp string) ([]*goprobuilds.Game, error) {

	rawGames := decodeGamesResponse(resp)

	games := getGames(rawGames)
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
