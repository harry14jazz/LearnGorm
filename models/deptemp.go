package models

//Deptemp struct
type Deptemp struct {
	ID           int
	Employee     *Employee
	EmployeeID   uint
	Department   *Department
	DepartmentID uint
}

//TableName func
func (*Deptemp) TableName() string {
	return "dept_emps"
}
