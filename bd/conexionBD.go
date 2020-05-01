package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//MongoCN its other
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:ubuntu64@cluster0-voq2g.mongodb.net/test?retryWrites=true&w=majority")
// ConectarBD() other

//func ConectarBD() es para
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error() )
		return client
	}
	err = client.Ping( context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa cin ka bd")
	return client
}
//func ChequeoConnection() es para
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
