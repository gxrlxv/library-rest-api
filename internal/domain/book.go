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

func NewBook(dto CreateBookDTO, author Author) Book {
	return Book{
		ID:     "",
		Name:   dto.Name,
		Year:   dto.Year,
		Author: author,
		Genre:  dto.Genre,
		Busy:   false,
		Owner:  User{},
	}
}

type CreateBookDTO struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorName string `json:"author_name"`
	Genre      string `json:"genre"`
}

type UpdateBookDTO struct {
	Name   string `json:"name,omitempty"`
	Year   int    `json:"year,omitempty"`
	Author Author `json:"author,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Busy   bool   `json:"busy,omitempty"`
	Owner  User   `json:"owner,omitempty"`
}
