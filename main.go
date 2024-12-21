package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos - Servidor iniciado")
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensaje": "Bienvenido al Sistema de Gestión de Libros Electrónicos"}`))
}
