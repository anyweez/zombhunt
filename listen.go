package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/anyweez/zombhunt/types"
	"github.com/anyweez/zombhunt/world"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {
	if os.Getenv("STEAM_API_KEY") == "" {
		log.Fatal("No Steam API key provided. Set STEAM_API_KEY.")
	}
}

type watchRequest struct {
	Path      string
	Handler   watchHandler
	SkipFirst bool
}

type watchHandler func(string, *types.World, chan watchRequest)

/**
 * Read all game data and print it to the screen. Eventually this will be a daemon process that runs
 * in the background and watches for changes but it currently only reads once and exits. Also makes
 * API requests to Steam to get player info.
 */
func main() {
	LoadConfig("src/github.com/anyweez/zombhunt/zombhunt.toml")
	w := world.Get()
	wr := make(chan watchRequest, 10)

	watch, err := fsnotify.NewWatcher()
	defer watch.Close()

	if err != nil {
		log.Fatal("Can't watch the filesystem.")
	}

	events := make(map[string]watchHandler)

	// Listen for change events and handle them when they come
	go func() {
		for {
			select {
			case event := <-watch.Events:
				log.Println("Reading latest version of " + path.Base(event.Name))

				events[event.Name](event.Name, w, wr)

			case err := <-watch.Errors:
				log.Println("Error: " + err.Error())
			}
		}
	}()

	// Set up watchers. There can only be one watch per file.
	go func() {
		for {
			request := <-wr
			log.Println("Watching " + path.Base(request.Path))

			events[request.Path] = request.Handler
			watch.Add(request.Path)

			// SkipFirst indicates that we should not read when we first start watching.
			// The default is to read @ first watch.
			if !request.SkipFirst {
				watch.Events <- fsnotify.Event{
					Name: request.Path,
					Op:   fsnotify.Create,
				}
			}
		}
	}()

	fmt.Println("Loading...")

	// Read the list of players and make sure all are accounted for in the world.
	wr <- watchRequest{
		Path:    GetConfig().Paths.PlayerData,
		Handler: CheckPlayers,
	}

	// Read the list of items and read in any new ones.
	wr <- watchRequest{
		Path:    GetConfig().Paths.ItemData,
		Handler: CheckItems,
	}

	wr <- watchRequest{
		Path:    GetConfig().Paths.RecipeData,
		Handler: CheckRecipes,
	}

	setupApi(w)
}

func setupApi(world *types.World) {
	router := mux.NewRouter()

	router.HandleFunc("/inventories", func(w http.ResponseWriter, r *http.Request) {
		inventories := make([]types.RESTInventory, 0)

		for _, player := range world.Players {
			inventories = append(inventories, types.RESTInventory{
				Player:    player,
				Inventory: player.Inventory,
			})
		}

		json.NewEncoder(w).Encode(world.Players)
	})

	router.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()["q"]) == 0 {
			json.NewEncoder(w).Encode(world.Recipes)
		} else {
			query := r.URL.Query()["q"][0]

			recipes := make([]*types.Recipe, 0)

			for _, recipe := range world.Recipes {
				if strings.Contains(strings.ToUpper(recipe.Name), strings.ToUpper(query)) {
					recipes = append(recipes, recipe)
				}
			}

			json.NewEncoder(w).Encode(recipes)
		}
	})

	handler := cors.Default().Handler(handlers.LoggingHandler(os.Stdout, router))
	http.ListenAndServe(":8080", handler)
}
