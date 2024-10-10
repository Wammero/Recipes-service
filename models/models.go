package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Reiting int    `json:"reiting"`
}

type Ingredient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author_id   int    `json:"author_id"`
}
