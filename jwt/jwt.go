package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitzukodavis/twittor/models"
	"time"
)

func GeneroJWT(t models.Usuario)(string, error){
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
//verificar con el codigo original
	tokenStr, err :=  token.SignedString(miClave)
	if err != nil {
		return tokenStr,  err
	}
	return tokenStr, nil
}