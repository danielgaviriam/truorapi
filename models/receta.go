package models

import "github.com/jinzhu/gorm"

//Receta is a struct
type Receta struct {
	gorm.Model
	Nombre string
	Public bool
	Tipos  []Tipo `gorm:"many2many:recetaTipos;"` //Relacion NtoN simple
}
