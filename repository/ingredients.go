package repository

import (
	"context"
	"recipe/models"
)

func (repo *PGRepo) GetIngredients() ([]models.Ingredient, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id, name FROM recipe_ingredients")

	if err != nil {
		return nil, err
	}

	data := make([]models.Ingredient, 0)

	for rows.Next() {
		var ingredient models.Ingredient

		err := rows.Scan(
			&ingredient.Id,
			&ingredient.Name,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, ingredient)
	}

	return data, nil
}
