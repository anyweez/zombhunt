package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/anyweez/zombhunt/parser"
	"github.com/anyweez/zombhunt/types"
)

/**
 * Check the players file to make sure all listed players are represented in the
 * world. Shouldn't modify existing players. This will include players that aren't
 * currently logged in.
 */
func CheckPlayers(filepath string, world *types.World, wr chan watchRequest) {
	var players types.XmlPlayers
	pData, _ := ioutil.ReadFile(filepath)

	xml.Unmarshal(pData, &players)

	for _, player := range players.Players {
		// If the player doesn't exist yet, fetch their data from Steam and add them
		// to the world.
		if !world.PlayerExists(player.Id) {
			extra := player.Fetch()
			newp := types.NewPlayer()

			newp.Id = player.Id
			newp.Name = extra.PersonaName
			newp.ProfileUrl = extra.ProfileUrl
			newp.AvatarUrl = extra.AvatarUrl

			log.Printf("Adding player %s\n", extra.PersonaName)

			world.AddPlayer(newp)

			wr <- watchRequest{
				Path:    fmt.Sprintf("%s/%d.ttp", GetConfig().Paths.PlayerDirectory, player.Id),
				Handler: CheckPlayer,
			}
		}
	}
}

func CheckItems(filepath string, world *types.World, wr chan watchRequest) {
	var items types.XmlItemTypes
	iData, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err.Error() + "\nCan't read " + filepath)
	}

	xml.Unmarshal(iData, &items)

	cnt := 0
	for _, item := range items.Items {
		if !world.ItemExists(item.Id) {
			world.AddItem(item)
			cnt++
		}
	}

	log.Printf("Loaded %d new items\n", cnt)
}

func CheckRecipes(filepath string, world *types.World, wr chan watchRequest) {
	var recipes types.XmlRecipeList
	iData, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err.Error() + "\nCan't read " + filepath)
	}

	xml.Unmarshal(iData, &recipes)

	cnt := 0
	for _, recipe := range recipes.Recipes {
		if !world.RecipeExists(recipe.Name) {
			ingredients := make([]*types.Recipe, 0, len(recipe.Ingredients))

			for _, item := range recipe.Ingredients {
				full := world.GetItem(item.Name)

				if full != nil {
					ingredients = append(ingredients, &types.Recipe{
						Name:     item.Name,
						Quantity: item.Count,
					})
				}
			}

			if len(ingredients) > 0 {
				world.AddRecipe(&types.Recipe{
					Name:        recipe.Name,
					Quantity:    recipe.Count,
					Ingredients: ingredients,
				})

				cnt++
			}
		}
	}

	for i, recipe := range world.Recipes {
		world.Recipes[i] = PopulateRecipe(world, recipe)
	}

	log.Printf("Loaded %d new recipes\n", cnt)
}

/**
 * Recursively look up ingredients for a recipe.
 */
func PopulateRecipe(world *types.World, recipe *types.Recipe) *types.Recipe {
	full := world.GetRecipe(recipe.Name)

	if full == nil {
		return recipe
	}

	for i, ingredient := range full.Ingredients {
		if ingredient.Name != recipe.Name {
			full.Ingredients[i] = PopulateRecipe(world, ingredient)
		}
	}

	recipe.Ingredients = full.Ingredients
	return recipe
}

func CheckPlayer(filepath string, world *types.World, wr chan watchRequest) {
	inventory := parser.LoadInventory(filepath)

	id, _ := strconv.ParseUint(strings.Split(path.Base(filepath), ".")[0], 10, 64)
	player := world.GetPlayer(id)

	// Check to ensure that changes in quantity are reflected. Also remove items that
	// no longer exist in the inventory.
	for _, existing := range player.Inventory.Items {
		exists := false

		for _, next := range inventory.Items {
			// Same item, different quantity (needs to change)
			if existing.Id == next.Id && existing.Quantity != next.Quantity {
				player.Inventory.AddItem(types.InventoryItem{
					Id:       existing.Id,
					Name:     existing.Name,
					Quantity: next.Quantity - existing.Quantity,
				})

				exists = true
			} else if existing.Id == next.Id {
				exists = true
			}
		}

		// If an item no longer exists in the player's inventory, remove it
		if !exists {
			player.Inventory.RemoveItem(*existing)
		}
	}

	// Check to ensure that new items are added.
	for _, next := range inventory.Items {
		isNew := true

		for _, existing := range player.Inventory.Items {
			if next.Id == existing.Id {
				isNew = false
			}
		}

		if isNew {
			player.Inventory.AddItem(*next)
		}
	}
}
