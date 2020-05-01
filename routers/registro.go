package routers

import (
	"encoding/json"
	"github.com/mitzukodavis/twittor/bd"
	"github.com/mitzukodavis/twittor/models"
	"net/http"
)

//Registro es la funcion para crear en la BD

func Registro(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "error en la base da datos"+err.Error(), 400)
		return
	}
	if len(t.Email)==0{
		http.Error(w, "el email del usuario es requerido", 400)
		return
	}
	if len(t.Password)<6{
		http.Error(w, "debe especificar una contraseña de al menos 6 caracteres", 400)
		return

	}
	_, encontrado,_:= bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado==true {
		http.Error(w, "Ya existe un usuario con ese correo", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err !=nil {            
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuarios"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "ni se ligrado insertar el regiustro", 400)

		return
	}
	w.WriteHeader(http.StatusCreated)
}
