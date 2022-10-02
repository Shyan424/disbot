package model

import (
	"context"
	"discordbot/datasource"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BackMessage struct {
	MongoBase
	Key   string
	Value []string
}

type backMessageConnection struct {
	collection *mongo.Collection
}

func GetBackMessageConnection() *backMessageConnection {
	var backMessage backMessageConnection
	backMessage.collection = datasource.GtCollection("backmessage")

	return &backMessage
}

func (c *backMessageConnection) FindByKey(key string) *BackMessage {
	filter := bson.D{{Key: "key", Value: key}}
	var message BackMessage
	result := c.collection.FindOne(context.TODO(), filter)
	err := result.Decode(&message)
	if err != nil {
		log.Fatal(err)
	}

	return &message
}
