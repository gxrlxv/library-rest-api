package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
)

type bookService struct {
	author     Author
	user       User
	repository BookRepository
	logger     *logging.Logger
}

func NewBookService(author Author, user User, repository BookRepository, logger *logging.Logger) *bookService {
	return &bookService{author: author, user: user, repository: repository, logger: logger}
}

func (bs *bookService) CreateBook(ctx context.Context, bookDTO domain.CreateBookDTO) error {
	bs.logger.Debug("create book service")

	author, err := bs.author.GetAuthorByName(ctx, bookDTO.AuthorName)
	if err != nil {
		return err
	}

	book := domain.NewBook(bookDTO, author)
	return bs.repository.Create(ctx, book)
}

func (bs *bookService) GetBookByID(ctx context.Context, id string) (domain.Book, error) {
	bs.logger.Debug("get book by id service")
	return bs.repository.FindByID(ctx, id)
}

func (bs *bookService) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	bs.logger.Debug("get all books service")
	return bs.repository.FindAll(ctx)
}

func (bs *bookService) UpdateBook(ctx context.Context, bookDTO domain.UpdateBookDTO, id string) error {
	bs.logger.Debug("update book service")
	return bs.repository.Update(ctx, bookDTO, id)
}

func (bs *bookService) DeleteBook(ctx context.Context, id string) error {
	bs.logger.Debug("delete book service")
	return bs.repository.Delete(ctx, id)
}
