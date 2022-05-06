package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type bookService struct {
	storage       BookRepository
	authorService Author
	genreService  Genre
	userService   User
}

//func NewService(storage Storage) book.Service {
//	return &service{storage: storage}
//}

func (bs *bookService) CreateBook(ctx context.Context, book domain.Book) (string, error) {
	return bs.storage.Create(ctx, book)
}

func (bs *bookService) GetBookByID(ctx context.Context, id string) (domain.Book, error) {
	return bs.storage.GetOne(ctx, id)
}

func (bs *bookService) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	return bs.storage.GetAll(ctx)
}
