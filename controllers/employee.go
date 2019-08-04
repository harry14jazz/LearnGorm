package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harry/jazzcorp/config"
	"github.com/harry/jazzcorp/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//CreateEmployee Func
func CreateEmployee(c *gin.Context) {
	item := models.Employee{
		FirstName: c.PostForm("first_name"),
		LastName:  c.PostForm("last_name"),
		BirthDate: c.PostForm("birth_date"),
		Gender:    c.PostForm("gender"),
		HireDate:  c.PostForm("hire_date"),
	}

	config.DB.Create(&item)

	c.JSON(201, gin.H{
		"status": "Created something OK",
		"data":   item.Serialize(),
	})
}
