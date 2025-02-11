package deletetodo

import (
	"context"
	"testing"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/features/createtodo"
	"todo-app/internal/features/listtodos"
	"todo-app/internal/infrastructure/db"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteTodoHandler_Integration(t *testing.T) {
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	createHandler := createtodo.NewHandler(testDB)
	deleteHandler := NewHandler(testDB)
	listHandler := listtodos.NewHandler(testDB)

	// First create a todo
	createReq := connect.NewRequest(&todov1.CreateTodoRequest{
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
	})

	createResp, err := createHandler.Handle(context.Background(), createReq)
	require.NoError(t, err)
	require.NotNil(t, createResp)
	require.NotNil(t, createResp.Msg.Todo)

	todoID := createResp.Msg.Todo.Id

	t.Run("should delete an existing todo", func(t *testing.T) {
		deleteReq := connect.NewRequest(&todov1.DeleteTodoRequest{
			Id: todoID,
		})

		resp, err := deleteHandler.Handle(context.Background(), deleteReq)

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.True(t, resp.Msg.Success)

		// Verify todo was deleted by listing todos
		listReq := connect.NewRequest(&todov1.ListTodosRequest{})
		listResp, err := listHandler.Handle(context.Background(), listReq)
		require.NoError(t, err)
		assert.Empty(t, listResp.Msg.Todos)
	})
}
