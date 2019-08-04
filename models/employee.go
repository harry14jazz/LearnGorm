package models

import (
	"github.com/harry/jazzcorp/common"
	"github.com/jinzhu/gorm"
)

//Employee struct
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

//Serialize employee data
func (e *Employee) Serialize() common.JSON {
	return common.JSON{
		"employee_id": e.ID,
		"first_name":  e.FirstName,
		"last_name":   e.LastName,
		"birth_date":  e.BirthDate,
		"gender":      e.Gender,
		"hire_date":   e.HireDate,
	}
}

//Read data emp
func (e *Employee) Read(m common.JSON) {
	e.ID = uint(m["employee_id"].(float64))
}
