package mongoDB

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Document represents an interface for MongoDB documents.
// Any struct that implements this interface can be used with the database operations.
type Document interface {
	// GetID returns the document's ObjectID
	GetID() primitive.ObjectID
}

// MongoHandler represents a MongoDB connection handler.
// It contains a reference to the MongoDB database.
type MongoHandler struct {
	DB *mongo.Database
}

// NewMongoHandler creates a new MongoDB connection handler.
// It takes database name and MongoDB URL as parameters.
// Returns a handler and error if connection fails.
func NewMongoHandler(dbName, url string) (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoHandler{
		DB: client.Database(dbName),
	}, nil
}

// CreateIndex creates an index in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - model: index model in BSON format
func CreateIndex(db *mongo.Database, collectionName string, model bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)
	collection := db.Collection(collectionName)
	indexModel := mongo.IndexModel{Keys: model}
	_, err := collection.Indexes().CreateOne(ctx, indexModel, indexOpts)
	return err
}

// / SaveOne saves a single document to the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - doc: document to save
//
// Returns InsertOneResult and error if operation fails.
func SaveOne[T Document](db *mongo.Database, collectionName string, doc T) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	return collection.InsertOne(ctx, doc)
}

// SaveMany saves multiple documents to the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - docs: slice of documents to save
//
// Returns InsertManyResult and error if operation fails.
func SaveMany[T Document](db *mongo.Database, collectionName string, docs []T) (*mongo.InsertManyResult, error) {
	interfaces := make([]interface{}, len(docs))
	for i, doc := range docs {
		interfaces[i] = doc
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	return collection.InsertMany(ctx, interfaces)
}

// UpdateOne updates a single document in the specified collection.
// The document is identified by its ID.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - doc: document with updated fields
//
// Returns UpdateResult and error if operation fails.
func UpdateOne[T Document](db *mongo.Database, collectionName string, doc T) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	return collection.UpdateOne(ctx,
		bson.D{{Key: "_id", Value: doc.GetID()}},
		bson.D{{Key: "$set", Value: doc}})
}

// FindOne finds a single document in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - filter: query filter in BSON format
//
// Returns found document and error if operation fails.
func FindOne[T Document](db *mongo.Database, collectionName string, filter bson.D) (*T, error) {
	var result T
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteOne deletes a single document by its ID.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - id: ObjectID of the document to delete
//
// Returns DeleteResult and error if operation fails.
func DeleteOne[T Document](db *mongo.Database, collectionName string, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	opts := options.Delete().SetHint(bson.D{{Key: "_id", Value: 1}})
	return collection.DeleteOne(ctx, bson.M{"_id": id}, opts)
}

// Find finds multiple documents in the specified collection.
// Parameters:
//   - db: MongoDB database reference
//   - collectionName: name of the collection
//   - sortModel: sorting criteria in BSON format
//   - filter: query filter in BSON format
//   - skip: number of documents to skip
//   - limit: maximum number of documents to return
//
// Returns slice of found documents and error if operation fails.
func Find[T Document](
	db *mongo.Database,
	collectionName string,
	sortModel bson.D,
	filter bson.D,
	skip int64,
	limit int64,
) (*[]T, error) {
	var results []T
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(sortModel).SetSkip(skip).SetLimit(limit)
	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// Close closes the database connection.
// Should be called when the handler is no longer needed.
// Returns error if disconnection fails.
func (h *MongoHandler) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return h.DB.Client().Disconnect(ctx)
}
