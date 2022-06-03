package domain

type Author struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Nationality string `json:"nationality" bson:"nationality"`
}

func NewAuthor(dto CreateAuthorDTO) Author {
	return Author{
		ID:          "",
		Name:        dto.Name,
		Nationality: dto.Nationality,
	}
}

type CreateAuthorDTO struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality,omitempty"`
}

type UpdateAuthorDTO struct {
	Name        string `json:"name,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}
