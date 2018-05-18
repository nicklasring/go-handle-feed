package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Settings struct {
	API struct {
		Name       string `json:"name"`
		URL        string `json:"url"`
		RemotePath string `json:"remote_path"`
		Hostname   string `json:"hostname"`
	} `json:"api"`
}

func getSettings() Settings {
	var settings Settings
	var settingsFile string

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	settingsFile = fmt.Sprintf("%s/.feed-settings.json", usr.HomeDir)
	f, err := os.Open(settingsFile)
	if err != nil {
		log.Fatal(err)
	}

	jsonContent, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonContent, &settings)
	return settings
}
