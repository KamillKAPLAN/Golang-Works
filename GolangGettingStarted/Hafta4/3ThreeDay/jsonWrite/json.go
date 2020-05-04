package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Catlog struct {
	ProductId string `json: "productId"`
	Quantity  int    `json: "quantity"`
}

type CatlogNodes struct {
	CatNodes []Catlog `json:"catlogNodes"`
}

func main() {
	log.Println("OKUuuuuuuuu")
	data2 := CatlogNodes{
		CatNodes: []Catlog{
			{
				ProductId: "asda",
				Quantity:  12,
			},
			{
				ProductId: "b",
				Quantity:  1,
			},
		},
	}
	json_bytelar, _ := json.Marshal(data2)
	err := ioutil.WriteFile("./jsonWrite/test.json", json_bytelar, 0777)
	if err != nil {
		// print it out
		log.Fatal(err)
	}
	log.Println("YAZzzzzzzzz")

	data := CatlogNodes{}
	_ = json.Unmarshal([]byte(json_bytelar), &data)

	for i := 0; i < len(data.CatNodes); i++ {
		log.Println("Product Id: ", data.CatNodes[i].ProductId)
		log.Println("Quantity: ", data.CatNodes[i].Quantity)
	}

}
