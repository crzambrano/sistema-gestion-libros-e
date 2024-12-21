package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos - Servidor iniciado")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/books", booksHandler)
	http.ListenAndServe(":8080", nil)
}
