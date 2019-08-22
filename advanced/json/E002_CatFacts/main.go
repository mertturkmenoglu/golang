// API website:
// https://alexwohlbruck.github.io/cat-facts/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	All []Entry `json:"all"`
}

type Entry struct {
	Text    string `json:"text"`
	Upvotes int    `json:"upvotes"`
}

func main() {
	baseURL := "https://cat-fact.herokuapp.com"
	endpoint := "/facts"

	resp, err := http.Get(baseURL + endpoint)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObj Response
	err = json.Unmarshal(data, &responseObj)

	if err != nil {
		log.Fatal(err)
	}

	for i, e := range responseObj.All {
		if i > 9 {
			break
		}

		fmt.Println("Fact -", i+1, ":", e.Text)
		fmt.Println("Upvotes: ", e.Upvotes)
		fmt.Println("========")
	}
}
