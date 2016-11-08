package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anyweez/zombhunt/parser"
	"github.com/anyweez/zombhunt/types"
	"github.com/anyweez/zombhunt/world"
)

/**
 * Read all game data and print it to the screen. Eventually this will be a daemon process that runs
 * in the background and watches for changes but it currently only reads once and exits. Also makes
 * API requests to Steam to get player info.
 */
func main() {
	LoadConfig("src/github.com/anyweez/zombhunt/zombhunt.toml")
	w := world.Get()

	w.Players = loadPlayers()
	w.Items = loadItems()

	fmt.Println("Zombhunt")
	fmt.Printf("Players: %d\t\tItems: %d\n", len(w.Players), len(w.Items))

	for _, player := range w.Players {
		player.Inventory = parser.LoadInventory(GetConfig().Paths.SaveGame, player)
	}

	for _, player := range w.Players {
		fmt.Printf("\n%s:\n", player.Name)

		for _, item := range player.Inventory {
			fmt.Printf("  - %dx %s\n", item.Quantity, item.Name)
		}
	}
}

/**
 * Load all players that have logged into the game. This includes players who may not
 * currently be playing.
 */
func loadPlayers() []*types.Player {
	var players types.XmlPlayers
	pData, err := ioutil.ReadFile(GetConfig().Paths.PlayerData)

	if err != nil {
		log.Fatal("Can't read " + GetConfig().Paths.PlayerData)
	}

	xml.Unmarshal(pData, &players)

	out := make([]*types.Player, 0, len(players.Players))
	for _, player := range players.Players {
		extra := player.Fetch()

		out = append(out, &types.Player{
			Id:         player.Id,
			Name:       extra.PersonaName,
			ProfileUrl: extra.ProfileUrl,
		})
	}

	return out
}

func loadItems() []*types.ItemType {
	// TODO: generalize paths
	var items types.XmlItemTypes
	iData, err := ioutil.ReadFile(GetConfig().Paths.ItemData)

	if err != nil {
		log.Fatal("Can't read " + GetConfig().Paths.ItemData)
	}

	xml.Unmarshal(iData, &items)

	return items.Items
}
