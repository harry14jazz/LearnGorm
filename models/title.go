package models

import "github.com/jinzhu/gorm"

//Title struct
type Title struct {
	gorm.Model
	Name       string
	EmployeeID uint
}
