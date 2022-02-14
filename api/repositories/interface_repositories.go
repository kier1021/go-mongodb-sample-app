package repositories

import "github.com/kier1021/go-mongodb-sample-app/api/models"

type ITodoRepository interface {
	GetTodos() (todos []models.Todo, err error)
	GetTodoByID(todoID string) (*models.Todo, error)
	AddTodo(todo models.Todo) (string, error)
	DeleteTodoByID(todoID string) error
	UpdateTodoByID(todoID string, data map[string]interface{}) error
}
