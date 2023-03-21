package service

import (
	"discordbot/model"
	"math/rand"
	"time"
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

	return err == nil
}

func (conn *BackMessageConnection) Insert(key string, value string) bool {
	values := []string{value}
	bm := model.BackMessage{Key: key, Value: values}
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
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	return r.Intn(maxIndex)
}
