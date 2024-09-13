package smalldb

// Tx represents a transaction with exclusive access to the database.
type Tx[T any] struct {
	db   *DB[T]
	data map[string]T
}

// Get retrieves the value associated with the given key within the transaction.
func (tx *Tx[T]) Get(key string) (T, bool) {
	value, exists := tx.data[key]
	return value, exists
}

// Set sets the value for the given key within the transaction.
func (tx *Tx[T]) Set(key string, value T) {
	tx.data[key] = value
}

// Delete removes the value associated with the given key within the transaction.
func (tx *Tx[T]) Delete(key string) {
	delete(tx.data, key)
}
