package backend

import (
	"net/http"

	"github.com/gorilla/mux"
)						
func main() {
	// Initialize the database
	InitDB("user=youruser dbname=yourdb sslmode=disable")
	defer CloseDB()

	// Set up the router
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")			

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)	
}