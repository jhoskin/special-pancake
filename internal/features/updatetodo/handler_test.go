package updatetodo

import (
	"context"
	"testing"

	"github.com/jhoskin/special-pancake/internal/features/createtodo"
	"github.com/jhoskin/special-pancake/internal/infrastructure/db"
	pb "github.com/jhoskin/special-pancake/proto/gen/todo/v1"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTodoHandler_Integration(t *testing.T) {
	// Set up test database
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	// Create handlers
	createHandler := createtodo.NewHandler(testDB)
	updateHandler := NewHandler(testDB)

	// Create a test todo
	createReq := connect.NewRequest(&pb.CreateTodoRequest{
		Title:       "Original Title",
		Description: "Original Description",
		Completed:   false,
	})

	createResp, err := createHandler.Handle(context.Background(), createReq)
	require.NoError(t, err)
	require.NotNil(t, createResp)
	require.NotNil(t, createResp.Msg.Todo)

	todoID := createResp.Msg.Todo.Id

	// Test updating todo
	t.Run("should update todo successfully", func(t *testing.T) {
		updateReq := connect.NewRequest(&pb.UpdateTodoRequest{
			Id:          todoID,
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   true,
		})

		updateResp, err := updateHandler.Handle(context.Background(), updateReq)
		require.NoError(t, err)
		require.NotNil(t, updateResp)
		require.NotNil(t, updateResp.Msg.Todo)

		// Verify updated fields
		assert.Equal(t, "Updated Title", updateResp.Msg.Todo.Title)
		assert.Equal(t, "Updated Description", updateResp.Msg.Todo.Description)
		assert.True(t, updateResp.Msg.Todo.Completed)
		assert.NotNil(t, updateResp.Msg.Todo.UpdatedAt)
	})
}
