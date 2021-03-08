package main

import (
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	goprobuilds "github.com/renanferr/goprobuilds/pkg"
	"github.com/renanferr/goprobuilds/pkg/client"
	cli "github.com/urfave/cli/v2"
)

const (
	SortGameDateAsc  = "gameDate-asc"
	SortGameDateDesc = "gameDate-desc"
)

func getGames(client client.Client) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.NArg() > 0 {
			champion := c.Args().First()
			limit := c.Int("limit")
			sort := c.String("sort")
			games, err := client.GetGames(champion, limit, sort)
			if err != nil {
				return err
			}
			printGames(games)
		}
		return nil
	}
}
func main() {
	c, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "get-games",
				Usage:  "get games from the given champion",
				Action: getGames(c),
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sort",
				Value: SortGameDateDesc,
				Usage: "Sort games",
			},
			&cli.IntFlag{
				Name:  "limit",
				Value: 10,
				Usage: "Limit games in one request",
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printGames(games []*goprobuilds.Game) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Time", "Champ", "Player", "Opponent", "KDA", "Gold"})
	for i, g := range games {
		table.Append([]string{strconv.Itoa(i + 1), g.Time, g.Champ, g.Player, g.Opponent, g.KDA, g.Gold})
	}
	table.Render()
}
