package entity

import "time"

type Event struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
	Persons   []Person
}
