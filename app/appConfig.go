package app

import (
	"discordbot/datasource"
	"discordbot/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	datasource.ConnectMongo()

	c := model.GetBackMessageConnection()
	back := c.FindByKey("test")
	fmt.Println(back)

	defer datasource.CloseMongo()
	// discordbot.ConnectDiscord()
	// defer discordbot.CloseDiscord()
}
