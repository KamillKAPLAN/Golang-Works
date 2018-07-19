package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Book struct {
	Id     string
	Kisim  string
	Kucret int
	Kyayin int
	Kyazar *Bilgi
}

type Bilgi struct {
	YfirstName string
	YlastName  string
	Ybirthday  int
}

// Tipi Book olan bir (slices)dizi tanımlıyoruz...
var books []Book

// -- All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Print("List books\n")
	json.NewEncoder(w).Encode(books)
}

// -- A Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get(GET) book\n")
	params := mux.Vars(r)
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{"IdBoş", "NameBoş", 0, 0000, &Bilgi{"FirstBoş", "LastBoş", 0000}})
}

// -- Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create(POST) book\n")
	params := mux.Vars(r)
	// Tipi Book olan bir değişken tanımlıyoruz....
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// -- Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Delete(DELETE) book\n")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

// Put Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Update(PUT) book\n")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&item)
			fmt.Print(item)
			books = append(books[:index], books[index+1:]...)
			books = append([]Book{item}, books...)
			json.NewEncoder(w).Encode(item)
		}
	}
}

func Slas(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Slasss...\n")
	w.Write([]byte("{\"KAMİL\":\"KAPLAN\"}"))
}
func main() {
	fmt.Print("Server starts...\n")
	router := mux.NewRouter()
	books = append(books, Book{Id: "1", Kisim: "Toprak Ana", Kucret: 20, Kyayin: 2012, Kyazar: &Bilgi{YfirstName: "Cengiz", YlastName: "Aymatov", Ybirthday: 1928}})
	books = append(books, Book{Id: "2", Kisim: "Hayat", Kucret: 10, Kyayin: 1930, Kyazar: &Bilgi{YfirstName: "Stefan", YlastName: "Zweig", Ybirthday: 1881}})
	books = append(books, Book{Id: "3", Kisim: "asdhgasd", Kucret: 15, Kyayin: 1935, Kyazar: &Bilgi{YfirstName: "asd", YlastName: "asd", Ybirthday: 1111}})

	router.HandleFunc("/", Slas)
	router.HandleFunc("/book", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":2550", handler)
}
