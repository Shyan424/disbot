package service

import (
	"discordbot/model"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BackMessageService interface {
	AddValue(string, string) bool
	Insert(key string, value string) bool
	GetRandomValue(key string) string
}

type BackMessageConnection struct {
	repo model.BackMessageRepository
}

func GetBackMessageService() BackMessageService {
	modelConn := model.GetBackMessageRepository()
	return &BackMessageConnection{modelConn}
}

// 用key新增value
// 有key將value新增到value feild
// 沒key就insert新的
func (conn *BackMessageConnection) AddValue(key string, value string) bool {
	_, err := conn.repo.FindByKeyAndUpdate(key, value)
	if err != nil && err != mongo.ErrNoDocuments {
		return false
	}

	return true
}

func (conn *BackMessageConnection) Insert(key string, value string) bool {
	values := []string{value}
	bm := model.BackMessage{Id: primitive.NewObjectID(), Key: key, Value: values}
	return conn.repo.Insert([]model.BackMessage{bm})
}

func (conn *BackMessageConnection) GetRandomValue(key string) string {
	value := ""
	bm := conn.repo.FindByKey(key)

	if bm != nil {
		values := bm.Value
		value = values[random(len(values))]
	}

	return value
}

func random(maxIndex int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxIndex)
}
