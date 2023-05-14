package model

import (
	"log"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"culture/entity"
)

type Person struct {
}

func (person Person) Init() *gorm.DB {
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
	return db
}

func (person Person) GetPerson(id int) (personEntity entity.Person, err error) {
	db := person.Init()
	if err = db.Preload("Persons").Where("id=?", id).Find(&personEntity).Error; err != nil {
		return
	}
	return
}

func (person Person) GetPersons() (personEntity []*entity.Person, err error) {
	db := person.Init()
	if err = db.Preload("Persons").Find(&personEntity).Error; err != nil {
		return
	}
	return
}

func (person Person) EditPerson(personEntity entity.Person) int {
	db := person.Init()
	db.Create(&personEntity)
	return 1
}
