package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	s "tanphat/Struct"
	m "tanphat/Manage"
	"tanphat/Repo"
)

func main() {
	var man  m.Manage

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&s.Student{})
	db.AutoMigrate(&s.Class{})
	db.AutoMigrate(&s.Course{})

	// Create
	db.Create(&s.Student{Name:"Tan Phat", Age:21})
	db.Create(&s.Student{Name:"Anh Phuong", Age:23})
	db.Create(&s.Class{Name:"CS161"})
	db.Create(&s.Class{Name:"CS422"})
	db.Create(&s.Course{Name:"CTTT"})

	// Set variable
	Phat := s.Student{Name:"Tan Phat", Age:21}
	Phuong := s.Student{Name:"Anh Phuong", Age:23}
	CS422 := s.Class{Name:"CS422"}
	CS161 := s.Class{Name :"CS161"}
	CTTT := s.Course{Name:"16CTT"}

	// Save
	db.Save(&CS161)
	db.Save(&CTTT)

	// Association
	db.Model(&CS161).Association("Students").Append(Phat,Phuong)
	db.Model(&CTTT).Association("Classes").Append(CS161,CS422)

	man.Student = &Repo.Student_repo{Db:db}
	man.Class = &Repo.Class_repo{Db:db}
	man.Course = &Repo.Course_repo{Db: db}

	res := man.Student.GetbyAge(21)
	resp := man.Class.GetbyName("CS161")
	ressp := man.Course.GetbyName("16CTT")

	fmt.Println(res)
	fmt.Println(resp)
	fmt.Println(ressp)

	db.DropTable(s.Student{})
	db.DropTable(s.Class{})
	db.DropTable(s.Course{})
}
