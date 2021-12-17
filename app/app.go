package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()

	router := mux.NewRouter()

	// define routers
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	
	router.HandleFunc("/customers/{customer_id:[0-9]+}",getCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getCustomers( w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomers(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Post request recevied")
}