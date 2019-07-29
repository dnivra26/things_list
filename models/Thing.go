package models

import (
	"github.com/jinzhu/gorm"
)

// Thing belongs to List
type Thing struct {
	gorm.Model
	Title  string
	ListID uint
}
