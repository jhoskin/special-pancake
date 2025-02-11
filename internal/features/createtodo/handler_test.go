package createtodo

import (
	"context"
	"testing"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/infrastructure/db"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTodoHandler_Integration(t *testing.T) {
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	handler := NewHandler(testDB)

	t.Run("should create a new todo", func(t *testing.T) {
		req := connect.NewRequest(&todov1.CreateTodoRequest{
			Title:       "Test Todo",
			Description: "Test Description",
			Completed:   false,
		})

		resp, err := handler.Handle(context.Background(), req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Msg.Todo)

		todo := resp.Msg.Todo
		assert.NotZero(t, todo.Id)
		assert.Equal(t, req.Msg.Title, todo.Title)
		assert.Equal(t, req.Msg.Description, todo.Description)
		assert.Equal(t, req.Msg.Completed, todo.Completed)
		assert.NotNil(t, todo.CreatedAt)
		assert.NotNil(t, todo.UpdatedAt)
	})
}
