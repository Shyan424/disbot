package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type BackMessageEntity struct {
	Id    primitive.ObjectID `bson:"_id"`
	Key   string
	Value []string
}
