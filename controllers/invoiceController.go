package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samsorrahman/golang-rest-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var invoiceCollection *mongo.Collection = database.OpenCollection("invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice items"})
		}

		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)
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
