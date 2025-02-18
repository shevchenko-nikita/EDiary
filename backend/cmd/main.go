package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"github.com/shevchenko-nikita/EDiary/internals/routes"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.ConnectDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	handler := handlers.NewHandler(db)
	router := gin.Default()
	routes.InitRoutes(router, handler)

	router.Run(":8080")

}
