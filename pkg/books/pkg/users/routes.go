package main

import (
	"encoding/json"
	"net/http"
	"pkg/books"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	book := books.NewBook("1", "1984", "George Orwell", "Fiction", 1949)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
