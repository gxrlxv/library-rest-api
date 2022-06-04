package service

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
)

type userService struct {
	repository UserRepository
	logger     *logging.Logger
}

func NewUserService(repository UserRepository, logger *logging.Logger) *userService {
	return &userService{repository: repository, logger: logger}
}

func (us *userService) CreateUser(ctx context.Context, userDTO domain.CreateUserDTO) error {
	us.logger.Debug("create user service")
	if userDTO.Password != userDTO.RepeatPassword {
		return fmt.Errorf("passwords don't match")
	}

	_, err := us.repository.FindByEmail(ctx, userDTO.Email)
	if err == nil {
		return apperrors.ErrUserAlreadyExistEmail
	}

	_, err = us.repository.FindByUsername(ctx, userDTO.Username)
	if err == nil {
		return apperrors.ErrUserAlreadyExistName
	}

	user := domain.NewUser(userDTO)

	err = user.GeneratePasswordHash(userDTO.Password)
	if err != nil {
		return apperrors.ErrUserPasswordNotGenerated
	}

	err = us.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return err
}

func (us *userService) SignIn(ctx context.Context, userDTO domain.SignInUserDTO) error {
	us.logger.Debug("sign in user service")
	model, err := us.repository.FindByEmail(ctx, userDTO.Email)
	if err != nil {
		return apperrors.ErrUserNotFound
	}

	if err = model.CompareHashAndPassword(userDTO.Password); err != nil {
		return apperrors.ErrUserIncorrectPassword
	}

	return nil
}

func (us *userService) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	us.logger.Debug("get user by id service")
	return us.repository.FindByID(ctx, id)
}

func (us *userService) GetUserByName(ctx context.Context, name string) (domain.User, error) {
	us.logger.Debug("get user by id service")
	return us.repository.FindByUsername(ctx, name)
}

func (us *userService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	us.logger.Debug("get all users service")
	return us.repository.FindAll(ctx)
}

func (us *userService) UpdateUser(ctx context.Context, userDTO domain.UpdateUserDTO, id string) error {
	us.logger.Debug("update user service")
	return us.repository.Update(ctx, userDTO, id)
}
func (us *userService) DeleteUser(ctx context.Context, id string) error {
	us.logger.Debug("delete user service")
	return us.repository.Delete(ctx, id)
}
