package models

import "github.com/jinzhu/gorm"

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

//Ingrediente is a struct
type Ingrediente struct {
	gorm.Model
	Nombre            string
	Descripcion       string
	Unidade           Unidade `gorm:"foreignkey:UnidadeID"`
	UnidadeID         int
	RecetaIngrediente []*RecetaIngrediente
}

//Receta is a struct
type Receta struct {
	gorm.Model
	Nombre            string
	Public            bool
	Tipos             []Tipo `gorm:"many2many:recetaTipos;"` //Relacion NtoN simple
	RecetaIngrediente []*RecetaIngrediente
}

//RecetaIngrediente is a struct
type RecetaIngrediente struct {
	gorm.Model
	Receta        *Receta
	RecetaID      int
	Ingrediente   *Ingrediente
	IngredienteID int
	Cantidad      float32
}
