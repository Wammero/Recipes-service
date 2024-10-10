package models

type User struct {
	Id      int
	Name    string
	Reiting int
}

type Ingredient struct {
	Id   int
	Name string
}

type Recipe struct {
	Id             int
	Name           string
	Description    string
	Ingredient_ids []int
	Author_id      int
}
