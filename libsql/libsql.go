package libsql

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //use
	model "github.com/neox-hk/truorapi/models"
)

//Connectdb export with GORM
func Connectdb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=baasu.db.elephantsql.com port=5432 user=zgxxbdhj dbname=zgxxbdhj password=whbzu3uTA38i6VjFoUS7w6S8xzdbv1Wh")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//--Crud de Recetas

//GetReceta export
func GetReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var receta model.Receta

	db.First(&receta, params["id"])

	json.NewEncoder(w).Encode(receta)
	defer db.Close()
	return
}

//CrearReceta export
func CrearReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := Connectdb()

	var newReceta model.Receta
	_ = json.NewDecoder(r.Body).Decode(&newReceta)

	exists := db.Where("Nombre = ?", newReceta.Nombre).First(&newReceta).RecordNotFound()

	if exists {
		result := db.NewRecord(newReceta)
		db.Create(&newReceta)

		if result {
			json.NewEncoder(w).Encode(1)
		} else {
			json.NewEncoder(w).Encode(0)
		}
		defer db.Close()

	} else {
		json.NewEncoder(w).Encode(3)
		defer db.Close()
	}
	return

}

//--End Crud de Recetas

//--Crud de Unidades

//GetUnidad export
func GetUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var unidad model.Unidade
	_ = db.First(&unidad, params["id"])

	json.NewEncoder(w).Encode(unidad)
	defer db.Close()
	return

}

//CrearUnidad export
func CrearUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUnidad model.Unidade
	_ = json.NewDecoder(r.Body).Decode(&newUnidad)

	db := Connectdb()

	// Existe en la bd?
	exists := db.Where("Nombre = ?", newUnidad.Nombre).First(&newUnidad).RecordNotFound()

	if exists {
		result := db.NewRecord(newUnidad)
		db.Create(&newUnidad)

		if result {
			json.NewEncoder(w).Encode(1)
		} else {
			json.NewEncoder(w).Encode(0)
		}
		defer db.Close()

	} else {
		json.NewEncoder(w).Encode(3)
		defer db.Close()
	}
	return

}

//UpdateUnidad export
func UpdateUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var unidad model.Unidade
	var updatedUnidad model.Unidade
	_ = json.NewDecoder(r.Body).Decode(&updatedUnidad)

	_ = db.Find(&unidad, params["id"])

	db.Model(&unidad).Updates(model.Unidade{Nombre: updatedUnidad.Nombre, Abrev: updatedUnidad.Abrev})

	defer db.Close()
}

//DeleteUnidad export
func DeleteUnidad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var unidad model.Unidade
	_ = db.Find(&unidad, params["id"])

	db.Delete(&unidad)

	defer db.Close()
}

//--End Crud de Unidades

//--Crud de Tipos

//GetTipo export
func GetTipo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var tipo model.Tipo

	_ = db.Find(&tipo, params["id"])

	json.NewEncoder(w).Encode(tipo)
	defer db.Close()
	return

}

//CrearTipo export
func CrearTipo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newTipo model.Tipo
	_ = json.NewDecoder(r.Body).Decode(&newTipo)

	db := Connectdb()

	// Existe en la bd?
	exists := db.Where("Nombre = ?", newTipo.Nombre).First(&newTipo).RecordNotFound()

	if exists {
		result := db.NewRecord(newTipo)
		db.Create(&newTipo)

		if result {
			json.NewEncoder(w).Encode(1)
		} else {
			json.NewEncoder(w).Encode(0)
		}
		defer db.Close()

	} else {
		json.NewEncoder(w).Encode(3)
		defer db.Close()
	}
	return
}

//UpdateTipo export
func UpdateTipo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var tipo model.Unidade
	var updatedTipo model.Tipo
	_ = json.NewDecoder(r.Body).Decode(&updatedTipo)

	_ = db.Find(&tipo, params["id"])

	db.Model(&tipo).Updates(model.Tipo{Nombre: updatedTipo.Nombre, Descripcion: updatedTipo.Descripcion})

	defer db.Close()
}

//DeleteTipo export
func DeleteTipo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var tipo model.Unidade
	_ = db.Find(&tipo, params["id"])

	db.Delete(&tipo)

	defer db.Close()
}

//--End Crud de Tipos
