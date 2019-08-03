package models

type Deptemp struct {
	ID           int
	Employee     *Employee
	EmployeeID   uint
	Department   *Department
	DepartmentID uint
}

func (*Deptemp) TableName() string {
	return "dept_emps"
}
