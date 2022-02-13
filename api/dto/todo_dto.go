package dto

type AddTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
