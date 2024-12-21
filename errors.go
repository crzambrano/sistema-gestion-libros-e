package main

import "fmt"

// Error personalizado para libros no encontrados
type BookNotFoundError struct {
	BookID string
}

func (e *BookNotFoundError) Error() string {
	return fmt.Sprintf("El libro con ID %s no fue encontrado", e.BookID)
}
