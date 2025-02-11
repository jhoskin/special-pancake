package createtodo

import (
	"context"
	"testing"

	"github.com/jhoskin/special-pancake/internal/infrastructure/db"
	pb "github.com/jhoskin/special-pancake/proto/gen/todo/v1"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTodoHandler_Integration(t *testing.T) {
	testDB, cleanup := db.NewTestDB(t)
	defer cleanup()

	handler := NewHandler(testDB)

	t.Run("should create todo successfully", func(t *testing.T) {
		req := connect.NewRequest(&pb.CreateTodoRequest{
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
		assert.Equal(t, "Test Todo", todo.Title)
		assert.Equal(t, "Test Description", todo.Description)
		assert.False(t, todo.Completed)
		assert.NotNil(t, todo.CreatedAt)
		assert.NotNil(t, todo.UpdatedAt)
	})
}
