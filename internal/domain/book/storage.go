package book

import "context"

type Storage interface {
	Create(ctx context.Context, book Book) (string, error)
	GetOne(ctx context.Context, id string) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id string) error
}
