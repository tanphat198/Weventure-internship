package Repo

import (
	"github.com/jinzhu/gorm"
	g "tanphat/Struct"
)

type Class_repo struct {
	Db *gorm.DB
}

func (s *Class_repo) GetbyID(ID int) g.Class{
	var class g.Class
	s.Db.Set("gorm:auto_preload", true).
		Where("ID = ?" , ID).
		Find(&class)

	return class
}

func (s *Class_repo) GetbyName(Name string) g.Class{
	var class g.Class
	s.Db.Set("gorm:auto_preload", true).
		Where("Name = ?" , Name).
		Find(&class)

	return class
}
