package todos

import (
	"log"

	"github.com/ybuilds/todo-api/internal/database"
	"github.com/ybuilds/todo-api/internal/models"
)

type Service interface {
	GetTodoById(id int64) (*models.Todo, error)
	GetTodos() ([]models.Todo, error)
	AddTodo(todo models.Todo) (int64, error)
	UpdateTodoById(id int64, todo models.Todo) (int64, error)
	DeleteTodoById(id int64) (int64, error)
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
		log.Println("error fetching todo from dao")
		return nil, err
	}

	return todo, nil
}

func (svc *service) GetTodos() ([]models.Todo, error) {
	todos, err := svc.dao.GetTodos()
	if err != nil {
		log.Println("error fetching todos from dao")
		return nil, err
	}

	return todos, nil
}

func (svc *service) AddTodo(todo models.Todo) (int64, error) {
	id, err := svc.dao.AddTodo(todo)
	if err != nil {
		log.Println("error inserting todo from dao")
		return -1, err
	}

	return id, nil
}

func (svc *service) UpdateTodoById(id int64, todo models.Todo) (int64, error) {
	id, err := svc.dao.UpdateTodoById(id, todo)
	if err != nil {
		log.Println("error updating todo from dao")
		return -1, err
	}

	return id, nil
}

func (svc *service) DeleteTodoById(id int64) (int64, error) {
	id, err := svc.dao.DeleteTodoById(id)
	if err != nil {
		log.Println("error deleting todo from dao")
		return -1, err
	}

	return id, nil
}
