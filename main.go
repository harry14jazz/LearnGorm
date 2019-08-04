package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harry/jazzcorp/config"
	"github.com/harry/jazzcorp/controllers"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/employee", controllers.CreateEmployee)
		// v1.POST("/salary", controllers.CreateSalary)
		v1.POST("/gaji", controllers.BikinGaji)
	}
	router.Run()
}
