package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/ybuilds/todo-api/internal/models"
)

type TodoDao interface {
	GetTodoById(id int64) (*models.Todo, error)
	GetTodos() ([]models.Todo, error)
	AddTodo(todo models.Todo) (int64, error)
	UpdateTodoById(id int64, todo models.Todo) (int64, error)
	DeleteTodoById(id int64) (int64, error)
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

func (dao *todoDao) AddTodo(todo models.Todo) (int64, error) {
	var id int64

	query := `INSERT INTO todos (name, description) VALUES ($1, $2) RETURNING id`

	err := dao.db.QueryRow(query, todo.Name, todo.Desc).Scan(&id)
	if err != nil {
		log.Println("error inserting todo to db", err)
		return -1, err
	}

	return id, nil
}

func (dao *todoDao) UpdateTodoById(id int64, todo models.Todo) (int64, error) {
	todo.Updated = time.Now()

	query := `UPDATE todos SET name=$2, description=$3, done=COALESCE($4, done), updated=$5 WHERE id=$1 RETURNING id`

	var updId int64

	err := dao.db.QueryRow(query, id, todo.Name, todo.Desc, todo.Done, todo.Updated).Scan(&updId)
	if err != nil {
		log.Println("error updating todo in db", err)
		return -1, err
	}

	return updId, nil
}

func (dao *todoDao) DeleteTodoById(id int64) (int64, error) {
	var delId int64

	query := `DELETE FROM todos WHERE id=$1 RETURNING id`

	err := dao.db.QueryRow(query, id).Scan(&delId)
	if err != nil {
		log.Println("error deleting todo from db", err)
		return -1, err
	}

	return id, nil
}
