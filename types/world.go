package types

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

/* An item in an inventory, belt, or chest. */
type InventoryItem struct {
	Id       uint32
	Name     string
	Quantity uint32
}

/*
	Information about a TYPE of item, not an actual instance in a player's inventory.
   	This information is consistent across all appearances of the item.
*/
type ItemType struct {
	Id   uint32 `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Player struct {
	Id         uint64
	Name       string
	ProfileUrl string

	Inventory []InventoryItem
}

type World struct {
	Players []*Player
	Items   []*ItemType
}

// TODO: reorganize. Feels sloppy to include this receiver on the XmlPlayer type.
func (p *XmlPlayer) Fetch() *steamPlayer {
	var steam steamPlayerBody
	apiKey := os.Getenv("STEAM_API_KEY")

	resp, err := http.Get("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + apiKey + "&steamids=" + strconv.FormatUint(p.Id, 10))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// TODO: handle error
	content, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(content, &steam)

	// TODO: check to ensure there is exactly one. Prioritize regions somehow?
	return steam.Response.Players[0]
}

// TODO: use map to make this O(1)
func (w *World) FindItem(id uint32) *ItemType {
	for _, item := range w.Items {
		if item.Id == id {
			return item
		}
	}

	return nil
}
