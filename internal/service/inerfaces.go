package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type (
	Author interface {
		CreateAuthor(ctx context.Context, authorDTO domain.CreateAuthorDTO) error
		GetAuthorByID(ctx context.Context, id string) (domain.Author, error)
		GetAllAuthors(ctx context.Context) ([]domain.Author, error)
		UpdateAuthor(ctx context.Context, authorDTO domain.UpdateAuthorDTO, id string) error
		DeleteAuthor(ctx context.Context, id string) error
	}

	AuthorRepository interface {
		Create(ctx context.Context, author domain.Author) error
		FindByID(ctx context.Context, id string) (domain.Author, error)
		FindAll(ctx context.Context) ([]domain.Author, error)
		FindByName(ctx context.Context, name string) (u domain.Author, err error)
		Update(ctx context.Context, authorDTO domain.UpdateAuthorDTO, id string) error
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

	User interface {
		CreateUser(ctx context.Context, userDTO domain.CreateUserDTO) error
		SignIn(ctx context.Context, userDTO domain.SignInUserDTO) error
		GetUserByID(ctx context.Context, id string) (domain.User, error)
		GetAllUsers(ctx context.Context) ([]domain.User, error)
		UpdateUser(ctx context.Context, userDTO domain.UpdateUserDTO, id string) error
		DeleteUser(ctx context.Context, id string) error
	}

	UserRepository interface {
		Create(ctx context.Context, user domain.User) error
		FindByID(ctx context.Context, id string) (domain.User, error)
		FindByEmail(ctx context.Context, email string) (u domain.User, err error)
		FindByUsername(ctx context.Context, email string) (u domain.User, err error)
		FindAll(ctx context.Context) ([]domain.User, error)
		Update(ctx context.Context, userDTO domain.UpdateUserDTO, id string) error
		Delete(ctx context.Context, id string) error
	}
)
