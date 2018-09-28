package libsql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getReceta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	fmt.Println(params["id"])

	//json.NewEncoder(w).Encode(&Book{})
}

//getCrearReceta exportable
func getCrearReceta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := json.NewDecoder(r.Body)
	fmt.Println(test)

}
