package genre

import "context"

type Service interface {
	CreateGenre(ctx context.Context, genre Genre) (string, error)
	GetGenreByID(ctx context.Context, id string) (Genre, error)
	GetAllGenres(ctx context.Context) ([]Genre, error)
}

type service struct {
	storage Storage
}

func (s *service) CreateGenre(ctx context.Context, genre Genre) (string, error) {
	return s.storage.Create(ctx, genre)
}

func (s *service) GetGenreByID(ctx context.Context, id string) (Genre, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetAllGenres(ctx context.Context) ([]Genre, error) {
	return s.storage.GetAll(ctx)
}
