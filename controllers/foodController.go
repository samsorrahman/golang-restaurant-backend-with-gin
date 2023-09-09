package controller

import (
	"context"
	// "fmt"
	"github.com/samsorrahman/golang-rest-backend/database"
	"github.com/samsorrahman/golang-rest-backend/models"

	// "log"
	"math"
	"net/http"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection("food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Foods",
		// })
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Create Food",
		// })
	}
}

func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output, .5)) / output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Update Food",
		// })
	}
}
