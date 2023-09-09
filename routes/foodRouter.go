package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samsorrahman/golang-rest-backend/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods)
	incomingRoutes.GET("/foods/:food_id", controller.GetFood)
	incomingRoutes.POST("/foods", controller.CreateFood)
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood)
	// incomingRoutes.DELETE("/foods/:id", controller.DeleteFood)
}
