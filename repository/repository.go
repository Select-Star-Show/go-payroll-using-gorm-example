// Copyright 2025 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package repository

import (
	"github.com/Select-Star-Show/go-payroll-using-gorm-example/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(employee models.Employee) (models.Employee, error)
	FindAll() ([]models.Employee, error)
	FindById(id uuid.UUID) (models.Employee, error)
	Delete(id uuid.UUID) error
	Update(employee models.Employee) (models.Employee, error)
	SqlQuery(query string, values ...interface{}) ([]models.Employee, error)
}

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (repo *EmployeeRepository) Save(employee *models.Employee) (*models.Employee, error) {
	err := repo.DB.Create(&employee).Error
	return employee, err
}

func (repo *EmployeeRepository) FindAll() (*[]models.Employee, error) {
	var employees []models.Employee
	err := repo.DB.Find(&employees).Error
	return &employees, err
}

func (repo *EmployeeRepository) FindById(id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee
	err := repo.DB.First(&employee, id).Error
	return &employee, err
}

func (repo *EmployeeRepository) Delete(id uuid.UUID) error {
	err := repo.DB.Delete(&models.Employee{}, id).Error
	return err
}

func (repo *EmployeeRepository) Update(employee *models.Employee) (*models.Employee, error) {
	err := repo.DB.Save(&employee).Error
	return employee, err
}

func (repo *EmployeeRepository) SqlQuery(query string, values ...interface{}) ([]models.Employee, error) {
	var employees []models.Employee
	err := repo.DB.Raw(query, values).Scan(&employees).Error
	return employees, err
}
