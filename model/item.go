package model

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `gorm:"not null"`
	Price float64 
}