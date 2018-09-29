package libsql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //use
	model "github.com/neox-hk/truorapi/models"
)

//Connectdb export with GORM
func Connectdb() *gorm.DB {

	db, err := gorm.Open("postgres", "host=baasu.db.elephantsql.com port=5432 user=zgxxbdhj dbname=zgxxbdhj password=P8GwHaKaKoHp4IYLvYMveuoAe0BIR0xn")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//--Crud de Recetas

//GetRecetas export
func GetRecetas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := Connectdb()

	recetas := []model.Receta{}
	db.Preload("Tipos").Preload("RecetaIngrediente").Find(&recetas)

	json.NewEncoder(w).Encode(recetas)
	defer db.Close()
	return
}

//GetReceta export
func GetReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var receta model.Receta

	db.First(&receta, params["id"])

	//Load many2many Information
	db.Preload("Tipos").Find(&receta, receta.ID)
	db.Preload("RecetaIngrediente").Find(&receta, receta.ID)
	db.Preload("Ingrediente").Find(&receta.RecetaIngrediente)

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

//UpdaterReceta export
func UpdaterReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	//registro Antiguo
	var receta model.Receta
	//Registro Nuevo
	var updatedReceta model.Receta
	_ = json.NewDecoder(r.Body).Decode(&updatedReceta)

	_ = db.First(&receta, params["id"])

	//Actualizando Ingrediente
	db.Model(&receta).Updates(&updatedReceta)

	defer db.Close()
}

//DeleteReceta export
func DeleteReceta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var receta model.Receta
	_ = db.Find(&receta, params["id"])

	_ = db.Where("receta_id = ?", params["id"]).Delete(model.RecetaIngrediente{})

	receta.Tipos = []model.Tipo{}
	_ = db.Save(&receta)

	db.Debug().Delete(&receta)

	defer db.Close()
}

//--End Crud de Recetas

//--Crud de Ingredientes

//GetIngrediente export
func GetIngrediente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var ingrediente model.Ingrediente
	var unidad model.Unidade
	_ = db.First(&ingrediente, params["id"])
	_ = db.First(&unidad, ingrediente.UnidadeID)
	ingrediente.Unidade = unidad

	json.NewEncoder(w).Encode(ingrediente)
	defer db.Close()
	return

}

//CrearIngrediente export
func CrearIngrediente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := Connectdb()

	var newIngrediente model.Ingrediente
	_ = json.NewDecoder(r.Body).Decode(&newIngrediente)

	exists := db.Where("Nombre = ?", newIngrediente.Nombre).First(&newIngrediente).RecordNotFound()

	if exists {
		result := db.NewRecord(newIngrediente)
		db.Create(&newIngrediente)

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

//UpdateIngrediente export
func UpdateIngrediente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	//registro Antiguo
	var ingrediente model.Ingrediente
	//Registro Nuevo
	var updatedIngrediente model.Ingrediente
	_ = json.NewDecoder(r.Body).Decode(&updatedIngrediente)

	_ = db.Find(&ingrediente, params["id"])

	//Actualizando Ingrediente
	db.Model(&ingrediente).Updates(model.Ingrediente{Nombre: updatedIngrediente.Nombre, Descripcion: updatedIngrediente.Descripcion, UnidadeID: updatedIngrediente.UnidadeID})

	defer db.Close()
}

//DeleteIngrediente export
func DeleteIngrediente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	var ingrediente model.Ingrediente
	_ = db.Find(&ingrediente, params["id"])

	db.Debug().Delete(&ingrediente)

	defer db.Close()
}

//--End Crud de Ingredientes

//--Crud de Unidades

//GetUnidad export
func GetUnidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := Connectdb()

	fmt.Println("test")

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
	_ = db.First(&tipo, params["id"])

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

	var tipo model.Tipo
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
