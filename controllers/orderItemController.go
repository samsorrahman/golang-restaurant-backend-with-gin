package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get OrderItems",
		// })
	}
}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get OrderItemByOrder",
		// })
	}
}

func ItemsByOrder(id string) (OrderItems[] primitive.M err error) {
	
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get OrderItem",
		// })
	}
}

func updateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "update OrderItem",
		// })
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Create OrderItem",
		// })
	}
}
