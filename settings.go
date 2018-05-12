package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	f, err := os.Open("settings.json")
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
