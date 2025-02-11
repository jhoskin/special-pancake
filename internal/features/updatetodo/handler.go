package updatetodo

import (
	"context"
	"encoding/json"
	"time"

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
	req *connect.Request[todov1.UpdateTodoRequest],
) (*connect.Response[todov1.UpdateTodoResponse], error) {
	todo := models.Todo{
		ID:          uint(req.Msg.Id),
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		Completed:   req.Msg.Completed,
		UpdatedAt:   time.Now(),
	}

	if err := h.updateTodo(&todo); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&todov1.UpdateTodoResponse{
		Todo: &todov1.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		},
	}), nil
}

func (h *Handler) updateTodo(todo *models.Todo) error {
	return h.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return bucket.Put(db.Itob(todo.ID), buf)
	})
}
