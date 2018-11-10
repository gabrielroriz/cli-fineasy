package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title string `gorm:"type:varchar(30);not null"`
}