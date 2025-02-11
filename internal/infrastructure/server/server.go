package server

import (
	"context"
	"log"
	"net/http"

	"github.com/jhoskin/special-pancake/internal/features/createtodo"
	"github.com/jhoskin/special-pancake/internal/features/deletetodo"
	"github.com/jhoskin/special-pancake/internal/features/listtodos"
	"github.com/jhoskin/special-pancake/internal/features/updatetodo"
	"github.com/jhoskin/special-pancake/internal/infrastructure/db"
	pb "github.com/jhoskin/special-pancake/proto/gen/todo/v1"
	"github.com/jhoskin/special-pancake/proto/gen/todo/v1/todov1connect"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server handles HTTP requests and manages todo operations
type Server struct {
	listtodos  *listtodos.Handler
	createtodo *createtodo.Handler
	updatetodo *updatetodo.Handler
	deletetodo *deletetodo.Handler
}

// NewServer creates a new Server instance with the given database
func NewServer(db *db.BoltDB) *Server {
	return &Server{
		listtodos:  listtodos.NewHandler(db),
		createtodo: createtodo.NewHandler(db),
		updatetodo: updatetodo.NewHandler(db),
		deletetodo: deletetodo.NewHandler(db),
	}
}

// ListTodos implements the TodoService interface
func (s *Server) ListTodos(ctx context.Context, req *connect.Request[pb.ListTodosRequest]) (*connect.Response[pb.ListTodosResponse], error) {
	return s.listtodos.Handle(ctx, req)
}

// CreateTodo implements the TodoService interface
func (s *Server) CreateTodo(ctx context.Context, req *connect.Request[pb.CreateTodoRequest]) (*connect.Response[pb.CreateTodoResponse], error) {
	return s.createtodo.Handle(ctx, req)
}

// UpdateTodo implements the TodoService interface
func (s *Server) UpdateTodo(ctx context.Context, req *connect.Request[pb.UpdateTodoRequest]) (*connect.Response[pb.UpdateTodoResponse], error) {
	return s.updatetodo.Handle(ctx, req)
}

// DeleteTodo implements the TodoService interface
func (s *Server) DeleteTodo(ctx context.Context, req *connect.Request[pb.DeleteTodoRequest]) (*connect.Response[pb.DeleteTodoResponse], error) {
	return s.deletetodo.Handle(ctx, req)
}

// Start begins listening for HTTP requests on the specified address
func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTodoServiceHandler(s)
	mux.Handle(path, handler)

	log.Printf("Starting server on %s...", addr)
	return http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
