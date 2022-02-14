package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-mongodb-sample-app/api/controllers"
	"github.com/kier1021/go-mongodb-sample-app/api/repositories"
	"github.com/kier1021/go-mongodb-sample-app/api/services"
	"github.com/kier1021/go-mongodb-sample-app/db"
)

type Routes struct {
	dbs    *db.DB
	Engine *gin.Engine
}

func NewRoutes(dbs *db.DB) *Routes {
	return &Routes{
		dbs:    dbs,
		Engine: gin.New(),
	}
}

func (routes *Routes) GetEngine() *gin.Engine {
	return routes.Engine
}

func (routes *Routes) SetRoutes() {

	todoRepo := repositories.NewTodoRepository(routes.dbs.MongoDB)
	todoSrv := services.NewTodoService(todoRepo)
	todoCtrl := controllers.NewTodoController(todoSrv)

	routes.Engine.GET(
		"/",
		func() gin.HandlerFunc {
			return func(c *gin.Context) {
				c.JSON(200, map[string]interface{}{"message": "Hello World!!"})
			}
		}(),
	)

	routes.Engine.GET(
		"/todos",
		todoCtrl.GetTodos(),
	)

	routes.Engine.POST(
		"/todo",
		todoCtrl.AddTodo(),
	)

	routes.Engine.GET(
		"/todo/:id",
		todoCtrl.GetTodoByID(),
	)

	routes.Engine.DELETE(
		"/todo/:id",
		todoCtrl.DeleteTodoByID(),
	)

	routes.Engine.PUT(
		"/todo/:id",
		todoCtrl.UpdateTodoByID(),
	)
}
