package repository

import (
	"Recipes_service/internal/models"
	"Recipes_service/pkg/password"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
)

// RegisterUser регистрирует нового пользователя
func (repo *PGRepo) RegisterUser(ctx context.Context, username, email, plainPassword string) (int64, error) {
	// Генерация хэша и соли для пароля
	hashedPassword, salt, err := password.HashPassword(plainPassword)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO users (username, email, password_hash, salt, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var userID int64
	err = repo.pool.QueryRow(ctx, query, username, email, hashedPassword, salt, time.Now(), time.Now()).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// AuthenticateUser проверяет учетные данные пользователя
func (repo *PGRepo) AuthenticateUser(ctx context.Context, username, plainPassword string) (int64, error) {
	query := `
		SELECT id, password_hash, salt FROM users
		WHERE username = $1
	`

	var userID int64
	var passwordHash, salt string

	err := repo.pool.QueryRow(ctx, query, username).Scan(&userID, &passwordHash, &salt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New("invalid username or password")
		}
		return 0, err
	}

	if !password.CheckPassword(plainPassword, salt, passwordHash) {
		return 0, errors.New("invalid username or password")
	}

	return userID, nil
}

// GetUserInfo получает информацию о пользователе по ID
func (repo *PGRepo) GetUserInfo(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT id, username, email, image_url, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := repo.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.ImageURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
