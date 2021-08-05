package main

import (
	"fmt"
	"net/http"

	"github.com/palash287gupta/golang-mux-api/controller"
	router "github.com/palash287gupta/golang-mux-api/http"
	"github.com/palash287gupta/golang-mux-api/repository"
	"github.com/palash287gupta/golang-mux-api/service"
)

var (
	todoRepository repository.TodoRepository = repository.NewTodoRepository()
	todoService    service.TodoService       = service.NewTodoService(todoRepository)
	todoController controller.TodoController = controller.NewTodoController(todoService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/todos", todoController.GetTodos)
	httpRouter.POST("/todos", todoController.AddTodo)

	httpRouter.SERVE(port)
}
