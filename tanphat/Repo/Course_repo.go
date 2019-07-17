package Repo

import (
	"github.com/jinzhu/gorm"
	g "tanphat/Struct"
)

type Course_repo struct {
	Db *gorm.DB
}

func (s *Course_repo) GetbyID(ID int) g.Course{
	var course g.Course
	s.Db.Set("gorm:auto_preload", true).
		Where("ID = ?" , ID).
		Find(&course)

	return course
}

func (s *Course_repo) GetbyName(Name string) g.Course{
	var course g.Course
	s.Db.Set("gorm:auto_preload", true).
		Where("Name = ?" , Name).
		Find(&course)

	return course
}
