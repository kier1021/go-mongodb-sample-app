package dto

type AddTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
