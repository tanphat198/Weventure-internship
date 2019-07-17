package Struct


type Class struct {
	Id uint `gorm:"column:id; primary_key"`
	Name string
	Students []Student `gorm:"foreignkey:ClassID"`
	CourseID uint
}