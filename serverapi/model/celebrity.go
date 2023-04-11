package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type event struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
	peoples   []people
}

type people struct {
	Id        int
	Name      int
	Desc      string
	start     time.Time
	end       time.Time
	CreatedAt time.Time
	UpdateAt  time.Time
}

var DbConn *gorm.DB

func init() {
	var err error
	DbConn, err = gorm.Open("mysql", "root:root@/test_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	DbConn.AutoMigrate()

}
