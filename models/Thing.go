package models

import (
	"github.com/jinzhu/gorm"
)

// Thing belongs to List
type Thing struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	ListID uint   `json:"list_id" binding:"required"`
}
