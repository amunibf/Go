package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//book struct /Model
type Book struct {
	ID     string  `json :"id"`
	Isbn   string  `json :"isbn"`
	Title  string  `json :"title"`
	Author *Author `json :"author"`
}

// init books var a slice book struct
var books []Book

//author struct
type Author struct {
	Firstname string `json :"firstname"`
	Lastname  string `json :"lastname"`
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get one books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get params
	params := mux.Vars(r)
	// loop thru books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create one book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //mock id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update all books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func main() {
	// fmt.Println("Server started!")
	//init router
	r := mux.NewRouter()

	//mock data @todo implement DB
	books = append(books, Book{ID: "1", Isbn: "61151", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Wayne"}})
	books = append(books, Book{ID: "2", Isbn: "61152", Title: "Book Two", Author: &Author{Firstname: "Harold", Lastname: "Crick"}})

	//endpoints
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
