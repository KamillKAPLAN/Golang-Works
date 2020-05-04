package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type databaseQuery map[string]interface{}

var CategorySlice Categories

var ProductSlice Products

var OptionSlice Options

var ChoiceSlice Choices

func Init() {
	os.Chdir("server")

	categoryFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer categoryFile.Close()
	byteValue, _ := ioutil.ReadAll(categoryFile)
	json.Unmarshal(byteValue, &CategorySlice)

	productFile, err2 := os.Open("products.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer productFile.Close()
	byteValue2, _ := ioutil.ReadAll(productFile)
	json.Unmarshal(byteValue2, &ProductSlice)

	optionFile, err3 := os.Open("options.json")
	if err3 != nil {
		fmt.Println(err3)
	}
	defer optionFile.Close()
	byteValue3, _ := ioutil.ReadAll(optionFile)
	json.Unmarshal(byteValue3, &OptionSlice)

	choiceFile, err4 := os.Open("choices.json")
	if err4 != nil {
		fmt.Println(err4)
	}
	defer choiceFile.Close()
	byteValue4, _ := ioutil.ReadAll(choiceFile)
	json.Unmarshal(byteValue4, &ChoiceSlice)
	fmt.Println(ChoiceSlice)
}
