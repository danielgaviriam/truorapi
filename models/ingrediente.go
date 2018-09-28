package models

import "github.com/jinzhu/gorm"

//Ingrediente is a struct
type Ingrediente struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Unidade     Unidade `gorm:"foreignkey:Unidad"`
}
