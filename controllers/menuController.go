package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samsorrahman/golang-rest-backend/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection("menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Menus",
		// })
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Menu",
		// })
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Create Menu",
		// })
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Update Menu",
		// })
	}
}
