package updatetodo

import (
	"context"
	"testing"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/features/createtodo"
	"todo-app/internal/infrastructure/db"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTodoHandler_Integration(t *testing.T) {
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	createHandler := createtodo.NewHandler(testDB)
	updateHandler := NewHandler(testDB)

	// First create a todo
	createReq := connect.NewRequest(&todov1.CreateTodoRequest{
		Title:       "Original Title",
		Description: "Original Description",
		Completed:   false,
	})

	createResp, err := createHandler.Handle(context.Background(), createReq)
	require.NoError(t, err)
	require.NotNil(t, createResp)
	require.NotNil(t, createResp.Msg.Todo)

	originalTodo := createResp.Msg.Todo

	t.Run("should update an existing todo", func(t *testing.T) {
		updateReq := connect.NewRequest(&todov1.UpdateTodoRequest{
			Id:          originalTodo.Id,
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   true,
		})

		resp, err := updateHandler.Handle(context.Background(), updateReq)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Msg.Todo)

		updatedTodo := resp.Msg.Todo
		assert.Equal(t, originalTodo.Id, updatedTodo.Id)
		assert.Equal(t, updateReq.Msg.Title, updatedTodo.Title)
		assert.Equal(t, updateReq.Msg.Description, updatedTodo.Description)
		assert.Equal(t, updateReq.Msg.Completed, updatedTodo.Completed)
		assert.NotEqual(t, originalTodo.UpdatedAt, updatedTodo.UpdatedAt)
	})
}
