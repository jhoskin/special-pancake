package listtodos

import (
	"context"
	"encoding/json"

	"github.com/jhoskin/special-pancake/internal/domain"
	"github.com/jhoskin/special-pancake/internal/infrastructure/db"
	pb "github.com/jhoskin/special-pancake/proto/gen/todo/v1"

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
	req *connect.Request[pb.ListTodosRequest],
) (*connect.Response[pb.ListTodosResponse], error) {
	todos, err := h.getAllTodos()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoTodos := make([]*pb.Todo, len(todos))
	for i, todo := range todos {
		protoTodos[i] = &pb.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		}
	}

	return connect.NewResponse(&pb.ListTodosResponse{
		Todos: protoTodos,
	}), nil
}

func (h *Handler) getAllTodos() ([]domain.Todo, error) {
	var todos []domain.Todo

	err := h.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)
		return bucket.ForEach(func(k, v []byte) error {
			var todo domain.Todo
			if err := json.Unmarshal(v, &todo); err != nil {
				return err
			}
			todos = append(todos, todo)
			return nil
		})
	})

	return todos, err
}
