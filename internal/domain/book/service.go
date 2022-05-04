package book

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain/author"
	"github.com/gxrlxv/library-rest-api/internal/domain/genre"
	"github.com/gxrlxv/library-rest-api/internal/domain/user"
)

type Service interface {
	CreateBook(ctx context.Context, book Book) (string, error)
	GetBookByID(ctx context.Context, id string) (Book, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
}

type service struct {
	storage       Storage
	authorService author.Service
	genreService  genre.Service
	userService   user.Service
}

//func NewService(storage Storage) book.Service {
//	return &service{storage: storage}
//}

func (s *service) CreateBook(ctx context.Context, book Book) (string, error) {
	return s.storage.Create(ctx, book)
}

func (s *service) GetBookByID(ctx context.Context, id string) (Book, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetAllBooks(ctx context.Context) ([]Book, error) {
	return s.storage.GetAll(ctx)
}
