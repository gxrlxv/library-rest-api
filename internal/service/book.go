package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
)

type bookService struct {
	repository BookRepository
	logger     *logging.Logger
}

func NewBookService(repository BookRepository, logger *logging.Logger) *bookService {
	return &bookService{repository: repository, logger: logger}
}

func (bs *bookService) CreateBook(ctx context.Context, bookDTO domain.CreateBookDTO) error {
	bs.logger.Debug("create book service")
	book := domain.NewBook(bookDTO)
	return bs.repository.Create(ctx, book)
}

func (bs *bookService) GetBookByID(ctx context.Context, id string) (domain.Book, error) {
	return domain.Book{}, nil
}

func (bs *bookService) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	return []domain.Book{}, nil
}

func (bs *bookService) UpdateBook(ctx context.Context, bookDTO domain.UpdateBookDTO, id string) error {
	return nil
}

func (bs *bookService) DeleteBook(ctx context.Context, id string) error {
	return nil
}
