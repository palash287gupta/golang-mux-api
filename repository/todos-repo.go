package repository

import (
	"github.com/palash287gupta/golang-mux-api/model"
)

var todos []model.Todo

type repo struct{}

type TodoRepository interface {
	Save(todo *model.Todo) (*model.Todo, error)
	FindAll() ([]model.Todo, error)
}

func NewTodoRepository() TodoRepository {
	return &repo{}
}

func (*repo) Save(todo *model.Todo) (*model.Todo, error) {
	todos = append(todos, *todo)
	return todo, nil
}

func (*repo) FindAll() ([]model.Todo, error) {
	var list []model.Todo = todos
	return list, nil
}
