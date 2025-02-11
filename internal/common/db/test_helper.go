package db

import (
	"os"
	"path/filepath"
	"testing"
)

// NewTestDB creates a new BoltDB instance for testing
func NewTestDB(t *testing.T) (*BoltDB, func()) {
	t.Helper()

	// Create a temporary directory for the test database
	dir, err := os.MkdirTemp("", "todo-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	// Create a database file in the temporary directory
	dbPath := filepath.Join(dir, "test.db")
	db, err := NewBoltDB(dbPath)
	if err != nil {
		os.RemoveAll(dir)
		t.Fatalf("failed to create test database: %v", err)
	}

	// Return the database and a cleanup function
	cleanup := func() {
		db.Close()
		os.RemoveAll(dir)
	}

	return db, cleanup
}
