package models

import "github.com/jinzhu/gorm"

type Salary struct {
	gorm.Model
	Salary     int
	EmployeeID uint
}
