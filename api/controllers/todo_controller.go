package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-mongodb-sample-app/api/apierrors"
	"github.com/kier1021/go-mongodb-sample-app/api/dto"
	"github.com/kier1021/go-mongodb-sample-app/api/services"
)

type TodoController struct {
	todoSrv *services.TodoService
}

func NewTodoController(todoSrv *services.TodoService) *TodoController {
	return &TodoController{
		todoSrv: todoSrv,
	}
}

func (ctrl *TodoController) GetTodos() gin.HandlerFunc {

	return func(c *gin.Context) {
		results, err := ctrl.todoSrv.GetTodos()
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]interface{}{"error": err.Error()},
			)
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"data": results,
		})
	}
}

func (ctrl *TodoController) GetTodoByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")

		results, err := ctrl.todoSrv.GetTodoByID(id)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]interface{}{"error": err.Error()},
			)
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"data": results,
		})
	}
}

func (ctrl *TodoController) AddTodo() gin.HandlerFunc {

	return func(c *gin.Context) {

		var todo dto.AddTodoDTO

		if err := c.ShouldBind(&todo); err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]interface{}{"error": "error in data input"},
			)
			return
		}

		results, err := ctrl.todoSrv.AddTodo(todo)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]interface{}{"error": err.Error()},
			)
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"data": results,
		})
	}
}

func (ctrl *TodoController) DeleteTodoByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")

		results, err := ctrl.todoSrv.DeleteTodoByID(id)
		if err != nil {
			if errors.Is(err, apierrors.NO_ENTITY_DELETED_ERROR) {
				c.AbortWithStatusJSON(
					http.StatusBadRequest,
					map[string]interface{}{"error": fmt.Sprintf("todo with ID %s does not exists", id)},
				)
				return
			}

			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]interface{}{"error": err.Error()},
			)
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"data": results,
		})
	}
}
