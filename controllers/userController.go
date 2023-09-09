package controller

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get Users",
		// })
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Get User",
		// })
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Sign Up",
		// })
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Login",
		// })
	}
}

func HashPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Hash Password",
		// })
	}
}

func VerfiyPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Verify Password",
		// })
	}
}
