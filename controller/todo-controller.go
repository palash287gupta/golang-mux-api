package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/palash287gupta/golang-mux-api/errors"
	"github.com/palash287gupta/golang-mux-api/model"
	"github.com/palash287gupta/golang-mux-api/service"
)

type controller struct{}

var (
	todoService service.TodoService
)

type TodoController interface {
	GetTodos(response http.ResponseWriter, request *http.Request)
	AddTodo(response http.ResponseWriter, request *http.Request)
}

func NewTodoController(service service.TodoService) TodoController {
	todoService = service
	return &controller{}
}

func (*controller) GetTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	todos, err := todoService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the todos"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(todos)
}

func (*controller) AddTodo(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Print(&request.Body)
	var todo model.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := todoService.Validate(&todo)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := todoService.Create(&todo)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the todo"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
