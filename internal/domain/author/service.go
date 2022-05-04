package author

import "context"

type Service interface {
	CreateAuthor(ctx context.Context, author Author) (string, error)
	GetAuthorByID(ctx context.Context, id string) (Author, error)
	GetAllAuthors(ctx context.Context) ([]Author, error)
}

type service struct {
	storage Storage
}

func (s *service) CreateAuthor(ctx context.Context, author Author) (string, error) {
	return s.storage.Create(ctx, author)
}

func (s *service) GetAuthorByID(ctx context.Context, id string) (Author, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetAllAuthors(ctx context.Context) ([]Author, error) {
	return s.storage.GetAll(ctx)
}
