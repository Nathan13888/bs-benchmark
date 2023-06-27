package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// settings
type SettingsProperties struct {
	BATTLESNAKE_BIN string
	OutputsDir      string `json:"outputs_dir"`
	ResultsDir      string `json:"results_dir"`
	UseBrowser      bool   `json:"use_browser"`
	BoardURL        string `json:"board_url"`

	Rounds int   `json:"rounds"`
	Sizes  []int `json:"sizes"`

	Seed     string `json:"seed"`
	Timeout  int    `json:"timeout"`
	Gametype string `json:"gametype"`
	Map      string `json:"map"`
}

var Settings = SettingsProperties{
	BATTLESNAKE_BIN: "battlesnake",
	OutputsDir:      "./outputs",
	ResultsDir:      "./results",
	UseBrowser:      true,
	BoardURL:        "http://localhost:3000",

	Rounds: 10,
	Sizes:  []int{5, 7, 9, 11},

	Seed:     "1656460409268690000",
	Timeout:  500,
	Gametype: "standard",
	Map:      "standard",
}

// read settings from JSON config file in current directory
func LoadSettings() error {
	// read config file if it exists
	// open config JSON and marshal into Settings struct
	path, err := os.OpenFile("./config.json", os.O_RDONLY, 0644)
	defer path.Close()
	if os.IsNotExist(err) {
		// if no config file, write default settings to config file
		fmt.Println("Writing default settings to config.json")

		path, err := os.OpenFile("./config.json", os.O_CREATE|os.O_WRONLY, 0644)
		defer path.Close()
		if err != nil {
			return err
		}

		// write default settings to config file
		// marshal Settings struct into JSON
		marshalled, err := json.MarshalIndent(Settings, "", "  ")
		if err != nil {
			return err
		}

		// write JSON to config file
		_, err = path.WriteString(string(marshalled) + "\n")
		if err != nil {
			return err
		}
		return nil
	}

	// read config file
	fmt.Println("Reading config from file" + path.Name())
	decoder := json.NewDecoder(path)
	err = decoder.Decode(&Settings)
	if err != nil {
		return err
	}

	return nil
}
