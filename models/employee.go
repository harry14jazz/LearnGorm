package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	Salary    []Salary
	Title     []Title
	DeptMng   []DeptMng
	DeptEmp   []*Deptemp
	FirstName string
	LastName  string
	BirthDate string
	Gender    string
	HireDate  string
}
