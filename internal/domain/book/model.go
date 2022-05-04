package book

import (
	"github.com/gxrlxv/library-rest-api/internal/domain/author"
	"github.com/gxrlxv/library-rest-api/internal/domain/genre"
	"github.com/gxrlxv/library-rest-api/internal/domain/user"
)

type Book struct {
	ID     string        `json:"id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Year   int           `json:"year,omitempty"`
	Author author.Author `json:"author,omitempty"`
	Genre  genre.Genre   `json:"genre,omitempty"`
	Busy   bool          `json:"busy,omitempty"`
	Owner  user.User     `json:"owner,omitempty"`
}
