package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var nextUrl string
var prevUrl string

func GetMap(back bool) {
	url := "https://pokeapi.co/api/v2/location/"
	if len(nextUrl) > 0 && !back {
		url = nextUrl
	}
	if len(prevUrl) > 0 && back {
		url = prevUrl
	}

	if url == "" {
		log.Println("URL is empty, nothing to fetch.")
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Fetching URL: %s\n", url)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	if locations.Next != "" {
		nextUrl = locations.Next
	} else {
		nextUrl = ""
	}

	if locations.Previous != "" {
		prevUrl = locations.Previous
	} else {
		prevUrl = ""
	}

	// fmt.Printf("Prev URL: %v\n", prevUrl)
	// fmt.Printf("Next URL: %v\n", nextUrl)

	// Выводим список мест
	for _, value := range locations.Results {
		fmt.Println(value.Name)
	}
}
