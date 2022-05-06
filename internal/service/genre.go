package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type genreService struct {
	storage GenreRepository
}

func (gs *genreService) CreateGenre(ctx context.Context, genre domain.Genre) (string, error) {
	return gs.storage.Create(ctx, genre)
}

func (gs *genreService) GetGenreByID(ctx context.Context, id string) (domain.Genre, error) {
	return gs.storage.GetOne(ctx, id)
}

func (gs *genreService) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	return gs.storage.GetAll(ctx)
}
