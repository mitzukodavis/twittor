package routers

import (
	"encoding/json"
	"github.com/mitzukodavis/twittor/bd"
	"github.com/mitzukodavis/twittor/jwt"
	"github.com/mitzukodavis/twittor/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r  *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "usuario y/o Contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "el email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false{
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "occurrio un error al intentar generar el token correspondiente a "+err.Error(),400)
		return
	}
	resp := models.RespuestaLogin{
		Token : jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(20 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}