package models

import "github.com/jinzhu/gorm"

type DeptMng struct {
	gorm.Model
	EmployeeID   uint
	DepartmentID uint
}
