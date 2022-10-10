package model

import (
	"context"
	"discordbot/datasource"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BackMessage struct {
	Id    primitive.ObjectID `bson:"_id"`
	Key   string
	Value []string
}

type BackMessageConnection struct {
	collection *mongo.Collection
}

type BackMessageRepository interface {
	Insert(backMessages []BackMessage) bool
	UpdateByKey(key string, values ...string) bool
	FindByKey(key string) *BackMessage
	FindByKeyAndUpdate(key string, values ...string) (*BackMessage, error)
}

func GetBackMessageRepository() BackMessageRepository {
	backMessage := datasource.GetCollection("backmessage")

	return &BackMessageConnection{backMessage}
}

func (c *BackMessageConnection) Insert(in []BackMessage) bool {
	inin := []interface{}{}
	for _, v := range in {
		if v.Id.IsZero() {
			v.Id = primitive.NewObjectID()
		}
		inin = append(inin, v)
	}

	_, err := c.collection.InsertMany(context.TODO(), inin)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (c *BackMessageConnection) UpdateByKey(key string, values ...string) bool {
	filter := bson.D{{Key: "key", Value: key}}
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "value", Value: bson.D{
				{Key: "$each", Value: values},
			}},
		}}}

	_, err := c.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (c *BackMessageConnection) FindByKey(key string) *BackMessage {
	filter := bson.D{{Key: "key", Value: key}}
	var message BackMessage
	result := c.collection.FindOne(context.TODO(), filter)
	err := result.Decode(&message)
	if err != nil {
		log.Println(err)
	}

	return &message
}

func (c *BackMessageConnection) FindByKeyAndUpdate(key string, values ...string) (*BackMessage, error) {
	filter := bson.D{{Key: "key", Value: key}}
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "value", Value: bson.D{
				{Key: "$each", Value: values},
			}},
		}},
	}
	option := options.FindOneAndUpdate().SetUpsert(true)

	var backMessage BackMessage
	// 有會直接update沒有會直接insert 回傳的result是還沒update前的
	err := c.collection.FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&backMessage)

	return &backMessage, err
}
