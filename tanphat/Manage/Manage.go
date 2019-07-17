package Manage

import (
	g "tanphat/Struct"
)

type Manage struct {
	Student StudentInfor
	Class 	ClassInfor
	Course 	CourseInfor
}

type ClassInfor interface {
	GetbyID(ID int) g.Class
	GetbyName(Name string) g.Class
}

type CourseInfor interface {
	GetbyID(ID int) g.Course
	GetbyName(Name string) g.Course
}

type StudentInfor interface {
	GetbyID(ID int) g.Student
	GetbyAge(Age int) g.Student

}
