package repository

import (
	"context"
	"recipe/models"
)

func (repo *PGRepo) GetRecipes() ([]models.Recipe, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id FROM recipes")

	if err != nil {
		return nil, err
	}

	data := make([]models.Recipe, 0)

	for rows.Next() {
		var recipe models.Recipe

		err := rows.Scan(
			&recipe.Id,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, recipe)
	}

	return data, nil
}
