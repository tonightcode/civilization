package entity

type Event struct {
	ID         uint `gorm:"primary_key"`
	Title      string
	Content    string
	HappenedAt LocalTime
	CreatedAt  LocalTime
	UpdatedAt  LocalTime
	DeletedAt  *LocalTime `sql:"index"`
	Persons    []Person   `gorm:"many2many:event_persons"`
}

func (event Event) TableName() string {
	return "event"
}
