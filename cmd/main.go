package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"github.com/shevchenko-nikita/EDiary/internals/routes"
)

func main() {

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
