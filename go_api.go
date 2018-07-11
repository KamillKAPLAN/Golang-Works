package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
	Address   *Address
}
type Address struct {
	City  string
	State string
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	//Veri setimizle, tüm insan dilimlerinden getirelim:
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	// Tek bir veri al
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{"4", "kamil", "kaplan", &Address{}})

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// yeni bir öğe oluştur
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Bir öğeyi sil
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", FirstName: "Koko", LastName: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", FirstName: "Francis", LastName: "Sunday"})

	router.HandleFunc("/people", GetPeople).Methods("GET")            // Telefon rehberi belgesindeki herkesi al
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")       // Tek bir kişi al
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")   // Yeni bir kişi oluştur
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE") // Bir kişi sil
	log.Fatal(http.ListenAndServe(":8000", router))
}
