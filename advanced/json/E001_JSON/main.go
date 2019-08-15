package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Repository struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Enter your GitHub username:")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = strings.Replace(username, "\n", "", -1)

	URL := "https://api.github.com/users/" + username + "/repos"
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal("GET error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("READ error")
	}

	var repos []Repository
	_ = json.Unmarshal(body, &repos)

	for _, r := range repos {
		fmt.Println(r.Name)
	}
}
