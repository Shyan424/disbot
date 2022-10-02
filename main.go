package main

import (
	"discordbot/app"
)

func main() {
	app.Run()
}

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://discordbot:botbotbot@localhost:27017/discordbot"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		if err := client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	collection := client.Database("discordbot").Collection("backmessage")

// 	var doc Backmessage
// 	doc.Key = "test"
// 	doc.Value = append(doc.Value, "test")
// 	fmt.Println(doc)

// 	result, err := collection.InsertOne(context.TODO(), doc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result)
// }

// type Backmessage struct {
// 	_id   string
// 	Key   string
// 	Value []string
// }
