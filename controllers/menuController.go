package controller

import "github.com/gin-gonic/gin"

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
