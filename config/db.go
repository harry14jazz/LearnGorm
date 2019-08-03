package config

import (
	"github.com/harry/jazzcorp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:kurkur14@/jazzcorp?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Employee{})
	DB.AutoMigrate(&models.Department{})
	DB.AutoMigrate(&models.Deptemp{})
	DB.AutoMigrate(&models.Salary{}).AddForeignKey("employee_id", "employees(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&models.Title{}).AddForeignKey("employee_id", "employees(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&models.DeptMng{}).AddForeignKey("employee_id", "employees(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&models.DeptMng{}).AddForeignKey("department_id", "departments(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&models.Deptemp{}).AddForeignKey("employee_id", "employees(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&models.Deptemp{}).AddForeignKey("department_id", "departments(id)", "CASCADE", "CASCADE")

	DB.Model(&models.Employee{}).Related(&models.Salary{})
	DB.Model(&models.Employee{}).Related(&models.Title{})
	DB.Model(&models.Employee{}).Related(&models.DeptMng{})
	DB.Model(&models.Department{}).Related(&models.DeptMng{})
}
