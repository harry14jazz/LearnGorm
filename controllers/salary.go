package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/harry/jazzcorp/common"
	"github.com/harry/jazzcorp/config"
	"github.com/harry/jazzcorp/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Employee alias
type Employee = models.Employee

// JSON alias
type JSON = common.JSON

//CreateSalary Func
// func CreateSalary(c *gin.Context) {
// 	gaji, _ := strconv.Atoi(c.PostForm("salary"))
// 	employee := c.MustGet("employee_id").(Employee)
// 	item := models.Salary{
// 		Salary:   gaji,
// 		Employee: employee,
// 	}

// 	config.DB.Create(&item)

// 	c.JSON(201, gin.H{
// 		"status": "Created something OK",
// 		"data":   item.Serialize(),
// 	})
// }

// BikinGaji asu
func BikinGaji(c *gin.Context) {
	salaries := []models.Salary{}

	c.BindJSON(&salaries)

	tx := config.DB.Begin()
	if err := tx.Create(&salaries).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		c.AbortWithStatus(404)
	} else {
		tx.Commit()
		c.JSON(201, salaries)
	}
}
