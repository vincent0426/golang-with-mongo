package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// should be same with the config in docker-compose
var credential = options.Credential{
	Username: "root",
	Password: "123",
}

var usersCollection *mongo.Collection
var client *mongo.Client
var err error

func InitDB() {
	usersCollection = client.Database("testing").Collection("users")
}

func CreateUser(user bson.M) *mongo.InsertOneResult {
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors
	if err != nil {
		panic(err)
	}
	return result
}

func CreateUsers(users []interface{}) *mongo.InsertManyResult {
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	return results
}

func FindUser(user bson.M) []byte {
	var findResult bson.M
	err = usersCollection.FindOne(context.TODO(), user).Decode(&findResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return []byte{}
		}
		panic(err)
	}
	findResultJson, err := json.Marshal(findResult)
	if err != nil {
		panic(err)
	}

	return findResultJson
}

func FindAll() []bson.M {
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func UpdateUser(updateId string, updateAge int) *mongo.UpdateResult {
	id, _ := primitive.ObjectIDFromHex("6272ae519925a594f5671f6b")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"age": updateAge}}
	// var result bson.D
	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}
	return result
}

func deleteUser(deleteName string) *mongo.DeleteResult {
	filter := bson.M{"fullName": deleteName}
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	InitDB()

	// Create One
	// user := bson.M{"fullName": "Vincent", "age": 30}
	// result := CreateUser(user)
	// fmt.Println(result.InsertedID)

	// Create Many
	// users := []interface{}{
	// 	bson.M{"fullName": "Janice", "age": 25},
	// 	bson.M{"fullName": "Bob", "age": 20},
	// 	bson.M{"fullName": "Alan", "age": 28},
	// }
	// results := CreateUsers(users)
	// fmt.Println(results.InsertedIDs)

	// Find One
	// findOneuser := bson.M{"fullName": "Vincent"}
	// findOneresult := FindUser(findOneuser)
	// fmt.Printf("%s", findOneresult)

	// Find all
	// findAllResults := FindAll()

	// fmt.Println("displaying all results in a collection")
	// for _, result := range findAllResults {
	// 	fmt.Println(result)
	// }

	// Update
	// get an id from database
	// updateId := "6272acb5f1776b8286f11d50"
	// updateAge := 100
	// updateResult := UpdateUser(updateId, updateAge)

	// fmt.Printf("%v", updateResult.MatchedCount)

	// Delete
	// deleteName := "Vincent"
	// deleteResult := deleteUser(deleteName)

	// fmt.Printf("Delete result: %v", deleteResult)
}
