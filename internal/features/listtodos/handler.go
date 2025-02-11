package listtodos

import (
	"context"
	"encoding/json"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/infrastructure/db"
	"todo-app/models"

	"github.com/bufbuild/connect-go"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	db *db.BoltDB
}

func NewHandler(db *db.BoltDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Handle(
	ctx context.Context,
	req *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {
	todos, err := h.getAllTodos()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoTodos := make([]*todov1.Todo, len(todos))
	for i, todo := range todos {
		protoTodos[i] = &todov1.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		}
	}

	return connect.NewResponse(&todov1.ListTodosResponse{
		Todos: protoTodos,
	}), nil
}

func (h *Handler) getAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	err := h.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)
		return bucket.ForEach(func(k, v []byte) error {
			var todo models.Todo
			if err := json.Unmarshal(v, &todo); err != nil {
				return err
			}
			todos = append(todos, todo)
			return nil
		})
	})

	return todos, err
}
