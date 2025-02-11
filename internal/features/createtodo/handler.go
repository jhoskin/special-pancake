package createtodo

import (
	"context"
	"time"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/common/db"
	"todo-app/models"

	"github.com/bufbuild/connect-go"
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
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {
	now := time.Now()
	todo := models.Todo{
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		Completed:   req.Msg.Completed,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := h.db.AddTodo(todo); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&todov1.CreateTodoResponse{
		Todo: &todov1.Todo{
			Id:          uint32(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		},
	}), nil
}
