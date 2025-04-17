package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetRandomWordFromAPI() (string, error) {
	url := "https://random-word-api.herokuapp.com/word?length=5"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result []string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result[0], err
}
