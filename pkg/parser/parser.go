package parser

import (
	"log"

	"github.com/anaskhan96/soup"
	"github.com/renanferr/probuilds-go/pkg/probuilds"
)

func mapRawMatches(raw []soup.Root) []*probuilds.Match {
	var matchups []*probuilds.Match
	for _, r := range raw {
		matchups = append(matchups, matchFromRaw(&r))
	}
	return matchups
}

func matchFromRaw(raw *soup.Root) *probuilds.Match {
	root := raw.Find("a")
	// block := root.Find("div", "class", "block")

	return &probuilds.Match{
		Path:     getPath(root),
		Champ:    getChamp(root),
		Opponent: getOpponent(root),
	}
}

func getPath(root *soup.Root) string {
	return root.Attrs()["href"]
}

func getChamp(root *soup.Root) string {
	rawChampion := getChampionFromRaw(root)
	return championFromRaw(rawChampion)
}

func getChampionFromRaw(root *soup.Root) string {
	return root.Find("div", "class", "block").
		Find("div", "class", "champ")
}

func getOpponent(root *soup.Root) string {
	return getChamp(root)
}

func championFromRaw(raw *soup.Root) string {
	return raw.Find("div").
		Find("img").Atts()["data-id"]
}

func getRawMatches(root soup.Root) []soup.Root {

	templateScroller := root.Find("div", "id", "template_scroller")
	mainContent := templateScroller.Find("div", "id", "maincontent")
	container := mainContent.Find("div", "id", "container")
	module := container.Find("div", "class", "module")
	wrap := module.Find("div", "class", "wrap")
	playerFeed := wrap.Find("div", "class", "pro-player-feed-5")
	gameFeed := playerFeed.Find("div", "id", "game-feed")

	rawMatches := gameFeed.FindAll("div")
	return rawMatches
}

func parseBuilds(url string) error {
	resp, err := soup.Get(url)
	if err != nil {
		return err
	}

	doc := soup.HTMLParse(resp)
	rawMatches := getRawMatches(doc)
	matches := mapRawMatches(rawMatches)
	log.Printf("%v", matches)
	return nil
}

func parseExample(url string) error {
	// resp, err := soup.Get(url)
	// if err != nil {
	// 	return err
	// }
	// // doc := soup.HTMLParse(resp)

	// // // links := doc.Find("div", "id", "template_scroller").
	// // // 	Find("div", "id", "maincontent").
	// // // 	Find("div", "id", "container").
	// // // 	Find("div", "class", "module").
	// // // 	Find("div", "class", "wrap").
	// // // 	Find("div", "class", "pro-player-feed-5").
	// // // 	Find("div", "id", "game-feed")
	// for _, link := range links {

	// 	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	// }
	return nil
}

func Parse(url string) error {
	return parseBuilds(url)
}
