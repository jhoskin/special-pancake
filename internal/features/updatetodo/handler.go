package updatetodo

import (
	"context"
	"encoding/json"
	"time"

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
	req *connect.Request[pb.UpdateTodoRequest],
) (*connect.Response[pb.UpdateTodoResponse], error) {
	todo := domain.Todo{
		ID:          uint(req.Msg.Id),
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		Completed:   req.Msg.Completed,
		UpdatedAt:   time.Now(),
	}

	if err := h.updateTodo(&todo); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.UpdateTodoResponse{
		Todo: &pb.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		},
	}), nil
}

func (h *Handler) updateTodo(todo *domain.Todo) error {
	return h.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return bucket.Put(db.Itob(todo.ID), buf)
	})
}
