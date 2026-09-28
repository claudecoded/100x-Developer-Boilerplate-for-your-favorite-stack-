package usecase

import (
	"context"
	"time"

	"://github.com"
	"://github.com"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(ctx context.Context, name, email string) (*domain.User, error) {
	now := time.Now()
	newUser := &domain.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userUsecase) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *userUsecase) ListUsers(ctx context.Context) ([]domain.User, error) {
	return u.repo.FetchAll(ctx)
}
