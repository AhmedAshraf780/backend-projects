package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:type`
	Repo struct{Name string    `json:name` } `json:repo`
}

func main() {
	userName := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/users/%v/events", userName)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("something went wrong")
		return
	}

	defer res.Body.Close()

	var events []Event

	if err := json.NewDecoder(res.Body).Decode(&events); err != nil{
		fmt.Println("Failed to decode the json")
		return
	}

	for _,event := range events[:7] {
		msg := fmt.Sprintf("%s has done %s at the Repository %s",userName,event.Type,event.Repo.Name)
		fmt.Println(msg)
	}

}
