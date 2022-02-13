package services

import (
	"github.com/kier1021/go-mongodb-sample-app/api/dto"
	"github.com/kier1021/go-mongodb-sample-app/api/models"
	"github.com/kier1021/go-mongodb-sample-app/api/repositories"
)

type TodoService struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoService(todoRepository *repositories.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (srv *TodoService) GetTodos() (map[string]interface{}, error) {
	todos, err := srv.todoRepository.GetTodos()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"todos": todos,
	}, nil
}

func (srv *TodoService) GetTodoByID(id string) (map[string]interface{}, error) {
	todo, err := srv.todoRepository.GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return map[string]interface{}{
			"todo": map[string]interface{}{},
		}, nil
	}

	return map[string]interface{}{
		"todo": todo,
	}, nil
}

func (srv *TodoService) AddTodo(todo dto.AddTodoDTO) (map[string]interface{}, error) {

	lastInsertID, err := srv.todoRepository.AddTodo(models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
	})

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Todo successfully added",
		"info": map[string]interface{}{
			"_id":         lastInsertID,
			"title":       todo.Title,
			"description": todo.Description,
		},
	}, nil
}

func (srv *TodoService) DeleteTodoByID(id string) (map[string]interface{}, error) {
	err := srv.todoRepository.DeleteTodoByID(id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Todo successfully deleted",
		"_id":     id,
	}, nil
}
