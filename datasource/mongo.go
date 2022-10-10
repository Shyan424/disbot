package datasource

import (
	"context"
	"log"
	"strings"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConnection *mongo.Database
	mongoClient     *mongo.Client
)

// 連線Mongodb
func ConnectMongo() {
	uri := viper.GetString("datasource.mongodb.uri")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	logErr(err)

	logErr(client.Ping(context.TODO(), nil))
	mongoClient = client

	db := uri[strings.LastIndex(uri, "/")+1:]
	mongoConnection = client.Database(db)
}

func GetCollection(collection string) *mongo.Collection {
	return mongoConnection.Collection(collection)
}

// 關閉Mongodb connection
func CloseMongo() {
	logErr(mongoClient.Disconnect(context.TODO()))
}

func logErr(err error) bool {
	isErr := err != nil
	if isErr {
		log.Fatal(err)
	}

	return isErr
}
