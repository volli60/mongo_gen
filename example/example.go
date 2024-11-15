package main

import (
	"log"

	"github.com/volli60/mongo_gen/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func (u User) GetID() primitive.ObjectID {
	return u.ID
}

func main() {
	// Create connection
	handler, err := mongoDB.NewMongoHandler("test_db", "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	const collectionName = "users"

	// Create user
	user := User{Name: "John Doe"}

	// Save to database
	result, err := mongoDB.SaveOne(handler.DB, collectionName, user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", result)

	// Find user
	filter := bson.D{{Key: "name", Value: "John Doe"}}
	foundUser, err := mongoDB.FindOne[User](handler.DB, collectionName, filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", foundUser)

	// Get list of users
	sortModel := bson.D{{Key: "name", Value: 1}}
	users, err := mongoDB.Find[User](handler.DB, collectionName, sortModel, bson.D{}, 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", users)
}
