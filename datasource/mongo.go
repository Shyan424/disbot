package datasource

import (
	"context"
	"log"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var self *MongoDatasource
var lock = sync.Mutex{}

type MongoDatasource struct {
	Uri             string
	mongoConnection *mongo.Database
	mongoClient     *mongo.Client
}

func GetDatasource() *MongoDatasource {

	if self == nil {
		lock.Lock()
		if self == nil {
			self = new()
		}
		lock.Unlock()
	}

	return self
}

func new() *MongoDatasource {
	return &MongoDatasource{}
}

// 連線Mongodb
func (d *MongoDatasource) ConnectMongo() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.Uri))
	logErr(err)

	logErr(client.Ping(context.TODO(), nil))
	d.mongoClient = client

	db := d.Uri[strings.LastIndex(d.Uri, "/")+1:]
	d.mongoConnection = client.Database(db)
}

func GetCollection(collection string) *mongo.Collection {
	return self.mongoConnection.Collection(collection)
}

// 關閉Mongodb connection
func (d *MongoDatasource) CloseMongo() {
	logErr(d.mongoClient.Disconnect(context.TODO()))
}

func logErr(err error) bool {
	isErr := err != nil
	if isErr {
		log.Fatal(err)
	}

	return isErr
}
