package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ybuilds/todo-api/internal/database"
	"github.com/ybuilds/todo-api/internal/todos"
)

func todoRoutes(srv *gin.Engine) {
	todosDao := database.NewTodoDao(database.DB)
	todosService := todos.NewService(todosDao)
	todosHandler := todos.NewHandler(todosService)

	srv.GET("/v1/todos", todosHandler.GetTodos)
	srv.GET("/v1/todos/:id", todosHandler.GetTodoById)
	srv.POST("/v1/todos", todosHandler.AddTodo)
	srv.PUT("/v1/todos/:id", todosHandler.UpdateTodoById)
	srv.DELETE("/v1/todos/:id", todosHandler.DeleteTodoById)
}
