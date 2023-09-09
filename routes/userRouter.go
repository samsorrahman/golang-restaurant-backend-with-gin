package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samsorrahman/golang-rest-backend/controllers"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetUser)
	incomingRoutes.POST("/users/signup", controllers.SignUp)
	incomingRoutes.PUT("/users/login", controllers.Login)
}
