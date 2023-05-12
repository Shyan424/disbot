package service

import (
	"discordbot/model/vo"
	"discordbot/repository"
	"math/rand"
	"time"
)

type BackMessageService interface {
	AddValue(string, string) bool
	Insert(key string, value string) bool
	GetRandomValue(key string) string
}

type BackMessageConnection struct {
	repo repository.BackMessageRepository
}

func GetBackMessageService() BackMessageService {
	modelConn := repository.GetBackMessageRepository()
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
	bm := vo.BackMessageVo{Key: key, Value: values}
	return conn.repo.Insert([]vo.BackMessageVo{bm})
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
