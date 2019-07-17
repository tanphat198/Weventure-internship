package Repo

import (
	"github.com/jinzhu/gorm"
	g "tanphat/Struct"
)

type Student_repo struct {
	Db *gorm.DB
}

func (s *Student_repo) GetbyID(ID int) g.Student{
	var student g.Student
	s.Db.Set("gorm:auto_preload", true).
		Where(" ID= ?" , ID).
			Find(&student)

	return student
}

func (s *Student_repo) GetbyAge(Age int) g.Student{
	var student g.Student
	s.Db.Set("gorm:auto_preload", true).
		Where(" Age= ?" , Age).
		Find(&student)

	return student
}
