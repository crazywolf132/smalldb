package smalldb

import (
	"os"
	"path/filepath"
	"sync"
)

// DB represents the small database instance.
// T is the type of values stored in the database.
type DB[T any] struct {
	filepath string
	mu       sync.RWMutex
	data     map[string]T
}

// Open initializes the database at the given file path.
// It creates the file and necessary directories if they don't exist.
func Open[T any](fp string) (*DB[T], error) {
	err := os.MkdirAll(filepath.Dir(fp), 0755)
	if err != nil {
		return nil, err
	}

	data, err := readData[T](fp)
	if err != nil {
		return nil, err
	}

	return &DB[T]{
		filepath: fp,
		data:     data,
	}, nil
}

// Get retrieves the value associated with the given key.
// Returns the value and a boolean indicating whether the key exists.
func (db *DB[T]) Get(key string) (T, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, exists := db.data[key]
	return value, exists
}

// Set sets the value for the given key.
// This operation is thread-safe.
func (db *DB[T]) Set(key string, value T) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = value
	return db.persist()
}

// Delete removes the value associated with the given key.
// This operation is thread-safe.
func (db *DB[T]) Delete(key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, key)
	return db.persist()
}

// GetAll returns a copy of all key-value pairs in the database.
func (db *DB[T]) GetAll() map[string]T {
	db.mu.RLock()
	defer db.mu.RUnlock()

	dataCopy := make(map[string]T, len(db.data))
	for k, v := range db.data {
		dataCopy[k] = v
	}
	return dataCopy
}

// Transaction provides a function to execute multiple operations atomically.
// The provided function fn is executed with exclusive access to the database.
func (db *DB[T]) Transaction(fn func(tx *Tx[T]) error) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	tx := &Tx[T]{
		db:   db,
		data: cloneMap(db.data),
	}

	if err := fn(tx); err != nil {
		return err
	}

	// Commit changes
	db.data = tx.data
	return db.persist()
}

// persist writes the in-memory data to the JSON file.
func (db *DB[T]) persist() error {
	return writeData(db.filepath, db.data)
}
