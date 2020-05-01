package middlew

import (
	"github.com/mitzukodavis/twittor/bd"
	"net/http"
)
//ChequeoBD es el middleware que verifica el estado de la conexcion de la db
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0{
			http.Error(w, "conexion de  perdida de la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}