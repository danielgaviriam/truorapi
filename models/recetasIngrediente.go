package models

import "github.com/jinzhu/gorm"

//recetaIngrediente is a struct
type recetaIngrediente struct {
	gorm.Model
	Receta      Receta      `gorm:"foreignkey:Receta"`
	Ingrediente Ingrediente `gorm:"foreignkey:Ingrediente"`
	Cantidad    float32
}
