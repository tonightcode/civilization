package entity

import "time"

type Person struct {
	Id         int
	Name       int
	Desc       string
	live_start time.Time
	live_end   time.Time
	CreatedAt  time.Time
	UpdateAt   time.Time
}
