package models

import "github.com/jinzhu/gorm"

//Tipo is a struct
type Tipo struct {
	gorm.Model
	Nombre      string
	Descripcion string
}
