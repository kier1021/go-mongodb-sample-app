package repositories

import (
	"github.com/kier1021/go-mongodb-sample-app/api/apierrors"
	"github.com/kier1021/go-mongodb-sample-app/api/models"
)

type TodoMockRepository struct {
	todos []models.Todo
}

func NewTodoMockRepository() *TodoMockRepository {

	todos := []models.Todo{
		{
			ID:          "620a6e96b2eaf1dbec4e9b9d",
			Title:       "Todo 1",
			Description: "My first todo",
		},
		{
			ID:          "620a6ea44c37b0eae9853369",
			Title:       "Todo 2",
			Description: "My second todo",
		},
		{
			ID:          "620a6eb51c011ec537b8be2c",
			Title:       "Todo 3",
			Description: "My third todo",
		},
		{
			ID:          "620a6ec858fce40cb69f686b",
			Title:       "Todo 4",
			Description: "My fourth todo",
		},
		{
			ID:          "620a6ed3916e6ef247fef781",
			Title:       "Todo 5",
			Description: "My fifth todo",
		},
	}

	return &TodoMockRepository{
		todos: todos,
	}
}

func (repo *TodoMockRepository) GetTodos() (todos []models.Todo, err error) {
	return repo.todos, nil
}

func (repo *TodoMockRepository) GetTodoByID(todoID string) (*models.Todo, error) {
	for _, todo := range repo.todos {
		if todo.ID == todoID {
			return &todo, nil
		}
	}

	return nil, nil
}

func (repo *TodoMockRepository) AddTodo(todo models.Todo) (string, error) {
	return "620a6fa19738088a4b89c1d9", nil
}

func (repo *TodoMockRepository) DeleteTodoByID(todoID string) error {

	isExisting := false

	for _, todo := range repo.todos {
		if todo.ID == todoID {
			isExisting = true
		}
	}

	if !isExisting {
		return apierrors.NO_ENTITY_DELETED_ERROR
	}

	return nil
}

func (repo *TodoMockRepository) UpdateTodoByID(todoID string, data map[string]interface{}) error {
	isExisting := false

	for _, todo := range repo.todos {
		if todo.ID == todoID {
			isExisting = true
		}
	}

	if !isExisting {
		return apierrors.NO_ENTITY_UPDATED_ERROR
	}

	return nil
}
