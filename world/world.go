/**
 * Reads and stores all information related to the world (items, players, etc).
 */

package world

import "github.com/anyweez/zombhunt/types"

// Singleton World that can be loaded once then fetched multiple times.
var w types.World

// func Load(dir string) types.World {
// 	w.Players = loadPlayers(dir + "/players.xml")
// 	w.Items = loadItems(dir)

// 	for _, player := range w.Players {
// 		player.Inventory = parser.LoadInventory(dir, player)
// 	}

// 	return w
// }

// /**
//  * Load all players that have logged into the game. This includes players who may not
//  * currently be playing.
//  */
// func loadPlayers(name string) []*types.Player {
// 	var players types.XmlPlayers
// 	pData, err := ioutil.ReadFile(name)

// 	if err != nil {
// 		log.Fatal("Can't read " + name)
// 	}

// 	xml.Unmarshal(pData, &players)

// 	out := make([]*types.Player, 0, len(players.Players))
// 	for _, player := range players.Players {
// 		extra := player.Fetch()

// 		out = append(out, &types.Player{
// 			Id:         player.Id,
// 			Name:       extra.PersonaName,
// 			ProfileUrl: extra.ProfileUrl,
// 		})
// 	}

// 	return out
// }

// func loadItems(dir string) []*types.ItemType {
// 	// TODO: generalize paths
// 	var items types.XmlItemTypes
// 	iData, err := ioutil.ReadFile(dir + "/Config/items.xml")

// 	if err != nil {
// 		log.Fatal("Can't read " + dir + "/Config/items.xml")
// 	}

// 	xml.Unmarshal(iData, &items)

// 	return items.Items
// }

func Get() *types.World {
	return &w
}
