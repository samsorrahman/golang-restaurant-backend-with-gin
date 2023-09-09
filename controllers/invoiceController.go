package controller

import "github.com/gin-gonic/gin"

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Invoices",
		// })
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Invoice",
		// })
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Create Invoice",
		// })
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Update Invoice",
		// })
	}
}
