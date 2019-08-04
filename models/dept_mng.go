package models

import "github.com/jinzhu/gorm"

//DeptMng struct
type DeptMng struct {
	gorm.Model
	EmployeeID   uint
	DepartmentID uint
}
