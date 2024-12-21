package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos - Servidor iniciado")

	// Inicializar biblioteca con un libro de ejemplo
	library = append(library, books.Book{
		ID:     "1",
		Title:  "1984",
		Author: "George Orwell",
		Genre:  "Fiction",
		Year:   1949,
	})

	// Definir rutas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/books", listBooksHandler)
	http.HandleFunc("/books/add", addBookHandler)
	http.HandleFunc("/books/get", getBookHandler)
	http.HandleFunc("/books/add-concurrent", addBookConcurrentHandler)



	// Iniciar servidor
	http.ListenAndServe(":8080", nil)
}
