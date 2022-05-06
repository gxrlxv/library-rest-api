package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type (
	Author interface {
		CreateAuthor(ctx context.Context, author domain.Author) (string, error)
		GetAuthorByID(ctx context.Context, id string) (domain.Author, error)
		GetAllAuthors(ctx context.Context) ([]domain.Author, error)
	}

	AuthorRepository interface {
		Create(ctx context.Context, author domain.Author) (string, error)
		GetOne(ctx context.Context, id string) (domain.Author, error)
		GetAll(ctx context.Context) ([]domain.Author, error)
		Update(ctx context.Context, author domain.Author) error
		Delete(ctx context.Context, id string) error
	}

	Book interface {
		CreateBook(ctx context.Context, book domain.Book) (string, error)
		GetBookByID(ctx context.Context, id string) (domain.Book, error)
		GetAllBooks(ctx context.Context) ([]domain.Book, error)
	}

	BookRepository interface {
		Create(ctx context.Context, book domain.Book) (string, error)
		GetOne(ctx context.Context, id string) (domain.Book, error)
		GetAll(ctx context.Context) ([]domain.Book, error)
		Update(ctx context.Context, book domain.Book) error
		Delete(ctx context.Context, id string) error
	}

	Genre interface {
		CreateGenre(ctx context.Context, genre domain.Genre) (string, error)
		GetGenreByID(ctx context.Context, id string) (domain.Genre, error)
		GetAllGenres(ctx context.Context) ([]domain.Genre, error)
	}

	GenreRepository interface {
		Create(ctx context.Context, genre domain.Genre) (string, error)
		GetOne(ctx context.Context, id string) (domain.Genre, error)
		GetAll(ctx context.Context) ([]domain.Genre, error)
		Update(ctx context.Context, genre domain.Genre) error
		Delete(ctx context.Context, id string) error
	}

	User interface {
		CreateUser(ctx context.Context, user domain.User) (string, error)
		GetUserByID(ctx context.Context, id string) (domain.User, error)
		GetAllUsers(ctx context.Context) ([]domain.User, error)
	}

	UserRepository interface {
		Create(ctx context.Context, user domain.User) (string, error)
		GetOne(ctx context.Context, id string) (domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		Update(ctx context.Context, user domain.User) error
		Delete(ctx context.Context, id string) error
	}
)
