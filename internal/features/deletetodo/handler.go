package deletetodo

import (
	"context"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/common/db"

	"github.com/bufbuild/connect-go"
)

type Handler struct {
	db *db.BoltDB
}

func NewHandler(db *db.BoltDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Handle(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTodoRequest],
) (*connect.Response[todov1.DeleteTodoResponse], error) {
	if err := h.db.DeleteTodo(uint(req.Msg.Id)); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&todov1.DeleteTodoResponse{
		Success: true,
	}), nil
}
