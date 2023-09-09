package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samsorrahman/golang-rest-backend/database"
	"github.com/samsorrahman/golang-rest-backend/middleware"
	"github.com/samsorrahman/golang-rest-backend/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = 8000
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemeRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
