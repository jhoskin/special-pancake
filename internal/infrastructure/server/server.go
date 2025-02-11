package server

import (
	"context"
	"log"
	"net/http"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/gen/proto/todo/v1/todov1connect"
	"todo-app/internal/features/createtodo"
	"todo-app/internal/features/deletetodo"
	"todo-app/internal/features/listtodos"
	"todo-app/internal/features/updatetodo"
	"todo-app/internal/infrastructure/db"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	listtodos  *listtodos.Handler
	createtodo *createtodo.Handler
	updatetodo *updatetodo.Handler
	deletetodo *deletetodo.Handler
}

func NewServer(db *db.BoltDB) *Server {
	return &Server{
		listtodos:  listtodos.NewHandler(db),
		createtodo: createtodo.NewHandler(db),
		updatetodo: updatetodo.NewHandler(db),
		deletetodo: deletetodo.NewHandler(db),
	}
}

// ListTodos implements the TodoService interface
func (s *Server) ListTodos(ctx context.Context, req *connect.Request[todov1.ListTodosRequest]) (*connect.Response[todov1.ListTodosResponse], error) {
	return s.listtodos.Handle(ctx, req)
}

// CreateTodo implements the TodoService interface
func (s *Server) CreateTodo(ctx context.Context, req *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	return s.createtodo.Handle(ctx, req)
}

// UpdateTodo implements the TodoService interface
func (s *Server) UpdateTodo(ctx context.Context, req *connect.Request[todov1.UpdateTodoRequest]) (*connect.Response[todov1.UpdateTodoResponse], error) {
	return s.updatetodo.Handle(ctx, req)
}

// DeleteTodo implements the TodoService interface
func (s *Server) DeleteTodo(ctx context.Context, req *connect.Request[todov1.DeleteTodoRequest]) (*connect.Response[todov1.DeleteTodoResponse], error) {
	return s.deletetodo.Handle(ctx, req)
}

func (s *Server) Start(addr string) error {
	path, handler := todov1connect.NewTodoServiceHandler(s)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	log.Printf("Starting server on %s...", addr)
	return http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
