package controller

import "github.com/gin-gonic/gin"

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Tables",
		// })
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Table",
		// })
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Create Table",
		// })
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Update Table",
		// })
	}
}
