package models

import "github.com/jinzhu/gorm"

type Wallet struct {
	gorm.Model
	Title string
}
