package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func downloadFile(url string) error {
	cmd := exec.Command("wget", url)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func transferFile(fileName string) error {
	settings := getSettings()
	remote := fmt.Sprintf("%s:%s", settings.API.Hostname, settings.API.RemotePath)
	cmd := exec.Command("rsync", "-rvz", "--remove-sent-files", fileName, remote)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
