package models

import (
	"github.com/harry/jazzcorp/common"
	"github.com/jinzhu/gorm"
)

//Salary struct
type Salary struct {
	gorm.Model
	Salary     int
	Employee   Employee `gorm:"foreignkey:EmployeeID"`
	EmployeeID uint
}

//Serialize func
func (s Salary) Serialize() common.JSON {
	return common.JSON{
		"id":          s.ID,
		"salary":      s.Salary,
		"employee_id": s.Employee.Serialize(),
		"created_at":  s.CreatedAt,
	}
}
