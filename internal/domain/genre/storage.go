package genre

import "context"

type Storage interface {
	Create(ctx context.Context, genre Genre) (string, error)
	GetOne(ctx context.Context, id string) (Genre, error)
	GetAll(ctx context.Context) ([]Genre, error)
	Update(ctx context.Context, genre Genre) error
	Delete(ctx context.Context, id string) error
}
