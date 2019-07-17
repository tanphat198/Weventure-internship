package Struct

type Course struct {
	Id uint `gorm:"column:id; primary_key"`
	Name string
	Classes []Class	`gorm:"foreignkey:CourseID"`
}



