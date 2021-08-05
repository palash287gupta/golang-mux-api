package service

import (
	"errors"
	"math/rand"

	"github.com/palash287gupta/golang-mux-api/model"
	"github.com/palash287gupta/golang-mux-api/repository"
)

type TodoService interface {
	Validate(todo *model.Todo) error
	Create(todo *model.Todo) (*model.Todo, error)
	FindAll() ([]model.Todo, error)
}

type service struct{}

var (
	repo repository.TodoRepository
)

func NewTodoService(repository repository.TodoRepository) TodoService {
	repo = repository
	return &service{}
}

func (*service) Validate(todo *model.Todo) error {
	if todo == nil {
		err := errors.New("todo is empty")
		return err
	}
	if todo.Title == "" {
		err := errors.New("todo title is empty")
		return err
	}
	return nil
}

func (*service) Create(todo *model.Todo) (*model.Todo, error) {
	todo.ID = rand.Int63()
	return repo.Save(todo)
}

func (*service) FindAll() ([]model.Todo, error) {
	return repo.FindAll()
}
