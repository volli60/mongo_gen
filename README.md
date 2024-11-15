Package mongoDB provides a simple and convenient way to interact with MongoDB databases in Go.
It includes common operations like creating connections, saving, updating, finding, and deleting documents.

Basic usage:
```go
// Create a connection
    handler, err := mongoDB.NewMongoHandler("your_db", "mongodb://localhost:27017")
    if err != nil {
        log.Fatal(err)
    }
    defer handler.Close()

// Save a document
    user := User{Name: "John"}
    result, err := mongoDB.SaveOne(handler.DB, "users", user)

// Find a document
    filter := bson.D{{Key: "name", Value: "John"}}
    foundUser, err := mongoDB.FindOne[User](handler.DB, "users", filter)
```
