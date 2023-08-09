package models

import (
	"context"
	"log"
	"time"

	"github.com/amosehiguese/mogo/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = store.GetCollection("user")

type User struct {
	Id 				primitive.ObjectID			`json:"id" bson:"_id,omitempty"`
	Name			string						`json:"name" bson:"name"`
	Gender			string						`json:"gender" bson:"gender"`
	Age				int							`json:"age" bson:"age"`
}



func (u *User) InsertOne() (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, u)
	if err != nil {
		log.Println("Insert not successful ->", err)
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (u *User) UpdateOne(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": u}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
func DeleteOne(userId string) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	id, err:= primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": id}
	_, err = userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}


func RetrieveOne(userId string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var user User

	id, err:= primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	filter := bson.M{"_id": id}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var users  []User
	cur, err := userCollection.Find((ctx), bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user User
		if err = cur.Decode(&user); err != nil {
			log.Println("Unable to decode -> ", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil{
		log.Println(err)
		return nil, err
	}
	cur.Close(context.TODO())
	return users, nil
}
