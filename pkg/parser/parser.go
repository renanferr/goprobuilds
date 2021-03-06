package parser

import (
	"encoding/json"
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

func getRawMatches(root soup.Root) []soup.Root {

	gameFeed := root.Find("div", "id", "game-feed")
	// rawMatches := filterNodes(playerFeed.FindAll("div"), func(n soup.Root, i int) bool {
	// 	class := n.Attrs()["class"]
	// 	return (class == "build-holder")
	// })
	return gameFeed.FindAll("div")
}

func parseBuilds(url string) error {
	resp, err := soup.Get(url)
	if err != nil {
		return err
	}

	rawMatches := decodeResponse(resp)

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

func filterNodes(nodes []soup.Root, fn func(soup.Root, int) bool) []soup.Root {
	var new []soup.Root
	for i, n := range nodes {
		if fn(n, i) {
			new = append(new, n)
		}
	}
	return new
}

func decodeResponse(resp string) []soup.Root {
	decoded := []string{}
	json.Unmarshal([]byte(resp), &decoded)
	var rawMatches []soup.Root
	for _, s := range decoded {
		rawMatches = append(rawMatches, soup.HTMLParse(s))
	}

	return rawMatches
}

func Parse(url string) error {
	return parseBuilds(url)
}
