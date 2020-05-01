package main

import (
	"log"
	"github.com/mitzukodavis/twittor/handlers"
	"github.com/mitzukodavis/twittor/bd"
)
func main() {
	if bd.ChequeoConnection()==0 {
		log.Fatal("sin conexcion a la bd")
		return
	}
	handlers.Manejadores()
}
