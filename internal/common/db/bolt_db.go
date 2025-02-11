package db

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/bbolt"
)

var TodoBucket = []byte("todos")

type BoltDB struct {
	db *bbolt.DB
}

func NewBoltDB(path string) (*BoltDB, error) {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("could not create database directory: %v", err)
	}

	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	// Create the bucket if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(TodoBucket)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("could not create bucket: %v", err)
	}

	return &BoltDB{db: db}, nil
}

func (b *BoltDB) Close() error {
	return b.db.Close()
}

// View executes a read-only transaction
func (b *BoltDB) View(fn func(*bbolt.Tx) error) error {
	return b.db.View(fn)
}

// Update executes a writable transaction
func (b *BoltDB) Update(fn func(*bbolt.Tx) error) error {
	return b.db.Update(fn)
}

// Begin starts a new transaction
func (b *BoltDB) Begin(writable bool) (*bbolt.Tx, error) {
	return b.db.Begin(writable)
}
