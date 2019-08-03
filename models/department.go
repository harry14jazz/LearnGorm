package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	DeptEmp  []*Deptemp
	DeptMng  []DeptMng
	DeptName string
}
