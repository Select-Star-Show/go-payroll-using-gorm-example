// Copyright 2025 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name string    `json:"name"`
	Role string    `json:"role"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
