package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Books struct {
	Title  string `json:"Title,omitempty"`
	Author string `json:"Author,omitempty"`
	ID     string `json:"ID,omitempty"`
}

var books []Books

func main() {
	books = append(books, Books{ID: "1", Title: "Panipat", Author: "Raja Paranjape"})
	books = append(books, Books{ID: "2", Title: "Ramayana", Author: "VyasMuni"})
	r := mux.NewRouter()
	r.HandleFunc("/books/", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", createBook).Methods("POST")
	r.HandleFunc("/books/{author}", getBookbyAuthor).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	para := mux.Vars(r)
	var book Books
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = para["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)

}

func getBookbyAuthor(w http.ResponseWriter, r *http.Request) {
	para := mux.Vars(r)
	for _, item := range books {
		if item.Author == para["author"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}
