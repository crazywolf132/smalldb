# smalldb 🚀

**Your Tiny, Type-Safe, JSON-Powered Sidekick for Go!**

---

Are you tired of overcomplicating your Go applications with heavyweight databases for simple data storage? Do you yearn for a lightweight, easy-to-use, and fun solution to handle your data needs? Look no further! Introducing **`smalldb`**, your new best friend for simple, efficient, and type-safe data storage.

---

## 🌟 What is smalldb?

`smalldb` is a **lightweight**, **generic**, **file-based** key-value store for Go applications. It leverages the power of **Go's generics** to provide **type-safe**, **thread-safe**, and **easy-to-use** data storage without the overhead of a full-fledged database system.

---

## 🎯 Why smalldb?

In a world dominated by complex databases, sometimes all you need is a simple, reliable place to store your data. Maybe you're building a CLI tool, a small web service, or a prototype app. You don't need the hassle of setting up and maintaining a database server. You need something small. You need `smalldb`.

### **Real-World Problems It Solves:**

- **Simplicity**: No more wrestling with database drivers or ORM configurations.
- **Portability**: Data stored in human-readable JSON files—easy to debug and transfer.
- **Performance**: Minimal overhead for applications that require quick data access without concurrency bottlenecks.
- **Type Safety**: Compile-time checks prevent bugs, making your code safer and more reliable.
- **Rapid Development**: Perfect for prototypes, demos, or small projects where setup time is crucial.

---

## 🚀 Features

- **🔒 Type Safety with Generics**: Enjoy the peace of mind that comes with compile-time type checking.
- **🧘 Simplified API**: Intuitive methods for common operations—no ceremony, just results.
- **⚡ Thread-Safe**: Concurrency handled gracefully with support for multiple readers.
- **💾 Atomic Transactions**: Perform multiple operations atomically to maintain data integrity.
- **📄 Human-Readable Storage**: Data stored in JSON—easy to read, edit, and manage.
- **📦 Minimal Dependencies**: A lean library that won't bloat your project.
- **✨ Fun to Use**: Because coding should bring joy!

---

## 🛠️ Installation

Ready to embrace simplicity? Install `smalldb` with a single command:

```bash
go get github.com/crazywolf132/smalldb
```

---

## 🎓 Getting Started

### **Step 1: Import smalldb**

```go
import "github.com/crazywolf132/smalldb"
```

### **Step 2: Define Your Data Type**

Create a struct that represents the data you want to store.

```go
type User struct {
    Name string
    Age  int
}
```

### **Step 3: Open the Database**

Initialize your `smalldb` instance, specifying your data type.

```go
db, err := smalldb.Open[User]("path/to/db.json")
if err != nil {
    log.Fatal(err)
}
```

---

## 📖 Usage Guide

### **Storing Data Made Simple**

```go
// Create a new user.
alice := User{Name: "Alice", Age: 30}

// Store the user in the database.
if err := db.Set("user:1001", alice); err != nil {
    log.Fatal(err)
}
```

### **Retrieving Your Data**

```go
// Retrieve the user from the database.
user, exists := db.Get("user:1001")
if !exists {
    fmt.Println("User not found!")
} else {
    fmt.Printf("Hello, %s! You are %d years old.\n", user.Name, user.Age)
}
```

### **Updating Data**

```go
// Update the user's age.
user.Age = 31
if err := db.Set("user:1001", user); err != nil {
    log.Fatal(err)
}
```

### **Deleting Data**

```go
// Remove the user from the database.
if err := db.Delete("user:1001"); err != nil {
    log.Fatal(err)
}
```

### **Atomic Transactions**

Need to perform multiple operations atomically? We've got you covered!

```go
err := db.Transaction(func(tx *smalldb.Tx[User]) error {
    // Create two new users.
    tx.Set("user:1002", User{Name: "Bob", Age: 25})
    tx.Set("user:1003", User{Name: "Charlie", Age: 28})

    // Delete an existing user.
    tx.Delete("user:1004")

    // Commit the transaction.
    return nil
})
if err != nil {
    log.Fatal(err)
}
```

---

## 🌐 Real-World Applications

- **Configuration Storage**: Store and manage application settings effortlessly.
- **Caching Layer**: Implement a simple caching mechanism for your app.
- **Session Management**: Keep track of user sessions in a lightweight manner.
- **Prototype Development**: Quickly iterate on ideas without database setup overhead.
- **Education**: Perfect for learning and teaching Go's concurrency and generics.

---

## 💡 Tips and Tricks

- **Custom Types**: Use any serializable type—structs, slices, maps, you name it!
- **Data Inspection**: Since data is stored in JSON, you can easily inspect and edit it with any text editor.
- **Concurrency Control**: Read-heavy applications benefit from concurrent reads—thanks to `sync.RWMutex`.
- **Error Handling**: Always check for errors to handle unexpected situations gracefully.
- **Backups**: Copy the JSON file for a quick backup of your data.

---

## 🤔 FAQs

**Q: Is `smalldb` suitable for production use?**

A: Absolutely! While it's designed for simplicity, `smalldb` is thread-safe and reliable for applications that fit its use case.

**Q: Can I store complex nested data structures?**

A: Yes! As long as your data types are serializable to JSON, you can store them in `smalldb`.

**Q: What happens if multiple goroutines try to write at the same time?**

A: Write operations are synchronized using a mutex to prevent data races, ensuring data integrity.

**Q: How large can the database get?**

A: Since `smalldb` loads the entire database into memory, it's best suited for smaller datasets. If you're dealing with gigabytes of data, consider a different solution.

**Q: Can I use `smalldb` in a web application?**

A: Yes! It's perfect for small web services, APIs, or microservices where a full database might be overkill.

---

## 🛡️ Safety First!

`smalldb` handles concurrency and data integrity with care, but remember:

- **Backup Regularly**: Keep copies of your data, especially before major changes.
- **Validate Your Data**: Ensure the data you're storing is correct and sanitized.
- **Handle Errors**: Don't ignore errors—handle them appropriately to prevent surprises.

---

## 🚧 Under the Hood

Curious about how `smalldb` works its magic? Here's a peek:

- **Generics**: Utilizes Go's generics to enforce type safety at compile time.
- **Mutex Locks**: Manages concurrent access with `sync.RWMutex`, allowing multiple readers and single writers.
- **JSON Storage**: Serializes data to JSON for easy storage and retrieval.
- **Atomic Transactions**: Provides a transactional interface to perform multiple operations atomically.

---

## 🌈 Contributing

We welcome contributions! Whether it's fixing bugs, adding features, or improving documentation, your help is appreciated.

1. **Fork the Repository**
2. **Create a Feature Branch**
3. **Commit Your Changes**
4. **Push to Your Fork**
5. **Submit a Pull Request**

---

## 📜 License

`smalldb` is licensed under the **MIT License**—because sharing is caring!

---

## 🙌 Acknowledgments

A big thank you to the Go community for making such an amazing language, and to all the developers who inspire simplicity and elegance in code.

---

## ⭐ Star Us!

If you find `smalldb` useful, give us a star on GitHub! It helps others discover the project and motivates us to keep improving.

---

## 🚀 Let's Keep It Small and Simple!

In a world full of complexity, `smalldb` is here to remind us that sometimes, less is more. So go ahead—embrace simplicity, write cleaner code, and make your data storage woes a thing of the past!

---

Happy Coding! 🎉