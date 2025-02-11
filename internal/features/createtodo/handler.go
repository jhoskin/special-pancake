package createtodo

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
	req *connect.Request[pb.CreateTodoRequest],
) (*connect.Response[pb.CreateTodoResponse], error) {
	now := time.Now()
	todo := domain.Todo{
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		Completed:   req.Msg.Completed,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := h.addTodo(&todo); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.CreateTodoResponse{
		Todo: &pb.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		},
	}), nil
}

func (h *Handler) addTodo(todo *domain.Todo) error {
	return h.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)

		id, _ := bucket.NextSequence()
		todo.ID = uint(id)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return bucket.Put(db.Itob(todo.ID), buf)
	})
}
