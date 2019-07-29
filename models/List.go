package models

import "github.com/jinzhu/gorm"

//List contains things
type List struct {
	gorm.Model
	Name   string
	Things []Thing
}
