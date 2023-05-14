package model

import (
	"log"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"culture/entity"
)

type Event struct {
}

func (event Event) Init() *gorm.DB {
	viper.SetConfigFile("config/db.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	viper.WatchConfig()
	db, err := gorm.Open("mysql", viper.GetString("mysql.local"))
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&entity.Person{}, &entity.Event{})
	return db
}

func (event Event) GetEvent(id int) (eventEntity entity.Event, err error) {
	db := event.Init()
	if err = db.Preload("Persons").Where("id=?", id).Find(&eventEntity).Error; err != nil {
		return
	}
	return
}

func (event Event) GetEvents() (eventEntitys []*entity.Event, err error) {
	db := event.Init()
	if err = db.Preload("Persons").Find(&eventEntitys).Error; err != nil {
		return
	}
	return
}

func (event Event) EditEvent(eventEntity entity.Event) int {
	db := event.Init()
	db.Create(&eventEntity)
	return 1
}
