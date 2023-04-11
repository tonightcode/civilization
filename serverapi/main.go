package main

import (
	"log"
	. "zongheng/routes"

	"github.com/spf13/viper"
)

func main() {
	startWebServer(":8080")
}

func startWebServer(port string) {
	loadConfig()
	r := NewRouter()

	log.Println("Starting HTTP service at " + port)
	r.Run(port)
}

func loadConfig() {
	viper.SetConfigFile("config/db.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	viper.WatchConfig()
}
