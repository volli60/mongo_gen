// Package mongoDB provides a simple and convenient way to interact with MongoDB databases in Go.
// It includes common operations like creating connections, saving, updating, finding, and deleting documents.
package mongoDB

// Document represents an interface for MongoDB documents.
// Any struct that implements this interface can be used with the database operations.
// type Document interface {
//     // GetID returns the document's ObjectID
//     GetID() primitive.ObjectID
// }

// MongoHandler represents a MongoDB connection handler.
// It contains a reference to the MongoDB database.
// type MongoHandler struct {
//     DB *mongo.Database
// }

// NewMongoHandler creates a new MongoDB connection handler.
// It takes database name and MongoDB URL as parameters.
// Returns a handler and error if connection fails.
// func NewMongoHandler(dbName, url string) (*MongoHandler, error)

// CreateIndex creates an index in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - model: index model in BSON format
// func CreateIndex(db *mongo.Database, collectionName string, model bson.D) error

// SaveOne saves a single document to the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - doc: document to save
// Returns InsertOneResult and error if operation fails.
// func SaveOne[T Document](db *mongo.Database, collectionName string, doc T) (*mongo.InsertOneResult, error)

// SaveMany saves multiple documents to the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - docs: slice of documents to save
// Returns InsertManyResult and error if operation fails.
// func SaveMany[T Document](db *mongo.Database, collectionName string, docs []T) (*mongo.InsertManyResult, error)

// UpdateOne updates a single document in the specified collection.
// The document is identified by its ID.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - doc: document with updated fields
// Returns UpdateResult and error if operation fails.
// func UpdateOne[T Document](db *mongo.Database, collectionName string, doc T) (*mongo.UpdateResult, error)

// FindOne finds a single document in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - filter: query filter in BSON format
// Returns found document and error if operation fails.
// func FindOne[T Document](db *mongo.Database, collectionName string, filter bson.D) (*T, error)

// DeleteOne deletes a single document by its ID.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - id: ObjectID of the document to delete
// Returns DeleteResult and error if operation fails.
// func DeleteOne[T Document](db *mongo.Database, collectionName string, id primitive.ObjectID) (*mongo.DeleteResult, error)

// Find finds multiple documents in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - sortModel: sorting criteria in BSON format
//   - filter: query filter in BSON format
//   - skip: number of documents to skip
//   - limit: maximum number of documents to return
// Returns slice of found documents and error if operation fails.
// func Find[T Document](db *mongo.Database, collectionName string, sortModel bson.D, filter bson.D, skip int64, limit int64) (*[]T, error)

// Close closes the database connection.
// Should be called when the handler is no longer needed.
// Returns error if disconnection fails.
// func (h *MongoHandler) Close() error
