package app

import (
	"discordbot/bot"
	"discordbot/datasource"
	"log"

	"github.com/spf13/viper"
)

func Run() {
	loadConfig()
	mongo := getMongo()
	bot := getBot()

	mongo.ConnectMongo()
	// test()
	defer mongo.CloseMongo()
	bot.ConnectDiscord()
	defer bot.CloseDiscord()
}

func loadConfig() {
	log.SetFlags(log.Lshortfile)
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func getMongo() *datasource.MongoDatasource {
	mongo := datasource.GetDatasource()
	mongo.Uri = viper.GetString("datasource.mongodb.uri")

	return mongo
}

func getBot() *bot.Discordbot {
	return bot.New("Bot " + viper.GetString("discordbot.token"))
}

// func test() {
// 	s := service.GetBackMessageService()
// 	// testTimes := 1
// 	// g := sync.WaitGroup{}
// 	// g.Add(testTimes)

// 	// for i := 0; i < testTimes; i++ {
// 	// 	go func(i int) {
// 	// 		k := i
// 	// 		for j := 0; j < 1000; j++ {
// 	// 			key := "key8" + strconv.Itoa(k)
// 	// 			value := "value" + strconv.Itoa(j)

// 	// 			s.AddValue(key, value)
// 	// 		}
// 	// 		g.Done()
// 	// 	}(i)
// 	// }
// 	// g.Wait()
// 	back := s.AddValue("insert1", "value3")
// 	fmt.Println(back)
// }
