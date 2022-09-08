package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Olexander753/REST_API_WebServer/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

func NewStorage(client *mongo.Client, database, collection string) user.Storage {
	return &db{
		collection: client.Database(database).Collection(collection),
	}
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user error: %v", err)
	}

	log.Println("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	log.Println(user)
	return "", fmt.Errorf("failed to convert objectID to HEX, oid: %s", oid)
}

func (d *db) FindByID(ctx context.Context, id string) (u user.User, err error) {
	log.Println("find user by id=", id)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert HEX to objectID, HEX: %s, error: %v", id, err)
	}

	//mongo.getDatabase("test").getCollection("docs").find({})
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("not found")
		}
		return u, fmt.Errorf("failed to FindOne user by id: %s, error: %v", id, result.Err())
	}

	log.Println("decode finde user")
	err = result.Decode(&u)
	if err != nil {
		return u, fmt.Errorf("failed to Decode user by id: %s, error: %v", id, err)
	}

	return u, nil
}

func (d *db) FindAll(ctx context.Context) (u []user.User, err error) {
	log.Println("find users")

	result, err := d.collection.Find(ctx, bson.M{})
	if result.Err() != nil {
		return u, fmt.Errorf("failed to FindAll users, error: %v", result.Err())
	}

	log.Println("decode finde users")
	err = result.All(ctx, &u)
	if err != nil {
		return u, fmt.Errorf("failed to Decode users, error: %v", err)
	}

	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	log.Println("update user, id=", user.ID)
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert userID to ObjectID, ID=%s, error: %v", user.ID, err)
	}

	filter := bson.M{"_id": objectID}

	log.Println("marshal user, id=", user.ID)
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user, ID=%s, error^ %v", user.ID, err)
	}

	log.Println("unmarshal userBytes")
	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal userBytes, error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{"$set": updateUserObj}

	log.Println("execute update user")
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query, error: %v", err)
	}

	if result.MatchedCount == 0 {
		//TODO ErrorNotFound
		return fmt.Errorf("user not found")
	}

	log.Printf("Matched %d documents and Modified %d documents\n", result.MatchedCount, result.ModifiedCount)

	return nil

}

func (d *db) Delete(ctx context.Context, id string) error {
	log.Println("delete user, id=", id)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert userID to ObjectID, ID=%s, error: %v", id, err)
	}

	filter := bson.M{"_id": objectID}

	log.Println("execute delete user query")
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute delete user query, error: %v", err)
	}

	if result.DeletedCount == 0 {
		//TODO ErrorNotFound
		return fmt.Errorf("user not found")
	}

	log.Printf("Deleted %d documents\n", result.DeletedCount)

	return nil
}
