package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/neox-hk/truorapi/libsql"
)

func connectdb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=baasu.db.elephantsql.com port=5432 user=hjamibre dbname=hjamibre password=DKGbN1fndOho8LYzWhtjtVjxetRgxxnH")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "html/index.html")

}

//-- CRUD UNIDADES GORM
func getUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := connectdb()

	var unidad Unidade

	_ = db.Find(&unidad, params["id"])

	json.NewEncoder(w).Encode(unidad)
	defer db.Close()
	return

}
func crearUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUnidad Unidade
	_ = json.NewDecoder(r.Body).Decode(&newUnidad)

	db := connectdb()
	result := db.NewRecord(newUnidad)
	db.Create(&newUnidad)
	defer db.Close()

	if result {
		fmt.Println("Funciono")
	} else {
		fmt.Println("Pues no")
	}

}
func updateUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := connectdb()

	var unidad Unidade
	var updatedUnidad Unidade
	_ = json.NewDecoder(r.Body).Decode(&updatedUnidad)

	_ = db.Find(&unidad, params["id"])

	db.Model(&unidad).Updates(Unidade{Nombre: updatedUnidad.Nombre, Abrev: updatedUnidad.Abrev})

	defer db.Close()
}

func deleteUnidad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := connectdb()

	var unidad Unidade
	_ = db.Find(&unidad, params["id"])

	db.Delete(&unidad)

	defer db.Close()
}

//Models------

//Tipo is a struct
type Tipo struct {
	gorm.Model
	Nombre      string
	Descripcion string
}

//Receta is a struct
type Receta struct {
	gorm.Model
	Nombre string
	Public bool
	Tipos  []Tipo `gorm:"many2many:recetaTipos;"` //Relacion NtoN simple
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
	Nombre      string
	Descripcion string
	Unidade     Unidade `gorm:"foreignkey:Unidad"`
}

//recetaIngrediente is a struct
type recetaIngrediente struct {
	gorm.Model
	Receta      Receta      `gorm:"foreignkey:Receta"`
	Ingrediente Ingrediente `gorm:"foreignkey:Ingrediente"`
	Cantidad    float32
}

//End Models------

func main() {

	//Iniciarndo Router
	r := mux.NewRouter()

	//Connect to database
	db := connectdb()

	db.AutoMigrate(&Unidade{})
	db.AutoMigrate(&Ingrediente{})
	db.AutoMigrate(&Tipo{})
	db.AutoMigrate(&Receta{})
	db.AutoMigrate(&recetaIngrediente{})

	log.Println("Server started on: http://localhost:8080")

	r.HandleFunc("/", index)
	//Recetas
	r.HandleFunc("/recetas", libsql.crearReceta).Methods("POST")
	r.HandleFunc("/recetas/{id}", getReceta).Methods("GET")
	//Unidades
	r.HandleFunc("/unidades/{id}", getUnidad).Methods("GET")
	r.HandleFunc("/unidades", crearUnidad).Methods("POST")
	r.HandleFunc("/unidades/{id}", updateUnidad).Methods("PUT")
	r.HandleFunc("/unidades/{id}", deleteUnidad).Methods("DELETE")

	http.ListenAndServe(":8080", r)
	defer db.Close()

}
