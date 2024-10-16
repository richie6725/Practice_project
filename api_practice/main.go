package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book struct(model)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init book var
var books []Book

// funtion get book
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	//loop through and find book
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	//loop through and find book
	for index, item := range books {
		if item.ID == params["id"] {

			books = append(books[:index], books[index+1:]...)
			//将前后两个子切片连接起来，跳过 index 位置，从而删除该位置的元素。
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
	json.NewEncoder(w).Encode(&Book{})
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	//loop through and find book
	for index, item := range books {
		if item.ID == params["id"] {

			books = append(books[:index], books[index+1:]...)
			//将前后两个子切片连接起来，跳过 index 位置，从而删除该位置的元素。
			break
		}
		json.NewEncoder(w).Encode(books)
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {
	// Init router
	router := mux.NewRouter()

	//data
	books = append(books, Book{ID: "1", Isbn: "12324", Title: "ONE", Author: Author{Firstname: "John", Lastname: "Coltrane"}})
	books = append(books, Book{ID: "2", Isbn: "1542", Title: "TWO", Author: Author{Firstname: "Bill", Lastname: "Evans"}})

	// Route handlers /endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))
}
