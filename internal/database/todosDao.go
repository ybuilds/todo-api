package database

import (
	"database/sql"
	"log"

	"github.com/ybuilds/todo-api/internal/models"
)

type TodoDao interface {
	GetTodoById(id int64) (*models.Todo, error)
	GetTodos() ([]models.Todo, error)
	AddTodo() models.Todo
	UpdateTodoById(id int64) models.Todo
	DeleteTodoById(id int64)
}

type todoDao struct {
	db *sql.DB
}

func NewTodoDao(db *sql.DB) TodoDao {
	return &todoDao{
		db: db,
	}
}

func (dao *todoDao) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo

	query := `SELECT * FROM todos`

	rows, err := dao.db.Query(query)
	if err != nil {
		log.Println("error fetching todos from db", err)
		return nil, err
	}

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Desc, &todo.Done, &todo.Created, &todo.Updated)
		if err != nil {
			log.Println("error scanning todos from db result", err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (dao *todoDao) GetTodoById(id int64) (*models.Todo, error) {
	var todo models.Todo

	query := `SELECT * FROM todos WHERE id=$1`

	err := dao.db.QueryRow(query, id).Scan(&todo.Id, &todo.Name, &todo.Desc, &todo.Done, &todo.Created, &todo.Updated)
	if err != nil {
		log.Println("error fetching todo from db", err)
		return nil, err
	}

	return &todo, nil
}

func (dao *todoDao) AddTodo() models.Todo {
	return models.Todo{}
}

func (dao *todoDao) UpdateTodoById(id int64) models.Todo {
	return models.Todo{}
}

func (dao *todoDao) DeleteTodoById(id int64) {

}
