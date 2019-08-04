package models

import "github.com/jinzhu/gorm"

//Department struct
type Department struct {
	gorm.Model
	DeptEmp  []*Deptemp
	DeptMng  []DeptMng
	DeptName string
}
