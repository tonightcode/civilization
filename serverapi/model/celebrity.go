package model

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"culture/entity"
)

var DbConn *gorm.DB

type Celebrity struct {
}

func (celebrity Celebrity) Init() *gorm.DB {
	viper.SetConfigFile("config/db.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	viper.WatchConfig()
	fmt.Println(viper.GetString("mysql.local"))
	DbConn, err = gorm.Open("mysql", viper.GetString("mysql.local"))
	if err != nil {
		panic(err)
	}
	DbConn.AutoMigrate(&entity.Event{}, &entity.Person{})
	return DbConn
}

func (celebrity Celebrity) GetEvent() (event entity.Event, err error) {
	DbConn := celebrity.Init()
	if err = DbConn.Find(event).Error; err != nil {
		return event, err
	}
	return
}
