package repository

import (
	"context"
	"recipe/models"
	"strconv"
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

func (repo *PGRepo) ChangeFavourite(userIDStr string, recipeID int) error {
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}

	_, err = repo.pool.Exec(context.Background(),
		"INSERT INTO favourites(user_id, recipe_id) VALUES($1, $2) ON CONFLICT (user_id, recipe_id) DO DELETE",
		userID, recipeID)

	if err != nil {
		return err
	}

	return nil
}
