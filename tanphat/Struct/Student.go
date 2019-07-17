package Struct

type Student struct {
	Id uint `gorm:"column:id; primary_key"`
	Name string
	Age int
	ClassID uint
}

