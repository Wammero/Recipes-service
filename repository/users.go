package repository

import (
	"context"
	"recipe/models"
)

func (repo *PGRepo) GetUsers() ([]models.User, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id, name, reiting FROM users")
	if err != nil {
		return nil, err
	}

	data := make([]models.User, 0)

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Reiting,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, user)
	}

	return data, err
}
