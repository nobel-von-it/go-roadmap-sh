package main

import (
	"encoding/json"
	"log"
	"os"
)

type Greet struct {
	Name string `json:"name"`
}

func (g Greet) String() string {
	return "Hello, " + g.Name
}

func main() {
	file, err := os.OpenFile("file.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var greet Greet

	if err = json.NewDecoder(file).Decode(&greet); err != nil {
		log.Printf("error: %v\n content dont have json", err)
	}
	if greet.Name == "" {
		greet.Name = "World"
	}

	body, err := json.MarshalIndent(greet, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	if err = json.NewEncoder(file).Encode(greet); err != nil {
		log.Fatal(err)
	}
}
