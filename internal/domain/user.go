package domain

import (
	"context"
	"time"

	"://github.com"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the strict database abstractions contract
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	FetchAll(ctx context.Context) ([]User, error)
}

// UserUsecase defines the pure core business rules contract
type UserUsecase interface {
	Register(ctx context.Context, name, email string) (*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	ListUsers(ctx context.Context) ([]User, error)
}
