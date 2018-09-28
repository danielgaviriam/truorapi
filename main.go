package main

import (
	"log"
	"net/http"

	libsql "github.com/neox-hk/truorapi/libsql"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Render Home Page Api-REST
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {

	//Iniciarndo Router
	r := mux.NewRouter()

	//Connect to database
	db := libsql.Connectdb()
	/*
		db.AutoMigrate(&Unidade{})
		db.AutoMigrate(&Ingrediente{})
		db.AutoMigrate(&Tipo{})
		db.AutoMigrate(&Receta{})
		db.AutoMigrate(&recetaIngrediente{})
	*/

	log.Println("Server started on: http://localhost:8080")

	//Enlaces:
	r.HandleFunc("/", index)
	//Recetas
	r.HandleFunc("/recetas/", libsql.GetReceta).Methods("POST")
	r.HandleFunc("/recetas/{id}", libsql.GetReceta).Methods("GET")
	//Unidades
	r.HandleFunc("/unidades/{id}", libsql.GetUnidad).Methods("GET")
	r.HandleFunc("/unidades/", libsql.CrearUnidad).Methods("POST")
	r.HandleFunc("/unidades/{id}", libsql.UpdateUnidad).Methods("PUT")
	r.HandleFunc("/unidades/{id}", libsql.DeleteUnidad).Methods("DELETE")
	//Tipo
	r.HandleFunc("/tipos/{id}", libsql.GetTipo).Methods("GET")
	r.HandleFunc("/tipos/", libsql.CrearTipo).Methods("POST")
	r.HandleFunc("/tipos/{id}", libsql.UpdateTipo).Methods("PUT")
	r.HandleFunc("/tipos/{id}", libsql.DeleteTipo).Methods("DELETE")

	http.ListenAndServe(":8080", r)

	//Close Connect
	defer db.Close()
}
