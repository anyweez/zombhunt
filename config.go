package main

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type ApplicationConfig struct {
	Paths ConfigPaths
}

type ConfigPaths struct {
	SaveGame        string `toml:"Saves"`
	GameData        string `toml:"GameData"`
	PlayerData      string // Derived from raw fields
	ItemData        string // Derived from raw fields
	PlayerDirectory string
}

var config ApplicationConfig

/**
 * Load the TOML configuration file. LoadConfig() only needs to be called once,
 * then GetConfig() can be used each time after that.
 */
func LoadConfig(path string) ApplicationConfig {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.Decode(string(data), &config); err != nil {
		log.Fatal(err)
	}

	config.Paths.PlayerData = config.Paths.SaveGame + "/players.xml"
	config.Paths.ItemData = config.Paths.GameData + "/items.xml"
	config.Paths.PlayerDirectory = config.Paths.SaveGame + "/Player"

	return config
}

func GetConfig() ApplicationConfig {
	return config
}
