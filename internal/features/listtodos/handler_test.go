package listtodos

import (
	"context"
	"testing"

	todov1 "todo-app/gen/proto/todo/v1"
	"todo-app/internal/common/db"
	"todo-app/internal/features/createtodo"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTodosHandler_Integration(t *testing.T) {
	// Set up test database
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	// Create handlers
	createHandler := createtodo.NewHandler(testDB)
	listHandler := NewHandler(testDB)

	// Create some test todos
	todos := []struct {
		title       string
		description string
		completed   bool
	}{
		{"Test Todo 1", "Description 1", false},
		{"Test Todo 2", "Description 2", true},
	}

	for _, todo := range todos {
		req := connect.NewRequest(&todov1.CreateTodoRequest{
			Title:       todo.title,
			Description: todo.description,
			Completed:   todo.completed,
		})

		resp, err := createHandler.Handle(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Msg.Todo)
	}

	// Test listing todos
	t.Run("should list all todos", func(t *testing.T) {
		req := connect.NewRequest(&todov1.ListTodosRequest{})
		resp, err := listHandler.Handle(context.Background(), req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Msg.Todos)
		assert.Len(t, resp.Msg.Todos, len(todos))

		// Verify todo contents
		for i, todo := range resp.Msg.Todos {
			assert.Equal(t, todos[i].title, todo.Title)
			assert.Equal(t, todos[i].description, todo.Description)
			assert.Equal(t, todos[i].completed, todo.Completed)
			assert.NotZero(t, todo.CreatedAt)
			assert.NotZero(t, todo.UpdatedAt)
		}
	})
}
