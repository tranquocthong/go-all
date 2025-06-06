package main

import (
	"context"
	"log"
	"test-jwt-auth/constants"
	"test-jwt-auth/server"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	opts := options.Client().ApplyURI(constants.URI).SetServerAPIOptions(options.ServerAPI(constants.DBAPIVersion))

	dbClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalln(err)
		return
	}

	db := dbClient.Database("jwt-auth")

	defer func() {
		if err := dbClient.Disconnect(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	bsSv := server.NewBasicServer(r, db)

	bsSv.RegisterBasicRoutes()

	bsSv.Run()
}
