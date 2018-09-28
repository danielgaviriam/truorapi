package models

import "github.com/jinzhu/gorm"

//Ingrediente is a struct
type Ingrediente struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Unidade     Unidade `gorm:"foreignkey:Unidad"`
}

//Receta is a struct
type Receta struct {
	gorm.Model
	Nombre string
	Public bool
	Tipos  []Tipo `gorm:"many2many:recetaTipos;"` //Relacion NtoN simple
}

//recetaIngrediente is a struct
type recetaIngrediente struct {
	gorm.Model
	Receta      Receta      `gorm:"foreignkey:Receta"`
	Ingrediente Ingrediente `gorm:"foreignkey:Ingrediente"`
	Cantidad    float32
}

//Tipo is a struct
type Tipo struct {
	gorm.Model
	Nombre      string
	Descripcion string
}

//Unidade is a struct
type Unidade struct {
	gorm.Model
	Nombre string
	Abrev  string
}
