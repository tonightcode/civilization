package entity

type Person struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  LocalTime
	UpdatedAt  LocalTime
	DeletedAt  *LocalTime `sql:"index"`
	Name       string
	Desc       string
	Live_start LocalTime
	Live_end   LocalTime
}

func (person Person) TableName() string {
	return "person"
}
