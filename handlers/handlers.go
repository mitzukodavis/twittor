package handlers

import (
	"github.com/mitzukodavis/twittor/middlew"
	"github.com/mitzukodavis/twittor/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores pongo a setear
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
