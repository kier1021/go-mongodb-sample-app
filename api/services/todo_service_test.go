package services

import (
	"reflect"
	"testing"

	"github.com/kier1021/go-mongodb-sample-app/api/dto"
	"github.com/kier1021/go-mongodb-sample-app/api/models"
	"github.com/kier1021/go-mongodb-sample-app/api/repositories"
)

func TestGetTodos(t *testing.T) {

	type testCase struct {
		name     string
		expected map[string]interface{}
		hasError bool
	}

	todoMockRepo := repositories.NewTodoMockRepository()
	todoSrv := NewTodoService(todoMockRepo)

	testCases := []testCase{
		{
			name: "Should return all todos",
			expected: map[string]interface{}{
				"todos": []models.Todo{
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
				},
			},
			hasError: false,
		},
	}

	for _, testCase := range testCases {

		res, err := todoSrv.GetTodos()

		if testCase.hasError { // Check if there is an expected error
			if err == nil {
				t.Errorf("Test '%s' failed. Expecting an error got nil", testCase.name)
			}
		} else { // Check if there is an unexpected error
			if err != nil {
				t.Errorf("Test '%s' failed. An unexpected error occurred", testCase.name)
			}
		}

		if !reflect.DeepEqual(res, testCase.expected) {
			t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", testCase.name, testCase.expected, res)
		}

	}
}

func TestGetTodoByID(t *testing.T) {

	type testCase struct {
		name     string
		expected map[string]interface{}
		todoID   string
		hasError bool
	}

	todoMockRepo := repositories.NewTodoMockRepository()
	todoSrv := NewTodoService(todoMockRepo)

	testCases := []testCase{
		{
			name: "Should return the correct todo given a correct ID",
			expected: map[string]interface{}{
				"todo": &models.Todo{
					ID:          "620a6e96b2eaf1dbec4e9b9d",
					Title:       "Todo 1",
					Description: "My first todo",
				},
			},
			todoID:   "620a6e96b2eaf1dbec4e9b9d",
			hasError: false,
		},
		{
			name: "Should return an empty todo given a wrong ID",
			expected: map[string]interface{}{
				"todo": map[string]interface{}{},
			},
			todoID:   "non_existing_id",
			hasError: false,
		},
	}

	for _, testCase := range testCases {

		res, err := todoSrv.GetTodoByID(testCase.todoID)

		if testCase.hasError { // Check if there is an expected error
			if err == nil {
				t.Errorf("Test '%s' failed. Expecting an error got nil", testCase.name)
			}
		} else { // Check if there is an unexpected error
			if err != nil {
				t.Errorf("Test '%s' failed. An unexpected error occurred", testCase.name)
			}
		}

		if !reflect.DeepEqual(res, testCase.expected) {
			t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", testCase.name, testCase.expected, res)
		}
	}
}

func TestAddTodo(t *testing.T) {
	type testCase struct {
		name     string
		expected map[string]interface{}
		todo     dto.AddTodoDTO
		hasError bool
	}

	todoMockRepo := repositories.NewTodoMockRepository()
	todoSrv := NewTodoService(todoMockRepo)

	testCases := []testCase{
		{
			name: "Should successfully add todo",
			expected: map[string]interface{}{
				"message": "Todo successfully added",
				"info": map[string]interface{}{
					"_id":         "620a6fa19738088a4b89c1d9",
					"title":       "Todo Test 1",
					"description": "My first test todo",
				},
			},
			todo: dto.AddTodoDTO{
				Title:       "Todo Test 1",
				Description: "My first test todo",
			},
			hasError: false,
		},
	}

	for _, testCase := range testCases {

		res, err := todoSrv.AddTodo(testCase.todo)

		if testCase.hasError { // Check if there is an expected error
			if err == nil {
				t.Errorf("Test '%s' failed. Expecting an error got nil", testCase.name)
			}
		} else { // Check if there is an unexpected error
			if err != nil {
				t.Errorf("Test '%s' failed. An unexpected error occurred", testCase.name)
			}
		}

		if !reflect.DeepEqual(res, testCase.expected) {
			t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", testCase.name, testCase.expected, res)
		}
	}
}

func TestDeleteTodoByID(t *testing.T) {
	type testCase struct {
		name     string
		expected map[string]interface{}
		hasError bool
		todoID   string
	}

	todoMockRepo := repositories.NewTodoMockRepository()
	todoSrv := NewTodoService(todoMockRepo)

	testCases := []testCase{
		{
			name: "Should successfully delete todo",
			expected: map[string]interface{}{
				"message": "Todo successfully deleted",
				"_id":     "620a6e96b2eaf1dbec4e9b9d",
			},
			hasError: false,
			todoID:   "620a6e96b2eaf1dbec4e9b9d",
		},
		{
			name:     "Should return an error when todo ID does not exists",
			expected: nil,
			hasError: true,
			todoID:   "non_existing_id",
		},
	}

	for _, testCase := range testCases {

		res, err := todoSrv.DeleteTodoByID(testCase.todoID)

		if testCase.hasError { // Check if there is an expected error
			if err == nil {
				t.Errorf("Test '%s' failed. Expecting an error got nil", testCase.name)
			}
		} else { // Check if there is an unexpected error
			if err != nil {
				t.Errorf("Test '%s' failed. An unexpected error occurred", testCase.name)
			}
		}

		if !reflect.DeepEqual(res, testCase.expected) {
			t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", testCase.name, testCase.expected, res)
		}
	}
}

func TestUpdateTodoByID(t *testing.T) {
	type testCase struct {
		name       string
		expected   map[string]interface{}
		hasError   bool
		todoID     string
		updateData dto.UpdateTodoDTO
	}

	todoMockRepo := repositories.NewTodoMockRepository()
	todoSrv := NewTodoService(todoMockRepo)

	testCases := []testCase{
		{
			name: "Should successfully update the title of todo",
			expected: map[string]interface{}{
				"message": "Todo successfully updated",
				"_id":     "620a6e96b2eaf1dbec4e9b9d",
				"info": map[string]interface{}{
					"title": "Updated Todo 1",
				},
			},
			hasError: false,
			todoID:   "620a6e96b2eaf1dbec4e9b9d",
			updateData: dto.UpdateTodoDTO{
				Title: "Updated Todo 1",
			},
		},
		{
			name: "Should successfully update the description of todo",
			expected: map[string]interface{}{
				"message": "Todo successfully updated",
				"_id":     "620a6e96b2eaf1dbec4e9b9d",
				"info": map[string]interface{}{
					"description": "Updated Todo Description 1",
				},
			},
			hasError: false,
			todoID:   "620a6e96b2eaf1dbec4e9b9d",
			updateData: dto.UpdateTodoDTO{
				Description: "Updated Todo Description 1",
			},
		},
		{
			name: "Should successfully update the both the tile and description of todo",
			expected: map[string]interface{}{
				"message": "Todo successfully updated",
				"_id":     "620a6e96b2eaf1dbec4e9b9d",
				"info": map[string]interface{}{
					"title":       "Updated Todo 1",
					"description": "Updated Todo Description 1",
				},
			},
			hasError: false,
			todoID:   "620a6e96b2eaf1dbec4e9b9d",
			updateData: dto.UpdateTodoDTO{
				Title:       "Updated Todo 1",
				Description: "Updated Todo Description 1",
			},
		},
		{
			name:     "Should return an error when todo ID does not exists",
			expected: nil,
			hasError: true,
			todoID:   "non_existing_id",
		},
	}

	for _, testCase := range testCases {

		res, err := todoSrv.UpdateTodoByID(testCase.todoID, testCase.updateData)

		if testCase.hasError { // Check if there is an expected error
			if err == nil {
				t.Errorf("Test '%s' failed. Expecting an error got nil", testCase.name)
			}
		} else { // Check if there is an unexpected error
			if err != nil {
				t.Errorf("Test '%s' failed. An unexpected error occurred", testCase.name)
			}
		}

		if !reflect.DeepEqual(res, testCase.expected) {
			t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", testCase.name, testCase.expected, res)
		}
	}
}
