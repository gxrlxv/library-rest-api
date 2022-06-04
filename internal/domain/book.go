package domain

type Book struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Year   int    `json:"year" bson:"year"`
	Author Author `json:"author" bson:"author"`
	Genre  string `json:"genre" bson:"genre"`
	Busy   bool   `json:"busy" bson:"busy"`
	Owner  User   `json:"owner" bson:"owner"`
}

func NewBook(name, genre string, year int, busy bool, author Author, owner User) Book {
	return Book{
		ID:     "",
		Name:   name,
		Year:   year,
		Author: author,
		Genre:  genre,
		Busy:   busy,
		Owner:  owner,
	}
}

type CreateBookDTO struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorName string `json:"author_name"`
	Genre      string `json:"genre"`
}

type UpdateBookDTO struct {
	Name       string `json:"name,omitempty"`
	Year       int    `json:"year,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Genre      string `json:"genre,omitempty"`
	Busy       bool   `json:"busy,omitempty"`
	OwnerName  string `json:"owner_name,omitempty"`
}
