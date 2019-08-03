package models

import "github.com/jinzhu/gorm"

type Title struct {
	gorm.Model
	Name       string
	EmployeeID uint
}
