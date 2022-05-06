package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type authorService struct {
	storage AuthorRepository
}

func (as *authorService) CreateAuthor(ctx context.Context, author domain.Author) (string, error) {
	return as.storage.Create(ctx, author)
}

func (as *authorService) GetAuthorByID(ctx context.Context, id string) (domain.Author, error) {
	return as.storage.GetOne(ctx, id)
}

func (as *authorService) GetAllAuthors(ctx context.Context) ([]domain.Author, error) {
	return as.storage.GetAll(ctx)
}
