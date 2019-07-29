package models

import "github.com/jinzhu/gorm"

//List contains things
type List struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Things []Thing
}
