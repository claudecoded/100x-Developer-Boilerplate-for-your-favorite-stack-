package repository

import (
	"context"
	"database/sql"
	"errors"

	"://github.com"
	"://github.com"
)

type pgUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *pgUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user entity not found")
	}
	return &user, err
}

func (r *pgUserRepository) FetchAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
