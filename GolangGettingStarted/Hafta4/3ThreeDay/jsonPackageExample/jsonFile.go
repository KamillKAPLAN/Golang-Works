package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {


	f, err := os.Open("file.json") // file.json has the json content
	if err != nil {
		log.Fatal(err)
	}

	bb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	doc := make(map[string]interface{})
	err = json.Unmarshal(bb, &doc);
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("full json: %s", string(bb))
}