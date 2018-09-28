package models

import "github.com/jinzhu/gorm"

//Unidade is a struct
type Unidade struct {
	gorm.Model
	Nombre string
	Abrev  string
}
