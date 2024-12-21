package main

import (
	"encoding/json"
	"net/http"
	"pkg/books"
)

var library []books.Book

// Vamos a manejar rutas adicionales para listar libros y registrar un nuevo libro.
// Listar todos los libros
func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(library)
}

// Agregar un nuevo libro
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var newBook books.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Error al decodificar datos", http.StatusBadRequest)
		return
	}

	library = append(library, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
