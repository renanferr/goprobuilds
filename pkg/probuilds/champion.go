package probuilds

import "github.com/anaskhan96/soup"

type Champion struct {
	Champ   *Champion
	Against *Champion
}

func ChampionFromRaw(raw *soup.Root) *Champion {
	return &Champion{}
}
