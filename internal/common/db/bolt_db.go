package db

import (
	"encoding/json"
	"fmt"
	"time"

	"todo-app/models"

	"go.etcd.io/bbolt"
)

var todoBucket = []byte("todos")

type BoltDB struct {
	db *bbolt.DB
}

func NewBoltDB(path string) (*BoltDB, error) {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	// Create the bucket if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todoBucket)
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

func (b *BoltDB) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		return bucket.ForEach(func(k, v []byte) error {
			var todo models.Todo
			if err := json.Unmarshal(v, &todo); err != nil {
				return err
			}
			todos = append(todos, todo)
			return nil
		})
	})

	return todos, err
}

func (b *BoltDB) AddTodo(todo models.Todo) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(todoBucket)

		// Auto-increment ID
		id, _ := bucket.NextSequence()
		todo.ID = uint(id)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return bucket.Put(itob(todo.ID), buf)
	})
}

func (b *BoltDB) UpdateTodo(todo models.Todo) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(todoBucket)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return bucket.Put(itob(todo.ID), buf)
	})
}

func (b *BoltDB) DeleteTodo(id uint) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		return bucket.Delete(itob(id))
	})
}

// itob converts an ID to a byte slice
func itob(v uint) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		b[i] = byte(v >> (i * 8))
	}
	return b
}
