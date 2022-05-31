package domain

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Year   int    `json:"year,omitempty"`
	Author Author `json:"author,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Busy   bool   `json:"busy,omitempty"`
	Owner  User   `json:"owner,omitempty"`
}
