// Copyright 2025 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package app

import (
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/controllers"
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/models"
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/repository"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	// Start
	asciiArt := figure.NewFigure("Employee Payroll", "standard", true)
	asciiArt.Print()

	// Database
	repository.ConnectDatabase()
	if err := repository.DB.AutoMigrate(&models.Employee{}); err != nil {
		panic("Failed to migrate database:")
	}
	employeeRepository := repository.NewEmployeeRepository(repository.DB)

	// Database - Init
	employeeRepository.Save(
		&models.Employee{
			Name: "Felipe",
			Role: "advocate",
		})
	employeeRepository.Save(
		&models.Employee{
			Name: "Glen",
			Role: "engineer",
		})

	// Web API
	controller := controllers.NewEmployeeController(employeeRepository)
	web := gin.Default()
	api := web.Group("/api")
	{
		api.POST("/employees", controller.CreateEmployee)
		api.GET("/employees", controller.GetEmployees)
		api.GET("/employees/:uuid", controller.FindEmployeeById)
		api.DELETE("/employees/:uuid", controller.DeleteEmployee)
		api.PUT("/employees/:uuid", controller.UpdateEmployee)
		api.GET("/employees/recent/role/:role", controller.GetEmployeesByRole)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085" // Default port
	}
	if err := web.Run(":" + port); err != nil {
		panic(err)
	}

}

func Start() {}
