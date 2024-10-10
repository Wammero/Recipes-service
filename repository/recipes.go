package repository

import (
	"context"
	"recipe/models"
	"strconv"
)

func (repo *PGRepo) GetRecipes() ([]models.Recipe, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id, name, description, author_id FROM recipes")

	if err != nil {
		return nil, err
	}

	data := make([]models.Recipe, 0)

	for rows.Next() {
		var recipe models.Recipe

		err := rows.Scan(
			&recipe.Id,
			&recipe.Name,
			&recipe.Description,
			&recipe.Author_id,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, recipe)
	}

	return data, nil
}

func (repo *PGRepo) AddRecipe(recipt *models.Recipe) error {
	_, err := repo.pool.Exec(context.Background(), "INSERT INTO recipes (name, description, author_id) VALUES ($1, $2, $3)", recipt.Name, recipt.Description, recipt.Author_id)

	if err != nil {
		return err
	}
	return nil
}

func (repo *PGRepo) DeleteRecipe(id_str string) error {
	id, err := strconv.Atoi(id_str)

	if err != nil {
		return err
	}

	_, err = repo.pool.Exec(context.Background(), "DELETE FROM recipes WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
