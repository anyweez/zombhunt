package types

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
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

/**
 * Inventories are maps of Item ID's to InventoryItem's. There is only one entry per type of
 * item so even if there are multiple stacks of an item, only a single entry will exist here.
 */
type Inventory struct {
	Items map[uint32]*InventoryItem
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
	AvatarUrl  string

	Inventory *Inventory
}

type Recipe struct {
	Name        string
	Count       uint32
	Ingredients []*InventoryItem
}

type World struct {
	Players []*Player
	Items   []*ItemType
	Recipes []*Recipe
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

func NewPlayer() *Player {
	var p Player
	p.Inventory = NewInventory()

	return &p
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

func (w *World) PlayerExists(id uint64) bool {
	for _, player := range w.Players {
		if player.Id == id {
			return true
		}
	}

	return false
}

func (w *World) AddPlayer(p *Player) {
	w.Players = append(w.Players, p)
}

func (w *World) GetPlayer(id uint64) *Player {
	for _, player := range w.Players {
		if player.Id == id {
			return player
		}
	}

	return nil
}

func (w *World) ItemExists(id uint32) bool {
	for _, item := range w.Items {
		if item.Id == id {
			return true
		}
	}

	return false
}

func (w *World) AddItem(item *ItemType) {
	w.Items = append(w.Items, item)
}

func (w *World) GetItem(name string) *ItemType {
	for _, item := range w.Items {
		if item.Name == name {
			return item
		}
	}

	return nil
}

func (w *World) RecipeExists(name string) bool {
	for _, recipe := range w.Recipes {
		if recipe.Name == name {
			return true
		}
	}

	return false
}

func (w *World) AddRecipe(r *Recipe) {
	w.Recipes = append(w.Recipes, r)
}

func NewInventory() *Inventory {
	var inv Inventory
	inv.Items = make(map[uint32]*InventoryItem)

	return &inv
}

/**
 * Adds an item to the inventory. Multiple stacks of the same item are collapsed.
 */
func (inv *Inventory) AddItem(item InventoryItem) {
	if _, exists := inv.Items[item.Id]; exists {
		inv.Items[item.Id].Quantity += item.Quantity
	} else {
		inv.Items[item.Id] = &item
	}

	inv.checkWipe(inv.Items[item.Id])
}

func (inv *Inventory) RemoveItem(item InventoryItem) {
	if _, exists := inv.Items[item.Id]; exists {
		inv.Items[item.Id].Quantity = uint32(math.Max(0, float64(inv.Items[item.Id].Quantity-item.Quantity)))
		inv.checkWipe(inv.Items[item.Id])
	}
}

func (inv *Inventory) checkWipe(item *InventoryItem) {
	if item.Quantity == 0 {
		delete(inv.Items, item.Id)
	}
}
