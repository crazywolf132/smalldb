package smalldb_test

import (
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/crazywolf132/smalldb"
)

// Define a test struct for our data type
type User struct {
	Name string
	Age  int
}

// Helper function to clean up test files
func cleanup(file string) {
	_ = os.Remove(file)
}

func TestOpen(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, err := smalldb.Open[User](file)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	if db == nil {
		t.Fatalf("Expected db instance, got nil")
	}
}

func TestSetAndGet(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, _ := smalldb.Open[User](file)
	user := User{Name: "Alice", Age: 30}

	err := db.Set("user:1", user)
	if err != nil {
		t.Fatalf("Failed to set data: %v", err)
	}

	retrievedUser, exists := db.Get("user:1")
	if !exists {
		t.Fatalf("Failed to get data: key does not exist")
	}

	if !reflect.DeepEqual(user, retrievedUser) {
		t.Fatalf("Expected %v, got %v", user, retrievedUser)
	}
}

func TestDelete(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, _ := smalldb.Open[User](file)
	user := User{Name: "Bob", Age: 25}

	_ = db.Set("user:2", user)
	err := db.Delete("user:2")
	if err != nil {
		t.Fatalf("Failed to delete data: %v", err)
	}

	_, exists := db.Get("user:2")
	if exists {
		t.Fatalf("Expected key to be deleted")
	}
}

func TestGetAll(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, _ := smalldb.Open[User](file)
	users := map[string]User{
		"user:1": {Name: "Alice", Age: 30},
		"user:2": {Name: "Bob", Age: 25},
	}

	for k, v := range users {
		_ = db.Set(k, v)
	}

	allData := db.GetAll()
	if len(allData) != len(users) {
		t.Fatalf("Expected %d items, got %d", len(users), len(allData))
	}

	for k, v := range users {
		if !reflect.DeepEqual(allData[k], v) {
			t.Fatalf("Expected %v for key %s, got %v", v, k, allData[k])
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, _ := smalldb.Open[User](file)
	user := User{Name: "Charlie", Age: 28}
	_ = db.Set("user:3", user)

	var wg sync.WaitGroup
	numReaders := 10

	// Start multiple readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, exists := db.Get("user:3")
			if !exists {
				t.Errorf("Reader failed to get data")
			}
		}()
	}

	// Start a writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		updatedUser := User{Name: "Charlie", Age: 29}
		err := db.Set("user:3", updatedUser)
		if err != nil {
			t.Errorf("Writer failed to set data: %v", err)
		}
	}()

	wg.Wait()
}

func TestTransaction(t *testing.T) {
	file := "test_db.json"
	defer cleanup(file)

	db, _ := smalldb.Open[User](file)

	err := db.Transaction(func(tx *smalldb.Tx[User]) error {
		tx.Set("user:4", User{Name: "Dave", Age: 40})
		tx.Set("user:5", User{Name: "Eve", Age: 35})
		tx.Delete("user:nonexistent")
		return nil // Commit the transaction
	})
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}

	// Verify data was committed
	user4, exists := db.Get("user:4")
	if !exists || user4.Name != "Dave" {
		t.Fatalf("Expected user:4 to be Dave")
	}

	user5, exists := db.Get("user:5")
	if !exists || user5.Name != "Eve" {
		t.Fatalf("Expected user:5 to be Eve")
	}
}
