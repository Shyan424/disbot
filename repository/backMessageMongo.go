package repository

import (
	"context"
	"discordbot/model/entity"
	"discordbot/model/vo"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BackMessageMongoConnection struct {
	collection *mongo.Collection
}

func bind(serviceVo vo.BackMessageVo) entity.BackMessageEntity {
	objId := primitive.ObjectID([12]byte{})

	if serviceVo.Id != "" {
		id, err := primitive.ObjectIDFromHex(serviceVo.Id)
		if err == nil {
			objId = id
		}
	}

	return entity.BackMessageEntity{Id: objId, Key: serviceVo.Key, Value: serviceVo.Value}
}

func unBind(entity entity.BackMessageEntity) vo.BackMessageVo {
	return vo.BackMessageVo{Id: entity.Id.Hex(), Key: entity.Key, Value: entity.Value}
}

func (c *BackMessageMongoConnection) Insert(serviceBo []vo.BackMessageVo) bool {
	insertObject := []interface{}{}
	for _, v := range serviceBo {
		bm := bind(v)
		if bm.Id.IsZero() {
			bm.Id = primitive.NewObjectID()
		}
		insertObject = append(insertObject, bm)
	}

	_, err := c.collection.InsertMany(context.TODO(), insertObject)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (c *BackMessageMongoConnection) UpdateByKey(key string, values ...string) bool {
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

func (c *BackMessageMongoConnection) FindByKey(key string) *vo.BackMessageVo {
	filter := bson.D{{Key: "key", Value: key}}
	var message entity.BackMessageEntity
	result := c.collection.FindOne(context.TODO(), filter)
	err := result.Decode(&message)
	if err != nil {
		return nil
	}

	back := unBind(message)

	return &back
}

func (c *BackMessageMongoConnection) FindByKeyAndUpdate(key string, values ...string) (*vo.BackMessageVo, error) {
	filter := bson.D{{Key: "key", Value: key}}
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "value", Value: bson.D{
				{Key: "$each", Value: values},
			}},
		}},
	}
	option := options.FindOneAndUpdate().SetUpsert(true)

	var backMessage entity.BackMessageEntity
	// 有會直接update沒有會直接insert 回傳的result是還沒update前的
	err := c.collection.FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&backMessage)
	back := unBind(backMessage)

	return &back, err
}
