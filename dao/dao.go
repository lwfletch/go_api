package dao

import (
	"context"
	"fmt"
	"log"

	"../models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "greyinc"

// COLLNAME Collection name
const COLLNAME = "users"

var db *mongo.Database

// InsertOneValue inserts one item from User model
func InsertOneValue(user models.User) {
	fmt.Println(user)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllUsers returns all people from DB
func GetAllUsers() []models.User {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.User
	var elem models.User
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

// DeleteUser deletes an existing user
func DeleteUser(user models.User) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), user, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateUser updates an existing user
func UpdateUser(user models.User, userId string) {
	fmt.Println("updateUser called")
}
