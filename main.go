package main

import (
	"context"
	"log"
	"net/http"
	"os"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/gen/proto/todo/v1/todov1connect"
	"todo-app/internal/common/db"
	"todo-app/internal/features/createtodo"
	"todo-app/internal/features/deletetodo"
	"todo-app/internal/features/listtodos"
	"todo-app/internal/features/updatetodo"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TodoServer struct {
	listtodos  *listtodos.Handler
	createtodo *createtodo.Handler
	updatetodo *updatetodo.Handler
	deletetodo *deletetodo.Handler
}

func (s *TodoServer) ListTodos(ctx context.Context, req *connect.Request[todov1.ListTodosRequest]) (*connect.Response[todov1.ListTodosResponse], error) {
	return s.listtodos.Handle(ctx, req)
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	return s.createtodo.Handle(ctx, req)
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *connect.Request[todov1.UpdateTodoRequest]) (*connect.Response[todov1.UpdateTodoResponse], error) {
	return s.updatetodo.Handle(ctx, req)
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *connect.Request[todov1.DeleteTodoRequest]) (*connect.Response[todov1.DeleteTodoResponse], error) {
	return s.deletetodo.Handle(ctx, req)
}

func main() {
	// Initialize Bolt database
	boltDB, err := db.NewBoltDB("data/todos.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer boltDB.Close()

	// Create data directory if it doesn't exist
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Initialize handlers
	todoServer := &TodoServer{
		listtodos:  listtodos.NewHandler(boltDB),
		createtodo: createtodo.NewHandler(boltDB),
		updatetodo: updatetodo.NewHandler(boltDB),
		deletetodo: deletetodo.NewHandler(boltDB),
	}

	path, handler := todov1connect.NewTodoServiceHandler(todoServer)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	log.Printf("Starting server on :8080...")
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
