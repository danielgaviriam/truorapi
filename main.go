package main

import (
	"log"
	"net/http"

	libsql "github.com/neox-hk/truorapi/libsql"
	models "github.com/neox-hk/truorapi/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Render Home Page Api-REST
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {

	//Iniciarndo Router
	r := mux.NewRouter()

	//Handlers
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Length", "Accept-Encoding", "X-CSRF-Token"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	//Connect to database
	db := libsql.Connectdb()

	//db.AutoMigrate(models.Unidade{}, models.Tipo{}, models.Receta{}, models.Ingrediente{}, models.RecetaIngrediente{})

	log.Println("Server started on: http://localhost:8080")

	//Enlaces:
	r.HandleFunc("/", index)

	//Recetas
	r.HandleFunc("/recetas/", libsql.GetRecetas).Methods("GET")
	r.HandleFunc("/recetas/", libsql.CrearReceta).Methods("POST")
	r.HandleFunc("/recetas/{id}", libsql.GetReceta).Methods("GET")
	r.HandleFunc("/recetas/{id}", libsql.UpdateReceta).Methods("PUT")
	r.HandleFunc("/recetas/{id}", libsql.DeleteReceta).Methods("DELETE")

	//Unidades
	r.HandleFunc("/unidades/", libsql.GetUnidades).Methods("GET")
	r.HandleFunc("/unidades/{id}", libsql.GetUnidad).Methods("GET")
	r.HandleFunc("/unidades/", libsql.CrearUnidad).Methods("POST")
	r.HandleFunc("/unidades/{id}", libsql.UpdateUnidad).Methods("PUT")
	r.HandleFunc("/unidades/{id}", libsql.DeleteUnidad).Methods("DELETE")

	//Tipo
	r.HandleFunc("/tipos/", libsql.GetTipos).Methods("GET")
	r.HandleFunc("/tipos/{id}", libsql.GetTipo).Methods("GET")
	r.HandleFunc("/tipos/", libsql.CrearTipo).Methods("POST")
	r.HandleFunc("/tipos/{id}", libsql.UpdateTipo).Methods("PUT")
	r.HandleFunc("/tipos/{id}", libsql.DeleteTipo).Methods("DELETE")

	//Ingredientes
	r.HandleFunc("/ingredientes/", libsql.GetIngredientes).Methods("GET")
	r.HandleFunc("/ingredientes/", libsql.CrearIngrediente).Methods("POST")
	r.HandleFunc("/ingredientes/{id}", libsql.GetIngrediente).Methods("GET")
	r.HandleFunc("/ingredientes/{id}", libsql.UpdateIngrediente).Methods("PUT")
	r.HandleFunc("/ingredientes/{id}", libsql.DeleteIngrediente).Methods("DELETE")

	http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r))

	//Close Connect
	defer db.Close()
}
