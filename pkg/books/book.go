package books

type Book struct {
	ID     string
	Title  string
	Author string
	Genre  string
	Year   int
}

func NewBook(id, title, author, genre string, year int) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		Genre:  genre,
		Year:   year,
	}
}
