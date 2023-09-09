package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samsorrahman/golang-rest-backend/controllers"
)

// OrderItemRoutes is a function that takes in a pointer to a gin.Engine and adds the routes for the OrderItem model

func OrderItemeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems)
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem)
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem)
	incomingRoutes.PATCH("/orderItems/:table_id", controller.UpdateOrderItem)
	// incomingRoutes.DELETE("/order_items/:id", controller.DeleteOrderItem)
}
