package deletetodo

import (
	"context"

	"github.com/jhoskin/special-pancake/internal/infrastructure/db"
	pb "github.com/jhoskin/special-pancake/proto/gen/todo/v1"

	"github.com/bufbuild/connect-go"
	"go.etcd.io/bbolt"
)

type Handler struct {
	db *db.BoltDB
}

func NewHandler(db *db.BoltDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Handle(
	ctx context.Context,
	req *connect.Request[pb.DeleteTodoRequest],
) (*connect.Response[pb.DeleteTodoResponse], error) {
	if err := h.deleteTodo(uint(req.Msg.Id)); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteTodoResponse{
		Success: true,
	}), nil
}

func (h *Handler) deleteTodo(id uint) error {
	return h.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(db.TodoBucket)
		return bucket.Delete(db.Itob(id))
	})
}
