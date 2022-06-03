package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
)

type authorService struct {
	repository AuthorRepository
	logger     *logging.Logger
}

func NewAuthorService(repository AuthorRepository, logger *logging.Logger) *authorService {
	return &authorService{repository: repository, logger: logger}
}

func (as *authorService) CreateAuthor(ctx context.Context, authorDTO domain.CreateAuthorDTO) error {
	as.logger.Debug("create author service")
	_, err := as.repository.FindByName(ctx, authorDTO.Name)
	if err == nil {
		return apperrors.ErrAuthorAlreadyExistName
	}

	author := domain.NewAuthor(authorDTO)

	return as.repository.Create(ctx, author)
}

func (as *authorService) GetAuthorByID(ctx context.Context, id string) (domain.Author, error) {
	as.logger.Debug("get author by id service")
	return as.repository.FindByID(ctx, id)
}

func (as *authorService) GetAllAuthors(ctx context.Context) ([]domain.Author, error) {
	as.logger.Debug("get all authors service")
	return as.repository.FindAll(ctx)
}

func (as *authorService) UpdateAuthor(ctx context.Context, authorDTO domain.UpdateAuthorDTO, id string) error {
	as.logger.Debug("update author service")
	return as.repository.Update(ctx, authorDTO, id)
}

func (as *authorService) DeleteAuthor(ctx context.Context, id string) error {
	as.logger.Debug("delete author service")
	return as.repository.Delete(ctx, id)
}
