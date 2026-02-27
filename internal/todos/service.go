package todos

import (
	"log"

	"github.com/ybuilds/todo-api/internal/database"
	"github.com/ybuilds/todo-api/internal/models"
)

type Service interface {
	GetTodoById(id int64) (*models.Todo, error)
	GetTodos() []models.Todo
	AddTodo() models.Todo
	UpdateTodoById(id int64) models.Todo
	DeleteTodoById(id int64)
}

type service struct {
	dao database.TodoDao
}

func NewService(dao database.TodoDao) Service {
	return &service{
		dao: dao,
	}
}

func (svc *service) GetTodoById(id int64) (*models.Todo, error) {
	todo, err := svc.dao.GetTodoById(id)
	if err != nil {
		log.Println("error fetching todos from dao")
		return nil, err
	}
	return todo, nil
}

func (svc *service) GetTodos() []models.Todo {
	todos, err := svc.dao.GetTodos()
	if err != nil {
		log.Println("error fetching todos from dao")
		return nil
	}
	return todos
}

func (svc *service) AddTodo() models.Todo {
	return svc.dao.AddTodo()
}

func (svc *service) UpdateTodoById(id int64) models.Todo {
	return svc.dao.UpdateTodoById(id)
}

func (svc *service) DeleteTodoById(id int64) {
	svc.dao.DeleteTodoById(id)
}
