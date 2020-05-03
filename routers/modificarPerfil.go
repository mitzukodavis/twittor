package routers

import (
	"encoding/json"
	"github.com/mitzukodavis/twittor/bd"
	"github.com/mitzukodavis/twittor/models"
	"net/http"
)

func ModificarPerfil(w http.ResponseWriter, r * http.Request){
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+ err.Error(), 400)
		return
	}
	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err !=nil{
		http.Error(w, "ocurrio un error al intentar modificar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logrado modificar el registro de usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}