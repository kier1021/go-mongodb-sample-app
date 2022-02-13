package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-mongodb-sample-app/db"
)

type APIServer struct {
	router     *gin.Engine
	httpServer *http.Server
	dbs        *db.DB
}

func NewAPIServer(dbs *db.DB) *APIServer {
	routes := NewRoutes(dbs)
	routes.SetRoutes()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: routes.Engine,
	}

	return &APIServer{
		router:     routes.Engine,
		httpServer: httpServer,
		dbs:        dbs,
	}
}

func (server *APIServer) Run() error {
	if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (server *APIServer) Shutdown(ctx context.Context) error {
	if err := server.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
