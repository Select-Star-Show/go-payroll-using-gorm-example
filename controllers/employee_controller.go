// Copyright 2025 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package controllers

import (
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/models"
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type EmployeeController struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeController(repo *repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{repo: repo}
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved, err := ec.repo.Save(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, saved)
}

func (ec *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ec.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, employees)
}

func (ec *EmployeeController) FindEmployeeById(c *gin.Context) {
	idStr := c.Param("uuid")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	employee, err := ec.repo.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, employee)
}

func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("uuid")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	err = ec.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	idStr := c.Param("uuid")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	found, err := ec.repo.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	found.Name = employee.Name
	found.Role = employee.Role

	updated, err := ec.repo.Update(found)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updated)
}

func (ec *EmployeeController) GetEmployeesByRole(c *gin.Context) {
	role := c.Param("role")

	employees, err := ec.repo.SqlQuery("SELECT * FROM employees AS OF SYSTEM TIME '-30s' WHERE role = ?", role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, employees)
}
